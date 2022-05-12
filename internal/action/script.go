package action

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"reflect"

	"go.starlark.net/starlark"
	"go.starlark.net/starlarkjson"
	"go.starlark.net/starlarkstruct"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/klog/v2"

	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
	plugininternal "github.com/ankeesler/spirits/internal/apis/spirits/plugin"
	"github.com/ankeesler/spirits/internal/loadcache"
)

type script struct {
	codec runtime.Codec

	starlarkCodec *starlarkCodec

	program *starlark.Program
}

func Script(apiVersion, source string, scheme *runtime.Scheme) (spiritsinternal.Action, error) {
	s := &script{}

	codecs := serializer.NewCodecFactory(scheme)
	encoder, err := getJSONEncoder(codecs)
	if err != nil {
		return nil, fmt.Errorf("get json encoder for scheme %s: %w", scheme, err)
	}

	actionRunGV, err := schema.ParseGroupVersion(apiVersion)
	if err != nil {
		return nil, fmt.Errorf("parse api version: %w", err)
	}

	s.codec = codecs.CodecForVersions(encoder, codecs.UniversalDecoder(), actionRunGV, nil)

	s.starlarkCodec, err = newStarlarkCodec()
	if err != nil {
		return nil, fmt.Errorf("new starlark codec: %w", err)
	}

	// TODO: fuzz in all fields?
	predeclared, err := s.getPredeclared(&plugininternal.ActionRun{
		Spec: plugininternal.ActionRunSpec{
			From: spiritsinternal.SpiritSpec{},
			To:   spiritsinternal.SpiritSpec{},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("get script predeclared symbols for compile: %w", err)
	}

	_, program, err := starlark.SourceProgram("<actionscript:source>", source, predeclared.Has)
	if err != nil {
		return nil, err
	}
	s.program = program

	return s, nil
}

func (s *script) Run(ctx context.Context, from, to *spiritsinternal.Spirit) error {
	out := bytes.NewBuffer([]byte{})
	thread := &starlark.Thread{
		Name: "<actionscript:main>",
		Print: func(thread *starlark.Thread, msg string) {
			fmt.Fprintf(out, "%s > %s", thread.Name, msg)
			klog.V(1).InfoS("run action script", "thread", thread.Name, "message", msg)
		},
		Load: func(thread *starlark.Thread, module string) (starlark.StringDict, error) {
			return loadcache.Load(module)
		},
	}
	predeclared, err := s.getPredeclared(&plugininternal.ActionRun{
		Spec: plugininternal.ActionRunSpec{
			From: from.Spec,
			To:   to.Spec,
		},
	})
	if err != nil {
		return fmt.Errorf("get script predeclared symbols for run: %w", err)
	}

	globals, err := s.run(ctx, thread, predeclared)
	if err != nil {
		return fmt.Errorf("run script: %w (out: %q)", err, out.String())
	}

	actionRunStarlarkValue := globals["status"]
	if actionRunStarlarkValue == nil {
		return fmt.Errorf("must set 'status' in script")
	}

	var actionRun plugininternal.ActionRun
	if err := s.fromStarlarkValue(actionRunStarlarkValue, &actionRun); err != nil {
		return fmt.Errorf("convert from starlark value %q: %w", actionRunStarlarkValue.String(), err)
	}

	from.Spec = actionRun.Status.From
	to.Spec = actionRun.Status.To

	return nil
}

func (s *script) DeepCopyAction() spiritsinternal.Action {
	// TODO: copy codec

	var copiedProgram *starlark.Program
	if s.program != nil {
		var err error
		copiedProgram, err = copyProgram(s.program)
		if err != nil {
			// TODO: probably shouldn't panic here
			panic("cannot copy program: " + err.Error())
		}
	}

	return &script{
		program: copiedProgram,
	}
}

func (s *script) getPredeclared(actionRun *plugininternal.ActionRun) (starlark.StringDict, error) {
	actionRunJSON, err := runtime.Encode(s.codec, actionRun)
	if err != nil {
		return nil, fmt.Errorf("encode ActionRun to JSON: %w", err)
	}

	actionRunStarlarkValue, err := s.starlarkCodec.decode(string(actionRunJSON))
	if err != nil {
		return nil, fmt.Errorf("decode ActionRun JSON to starlark value: %w", err)
	}

	actionRunStarlarkAttrs, ok := actionRunStarlarkValue.(starlark.HasAttrs)
	if !ok {
		return nil, fmt.Errorf(
			"cannot cast starlark value %q with type %T to %T",
			actionRunStarlarkValue.String(),
			actionRunStarlarkValue,
			starlark.HasAttrs(nil),
		)
	}

	// TODO: remove unneeded code
	// TODO: don't refer to starlark in errors

	apiVersionStarlarkValue, err := actionRunStarlarkAttrs.Attr("apiVersion")
	if err != nil {
		return nil, fmt.Errorf("get attr api version from ActionRun script value %q: %w", actionRunStarlarkAttrs.String(), err)
	}
	specStarlarkValue, err := actionRunStarlarkAttrs.Attr("spec")
	if err != nil {
		return nil, fmt.Errorf("get attr spec version from ActionRun script value %q: %w", actionRunStarlarkAttrs.String(), err)
	}

	return starlark.StringDict{
		"apiVersion": apiVersionStarlarkValue,
		"spec":       specStarlarkValue,
	}, nil
}

func (s *script) run(
	ctx context.Context,
	thread *starlark.Thread,
	predeclared starlark.StringDict,
) (starlark.StringDict, error) {
	type starlarkInitRet struct {
		globals starlark.StringDict
		err     error
	}
	done := make(chan *starlarkInitRet)
	defer close(done)
	go func() {
		globals, err := s.program.Init(thread, predeclared)
		if err != nil {
			err = fmt.Errorf("script failed: %w", err)
		}
		done <- &starlarkInitRet{globals: globals, err: err}
	}()

	var initRet *starlarkInitRet
	select {
	case <-ctx.Done():
		thread.Cancel(ctx.Err().Error())
		initRet = <-done
	case initRet = <-done:
	}

	return initRet.globals, initRet.err
}

func (s *script) fromStarlarkValue(starlarkValue starlark.Value, obj runtime.Object) error {
	objJSON, err := s.starlarkCodec.encode(starlarkValue)
	if err != nil {
		return fmt.Errorf("encode starlark value: %w", err)
	}

	wantGVK := plugininternal.SchemeGroupVersion.WithKind("ActionRun")
	_, gotGVK, err := s.codec.Decode([]byte(objJSON), nil, obj)
	if err != nil {
		return fmt.Errorf("decode ActionRun: %w", err)
	}

	if wantGVK != *gotGVK {
		return fmt.Errorf("decode ActionRun: want %s got %s", wantGVK, gotGVK)
	}

	return nil
}

func getJSONEncoder(codecs serializer.CodecFactory) (runtime.Encoder, error) {
	serializerInfos := codecs.WithoutConversion().SupportedMediaTypes()
	for _, serializerInfo := range serializerInfos {
		if serializerInfo.MediaType == "application/json" {
			return serializerInfo.Serializer, nil
		}
	}
	return nil, fmt.Errorf("cannot find json serializer in %v", serializerInfos)
}

type starlarkCodec struct {
	encodeBuiltin, decodeBuiltin *starlark.Builtin
}

func newStarlarkCodec() (*starlarkCodec, error) {
	starlarkJSONEncodeValue, ok := starlarkjson.Module.Members["encode"]
	if !ok {
		return nil, fmt.Errorf("cannot find encode in json module")
	}
	starlarkJSONEncode, ok := starlarkJSONEncodeValue.(*starlark.Builtin)
	if !ok {
		return nil, fmt.Errorf("cannot cast starlark json.encode builtin (got %T)", starlarkJSONEncodeValue)
	}

	starlarkJSONDecodeValue, ok := starlarkjson.Module.Members["decode"]
	if !ok {
		return nil, fmt.Errorf("cannot find decode in json module")
	}
	starlarkJSONDecode, ok := starlarkJSONDecodeValue.(*starlark.Builtin)
	if !ok {
		return nil, fmt.Errorf("cannot cast starlark json.decode builtin (got %T)", starlarkJSONDecodeValue)
	}

	return &starlarkCodec{
		encodeBuiltin: starlarkJSONEncode,
		decodeBuiltin: starlarkJSONDecode,
	}, nil
}

func (c *starlarkCodec) encode(starlarkValue starlark.Value) (string, error) {
	thread := starlark.Thread{
		Name: "<actionscript:starlarkcodec:encode>",
		Print: func(thread *starlark.Thread, msg string) {
			klog.Info(thread.Name, "message", msg)
		},
	}
	starlarkValue, err := c.encodeBuiltin.CallInternal(
		&thread,
		[]starlark.Value{starlarkValue},
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("call starlark json.decode builtin: %w", err)
	}

	jsonString, ok := starlarkValue.(starlark.String)
	if !ok {
		return "", fmt.Errorf("expected starlark string return value, got %T", starlarkValue)
	}

	return string(jsonString), nil
}

func (c *starlarkCodec) decode(jsonString string) (starlark.Value, error) {
	// decoder := json.NewDecoder(bytes.NewBuffer([]byte(jsonString)))
	// var (
	// 	starlarkListElems []starlark.Value
	// 	starlarkDictItems []starlark.Tuple
	// )
	// for {
	// 	token, err := decoder.Token()
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			if starlarkListElems != nil || starlarkDictItems != nil {
	// 				return nil, fmt.Errorf("unfinished list/map json at end of %q", jsonString)
	// 			}
	// 			break
	// 		}
	// 		return nil, fmt.Errorf("invalid json at offset %d of %q", decoder.InputOffset(), jsonString)
	// 	}
	// 	switch token.(type) {
	// 	case json.Delim:
	// 		switch token {
	// 		case '[':
	// 			starlarkListElems = []starlark.Value{}
	// 		case ']':
	// 			starlark.NewList(starlarkListElems)
	// 			starlarkListElems = nil
	// 	}
	// }

	thread := starlark.Thread{
		Name: "<actionscript:starlarkcodec:decode>",
		Print: func(thread *starlark.Thread, msg string) {
			klog.Info(thread.Name, "message", msg)
		},
	}
	starlarkValue, err := c.decodeBuiltin.CallInternal(
		&thread,
		[]starlark.Value{starlark.String(string(jsonString))},
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("call starlark json.decode builtin: %w", err)
	}

	starlarkValue, err = convertDictsToStructs(starlarkValue)
	if err != nil {
		return nil, fmt.Errorf("convert script dict to struct: %w", err)
	}

	return starlarkValue, nil
}

func convertDictsToStructs(starlarkValue starlark.Value) (starlark.Value, error) {
	starlarkAttrsValue, convertible := starlarkValue.(starlark.IterableMapping)
	if !convertible {
		return starlarkValue, nil
	}

	stringDict := starlark.StringDict{}
	var err error
	for _, kv := range starlarkAttrsValue.Items() {
		k, v := kv[0], kv[1]
		if _, ok := k.(starlark.String); !ok {
			return nil, fmt.Errorf("got non-string key %q with type %T", k.String(), k)
		}
		kString := k.String()
		if len(kString) > 2 {
			// get rid of starting and ending quotes
			kString = kString[1 : len(kString)-1]
		}
		stringDict[kString], err = convertDictsToStructs(v)
		if err != nil {
			return nil, err
		}
	}

	return starlarkstruct.FromStringDict(starlarkstruct.Default, stringDict), nil
}

func copyProgram(program *starlark.Program) (*starlark.Program, error) {
	pipeReader, pipeWriter := io.Pipe()

	if err := program.Write(pipeWriter); err != nil {
		return nil, fmt.Errorf("program writer: %w", err)
	}
	if err := pipeWriter.Close(); err != nil {
		return nil, fmt.Errorf("close write pipe end: %w", err)
	}

	copiedProgram, err := starlark.CompiledProgram(pipeReader)
	if err != nil {
		return nil, fmt.Errorf("compiled program: %w", err)
	}
	if err := pipeReader.Close(); err != nil {
		return nil, fmt.Errorf("close read pipe end: %w", err)
	}

	return copiedProgram, nil
}

func getPredeclared(from, to *spiritsinternal.Spirit) starlark.StringDict {
	infoFunc := func(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
		if len(args) == 0 {
			return nil, nil
		}

		arg0 := args[0]
		log := klog.V(0)
		if arg0Int, ok := arg0.(starlark.Int); ok {
			log = klog.V(klog.Level(arg0Int.BigInt().Int64()))
			args = args[1:]
		}
		log.InfoS("action script", "thread", thread.Name, args)

		return nil, nil
	}

	errorFunc := func(thread *starlark.Thread, builtin *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
		reason := fmt.Sprintf("error() called at\n%s" + thread.CallStack().String())
		if len(args) > 0 {
			reason = args.String()
		}
		thread.Cancel(reason)
		return nil, nil
	}

	return starlark.StringDict{
		"info":  starlark.NewBuiltin("info", infoFunc),
		"error": starlark.NewBuiltin("error", errorFunc),
		"from":  toStarlarkStruct(from),
		"to":    toStarlarkStruct(from),
	}
}

func toStarlarkStruct(obj any) *starlarkstruct.Struct {
	starlarkDict := starlark.StringDict{}

	objType := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)

	if objType == nil {
		return nil
	}

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		if !field.IsExported() {
			continue
		}
		starlarkDict[field.Name] = toStarlarkValue(field.Type, objVal.Field(i))
	}
	return starlarkstruct.FromStringDict(starlarkstruct.Default, starlarkDict)
}

func toStarlarkValue(objType reflect.Type, objVal reflect.Value) starlark.Value {
	if objType == nil {
		return starlark.None
	}

	switch objType.Kind() {
	case reflect.Bool:
		return starlark.Bool(objVal.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return starlark.MakeInt64(objVal.Int())
	case reflect.Float32, reflect.Float64:
		return starlark.Float(objVal.Float())
	case reflect.Array:
		return toStarlarkList(objType.Elem(), objVal)
	case reflect.Map:
		return toStarlarkDict(objType.Key(), objType.Elem(), objVal)
	case reflect.Pointer:
		return toStarlarkValue(objType.Elem(), objVal.Elem())
	case reflect.Slice:
		return toStarlarkList(objType.Elem(), objVal)
	case reflect.String:
		return starlark.String(objVal.String())
	case reflect.Struct:
		return toStarlarkStruct(objVal.Interface())
	default:
		// TODO: probably shouldn't panic here
		panic("unsupported starlark value type: " + objType.Kind().String())
	}
}

func toStarlarkList(objType reflect.Type, objVal reflect.Value) *starlark.List {
	var starlarkValues []starlark.Value
	for i := 0; i < objVal.Len(); i++ {
		starlarkValues = append(starlarkValues, toStarlarkValue(objType, objVal.Index(i)))
	}
	return starlark.NewList(starlarkValues)
}

func toStarlarkDict(keyType, valType reflect.Type, objVal reflect.Value) *starlark.Dict {
	starlarkDict := starlark.NewDict(objVal.Len())
	objMapRange := objVal.MapRange()
	for objMapRange.Next() {
		keyVal := objMapRange.Key()
		valVal := objMapRange.Value()
		starlarkDict.SetKey(toStarlarkValue(keyType, keyVal), toStarlarkValue(valType, valVal))
	}
	return starlarkDict
}

func fromStarlarkValue(starlarkValue starlark.Value, obj any) error {
	type mappingValue struct {
		goTypeKindMask reflect.Kind
		mapFunc        func(starlark.Value, reflect.Value) error
	}
	mappings := map[reflect.Type]mappingValue{
		reflect.TypeOf(starlark.Bool(false)): {
			goTypeKindMask: reflect.Bool,
			mapFunc: func(starlarkValue starlark.Value, objVal reflect.Value) error {
				objVal.SetBool(bool(starlarkValue.(starlark.Bool)))
				return nil
			},
		},
		reflect.TypeOf(starlark.Dict{}): {
			goTypeKindMask: reflect.Map,
			mapFunc: func(starlarkValue starlark.Value, objVal reflect.Value) error {
				starlarkDict := starlarkValue.(*starlark.Dict)
				for _, kv := range starlarkDict.Items() {
					starlarkMapKey, starlarkMapVal := kv[0], kv[1]
					goMapKeyObj, goMapValObj := reflect.New(objVal.Type().Key()).Interface(), reflect.New(objVal.Type().Elem()).Interface()
					if err := fromStarlarkValue(starlarkMapKey, goMapKeyObj); err != nil {
						return fmt.Errorf("from starlark map key: %w", err)
					}
					if err := fromStarlarkValue(starlarkMapVal, goMapValObj); err != nil {
						return fmt.Errorf("from starlark map value: %w", err)
					}
				}
				return nil
			},
		},
		reflect.TypeOf(starlark.Float(0)): {
			goTypeKindMask: reflect.Float32 | reflect.Float64,
			mapFunc: func(starlarkValue starlark.Value, objVal reflect.Value) error {
				// TODO: handle underflow/overflow
				objVal.SetFloat(float64(starlarkValue.(starlark.Float)))
				return nil
			},
		},
		reflect.TypeOf(starlark.Int{}): {
			goTypeKindMask: reflect.Int | reflect.Int8 | reflect.Int16 | reflect.Int32 | reflect.Int64 |
				reflect.Uint | reflect.Uint8 | reflect.Uint16 | reflect.Uint32 | reflect.Uint64,
			mapFunc: func(starlarkValue starlark.Value, objVal reflect.Value) error {
				// TODO: handle underflow/overflow
				objVal.SetInt(starlarkValue.(starlark.Int).BigInt().Int64())
				return nil
			},
		},
		reflect.TypeOf(starlark.List{}): {
			goTypeKindMask: reflect.Array | reflect.Slice,
			mapFunc: func(starlarkValue starlark.Value, objVal reflect.Value) error {
				starlarkList := starlarkValue.(*starlark.List)
				for i := 0; i < starlarkList.Len(); i++ {
					starlarkValue := starlarkList.Index(i)
					goValue := reflect.New(objVal.Type().Elem()).Interface()
					if err := fromStarlarkValue(starlarkValue, goValue); err != nil {
						return fmt.Errorf("from starlark list value: %w", err)
					}
				}
				return nil
			},
		},
	}
	mapping, ok := mappings[reflect.TypeOf(starlarkValue)]
	if !ok {
		return fmt.Errorf("unknown mapping for skylark type %T", starlarkValue)
	}

	if reflect.TypeOf(obj).Kind()&mapping.goTypeKindMask != 0 {
		return fmt.Errorf("can't convert starlark value type %T to go value type %T", starlarkValue, obj)
	}

	if err := mapping.mapFunc(starlarkValue, reflect.ValueOf(obj)); err != nil {
		return fmt.Errorf("can't map starlark value %q to go value", starlarkValue.String())
	}

	return nil
}

func fromStarlarkStruct(starlarkStruct *starlarkstruct.Struct, obj any) any {
	return nil
}

package script

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"

	"github.com/ankeesler/spirits0/internal/api"
	"github.com/ankeesler/spirits0/internal/spirit"
	fuzz "github.com/google/gofuzz"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

const ctxKey = "action.ctx"

var spiritWithAllFieldsSet = spirit.Spirit{}

func init() {
	fuzz.
		New().
		RandSource(rand.NewSource(0)).
		NilChance(0).
		NumElements(1, 1).
		MaxDepth(10).
		Funcs(func(meta **api.Meta, c fuzz.Continue) {
			*meta = &api.Meta{}
			c.Fuzz(*meta)
		}).
		Funcs(func(stats **api.SpiritStats, c fuzz.Continue) {
			*stats = &api.SpiritStats{}
			c.Fuzz(*stats)
		}).
		Funcs(func(action **api.SpiritAction, c fuzz.Continue) {
			*action = &api.SpiritAction{}
			c.Fuzz(*action)
		}).
		// Skip fuzzing Action implementation - the script shouldn't touch this field
		Funcs(func(_ *spirit.Action, c fuzz.Continue) {}).
		Fuzz(&spiritWithAllFieldsSet)
	log.Printf("Spirit with all fields set: %+v", spiritWithAllFieldsSet.API)
}

type Script struct {
	program *starlark.Program
}

func Compile(source string) (*Script, error) {
	s := &Script{}

	predeclared, err := newPredeclared(
		&spiritWithAllFieldsSet,
		[]*spirit.Spirit{&spiritWithAllFieldsSet},
	)
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

func (s *Script) Run(ctx context.Context, source *spirit.Spirit, targets []*spirit.Spirit) error {
	out := bytes.NewBuffer([]byte{})
	thread := &starlark.Thread{
		Name: "<actionscript:load>",
		Print: func(thread *starlark.Thread, msg string) {
			msg = fmt.Sprintf("%s: %s", thread.Name, msg)
			fmt.Fprintf(out, msg)
			log.Printf(msg)
		},
	}
	predeclared, err := newPredeclared(source, targets)
	if err != nil {
		return fmt.Errorf("get script predeclared symbols for run: %w", err)
	}

	if err := s.run(ctx, thread, predeclared); err != nil {
		return fmt.Errorf("run script: %w (out: %q)", err, out.String())
	}

	return nil
}

func (s *Script) run(
	ctx context.Context,
	thread *starlark.Thread,
	predeclared starlark.StringDict,
) error {
	type starlarkInitRet struct {
		globals starlark.StringDict
		err     error
	}
	done := make(chan *starlarkInitRet)
	defer close(done)

	thread.SetLocal(ctxKey, ctx)
	go func() {
		globals, err := s.program.Init(thread, predeclared)
		if err != nil {
			err = fmt.Errorf("script failed: %w (%s, %s)", err, thread.Local("resolve"), thread.Local("reject"))
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

	return initRet.err
}

func newPredeclared(
	source *spirit.Spirit,
	targets []*spirit.Spirit,
) (starlark.StringDict, error) {
	starlarkCtxStruct := starlarkstruct.FromStringDict(starlarkstruct.Default, starlark.StringDict{
		"value": starlark.NewBuiltin(
			"ctx.value",
			func(
				thread *starlark.Thread,
				b *starlark.Builtin,
				args starlark.Tuple,
				kwargs []starlark.Tuple,
			) (starlark.Value, error) {
				var key starlark.Value
				if err := starlark.UnpackPositionalArgs(b.Name(), args, kwargs, 1, &key); err != nil {
					return nil, err
				}

				ctx, ok := thread.Local(ctxKey).(context.Context)
				if !ok {
					return nil, errors.New("missing thread local context")
				}

				val := ctx.Value(key)
				if val == nil {
					return starlark.None, nil
				}

				return starlark.String(val.(string)), nil
			},
		),
		"set_value": starlark.NewBuiltin(
			"ctx.value",
			func(
				thread *starlark.Thread,
				b *starlark.Builtin,
				args starlark.Tuple,
				kwargs []starlark.Tuple,
			) (starlark.Value, error) {
				var key, val starlark.Value
				if err := starlark.UnpackPositionalArgs(b.Name(), args, kwargs, 2, &key, &val); err != nil {
					return nil, err
				}
				return val, nil
			},
		),
	})

	var starlarkTargets []starlark.Value
	for _, target := range targets {
		starlarkTargets = append(starlarkTargets, newSpiritStarlarkStruct(target.API))
	}

	starlarkAbortFunc := starlark.NewBuiltin(
		"abort",
		func(
			thread *starlark.Thread,
			b *starlark.Builtin,
			args starlark.Tuple,
			kwargs []starlark.Tuple,
		) (starlark.Value, error) {
			thread.Cancel(fmt.Sprintf("abort(%v, %v)", args, kwargs))
			return starlark.None, nil
		},
	)

	starlarkActionStruct := starlarkstruct.FromStringDict(starlarkstruct.Default, starlark.StringDict{
		"ctx":     starlarkCtxStruct,
		"source":  newSpiritStarlarkStruct(source.API),
		"targets": starlark.NewList(starlarkTargets),
		"abort":   starlarkAbortFunc,
	})

	return starlark.StringDict{
		"action": starlarkActionStruct,
	}, nil
}

func newSpiritStarlarkStruct(spirit *api.Spirit) *starlarkstruct.Struct {
	starlarkMetaStruct := starlarkstruct.FromStringDict(starlarkstruct.Default, starlark.StringDict{
		"id":           starlark.String(spirit.GetMeta().GetId()),
		"created_time": starlark.MakeInt64(spirit.GetMeta().GetCreatedTime().AsTime().Unix()),
		"updated_time": starlark.MakeInt64(spirit.GetMeta().GetUpdatedTime().AsTime().Unix()),
	})

	starlarkStatsDict := starlark.StringDict{}
	addStatStarlarkBuitlins(
		starlarkStatsDict, "health", &spirit.GetStats().Health)
	addStatStarlarkBuitlins(
		starlarkStatsDict, "physical_power", &spirit.GetStats().PhysicalPower)
	addStatStarlarkBuitlins(
		starlarkStatsDict, "physical_constitution", &spirit.GetStats().PhysicalConstitution)
	addStatStarlarkBuitlins(
		starlarkStatsDict, "mental_power", &spirit.GetStats().MentalPower)
	addStatStarlarkBuitlins(
		starlarkStatsDict, "mental_constitution", &spirit.GetStats().MentalConstitution)
	addStatStarlarkBuitlins(
		starlarkStatsDict, "agility", &spirit.GetStats().Agility)
	starlarkStatsStruct := starlarkstruct.FromStringDict(starlarkstruct.Default, starlarkStatsDict)

	var starlarkActionsList []starlark.Value
	for _, action := range spirit.GetActions() {
		starlarkActionStruct := starlarkstruct.FromStringDict(starlarkstruct.Default, starlark.StringDict{
			"name":      starlark.String(action.GetName()),
			"action_id": starlark.String(action.GetActionId()),
		})
		starlarkActionsList = append(starlarkActionsList, starlarkActionStruct)
	}
	starlarkActionList := starlark.NewList(starlarkActionsList)

	return starlarkstruct.FromStringDict(starlarkstruct.Default, starlark.StringDict{
		"meta":    starlarkMetaStruct,
		"stats":   starlarkStatsStruct,
		"actions": starlarkActionList,
	})
}

func addStatStarlarkBuitlins(
	starlarkDict starlark.StringDict,
	statName string,
	stat *int64,
) {
	getter := fmt.Sprintf("%s", statName)
	setter := fmt.Sprintf("set_%s", statName)

	starlarkDict[getter] = starlark.NewBuiltin(
		getter,
		func(
			thread *starlark.Thread,
			b *starlark.Builtin,
			args starlark.Tuple,
			kwargs []starlark.Tuple,
		) (starlark.Value, error) {
			return starlark.MakeInt64(*stat), nil
		},
	)

	starlarkDict[setter] = starlark.NewBuiltin(
		setter,
		func(
			thread *starlark.Thread,
			b *starlark.Builtin,
			args starlark.Tuple,
			kwargs []starlark.Tuple,
		) (starlark.Value, error) {
			var newStat int64
			if err := starlark.UnpackPositionalArgs(b.Name(), args, kwargs, 1, &newStat); err != nil {
				return nil, err
			}
			*stat = newStat
			return starlark.MakeInt64(*stat), nil
		},
	)
}

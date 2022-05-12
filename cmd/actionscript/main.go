package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

	actioninternal "github.com/ankeesler/spirits/internal/action"
	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
	plugininternal "github.com/ankeesler/spirits/internal/apis/spirits/plugin"
	pluginv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/plugin/v1alpha1"
	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
)

var (
	scheme = runtime.NewScheme()
	codecs = serializer.NewCodecFactory(scheme)
)

func init() {
	utilruntime.Must(spiritsinternal.AddToScheme(scheme))
	utilruntime.Must(spiritsv1alpha1.AddToScheme(scheme))
	utilruntime.Must(plugininternal.AddToScheme(scheme))
	utilruntime.Must(pluginv1alpha1.AddToScheme(scheme))
}

func main() {
	if err := reallyMain(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err.Error())
		os.Exit(1)
	}
}

func reallyMain() error {
	var (
		contextPath string
		scriptPath  string
		help        bool
	)
	flag.StringVar(&contextPath, "context", "context.json", "path to file containing ActionRun JSON")
	flag.StringVar(&scriptPath, "script", "script.star", "path to file containing actionscript")
	flag.BoolVar(&help, "help", false, "print this help message")
	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(1)
	}

	actionRun, err := readActionRun(contextPath)
	if err != nil {
		return fmt.Errorf("read ActionRun: %w", err)
	}

	script, err := os.ReadFile(scriptPath)
	if err != nil {
		return fmt.Errorf("read script: %w", err)
	}

	actionRunGV, err := schema.ParseGroupVersion(actionRun.APIVersion)
	if err != nil {
		return fmt.Errorf("parse APIVersion: %w", err)
	}

	encoder, err := getJSONEncoder()
	codec := codecs.CodecForVersions(encoder, codecs.UniversalDecoder(), actionRunGV, nil)

	action, err := actioninternal.Script(codec, string(script))
	if err != nil {
		return fmt.Errorf("compile script: %w", err)
	}

	from := &spiritsinternal.Spirit{Spec: actionRun.Spec.From}
	to := &spiritsinternal.Spirit{Spec: actionRun.Spec.To}
	if err := action.Run(context.Background(), from, to); err != nil {
		return fmt.Errorf("run: %w", err)
	}

	printSpirit("from", from)
	printSpirit("to", to)

	return nil
}

func readActionRun(path string) (*plugininternal.ActionRun, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	obj, err := runtime.Decode(codecs.UniversalDecoder(), data)
	if err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}

	var actionRun plugininternal.ActionRun
	if err := scheme.Convert(obj, &actionRun, nil); err != nil {
		return nil, fmt.Errorf("convert: %w", err)
	}

	return obj.(*plugininternal.ActionRun), nil
}

func getJSONEncoder() (runtime.Encoder, error) {
	serializerInfos := codecs.WithoutConversion().SupportedMediaTypes()
	for _, serializerInfo := range serializerInfos {
		if serializerInfo.MediaType == "application/json" {
			return serializerInfo.Serializer, nil
		}
	}
	return nil, fmt.Errorf("cannot find json serializer in %v", serializerInfos)
}

func printSpirit(name string, spirit *spiritsinternal.Spirit) {
	fmt.Println(name)
	fmt.Println(spirit.Spec.Attributes.Stats)
}

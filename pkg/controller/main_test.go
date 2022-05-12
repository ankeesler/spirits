package controller

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

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

package main

import (
	"log"

	inputinternal "github.com/ankeesler/spirits/internal/apis/spirits/input"
	inputv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/input/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

func main() {
	scheme := runtime.NewScheme()
	// metav1.AddToGroupVersion(scheme, metav1.Unversioned)
	schemeBuilder := runtime.NewSchemeBuilder(
		inputv1alpha1.AddToScheme,
		inputinternal.AddToScheme,
	)
	utilruntime.Must(schemeBuilder.AddToScheme(scheme))
	log.Printf("%#v", scheme)

	log.Print(scheme.ObjectKinds(&inputv1alpha1.ActionCall{}))
}

package apiserver

import (
	"context"
	"fmt"
	"net"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	apiserverrest "k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/apiserver/pkg/server/dynamiccertificates"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	"k8s.io/client-go/rest"
	certutil "k8s.io/client-go/util/cert"
	"k8s.io/klog/v2"

	inputinternal "github.com/ankeesler/spirits/internal/apis/spirits/input"
	inputv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/input/v1alpha1"
)

type ActionSink interface {
	Post(namespace, battleName, spiritName, actionName string) error
}

type APIServer struct {
	// Required

	Port    int
	DNSName string

	ActionSink ActionSink

	// Optional

	PostStartHook func() error
}

func (m *APIServer) Start(ctx context.Context) error {
	scheme := getScheme()
	codecFactory := serializer.NewCodecFactory(scheme)
	config, err := m.getConfig(codecFactory)
	if err != nil {
		return fmt.Errorf("get apiserver config: %w", err)
	}

	apiServer, err := config.Complete().New("spirits-apiserver", genericapiserver.NewEmptyDelegate())
	if err != nil {
		return fmt.Errorf("new apiserver: %w", err)
	}

	actionCallGVR := inputv1alpha1.SchemeGroupVersion.WithResource("actioncalls")
	apiGroup := genericapiserver.NewDefaultAPIGroupInfo(actionCallGVR.Group, scheme, metav1.ParameterCodec, codecFactory)
	apiGroup.VersionedResourcesStorageMap[actionCallGVR.Version] = map[string]apiserverrest.Storage{
		actionCallGVR.Resource: &actionCallHandler{ActionSink: m.ActionSink},
	}
	if err := apiServer.InstallAPIGroup(&apiGroup); err != nil {
		return fmt.Errorf("install api group: %w", err)
	}

	if m.PostStartHook != nil {
		if err := apiServer.AddPostStartHook(
			"aggregated-api-manager-post-start-hook",
			func(postStartContext genericapiserver.PostStartHookContext) error {
				return m.PostStartHook()
			},
		); err != nil {
			return fmt.Errorf("add post start hook")
		}
	}

	return apiServer.PrepareRun().Run(ctx.Done())
}

func (m *APIServer) getConfig(codecFactory serializer.CodecFactory) (*genericapiserver.RecommendedConfig, error) {
	// Create new recommended set of apiserver options
	options := genericoptions.NewRecommendedOptions("", codecFactory.LegacyCodec(inputv1alpha1.SchemeGroupVersion))

	// Set serving port
	options.SecureServing.BindPort = m.Port

	// Support for running outside of a cluster (e.g., debugging on a local machine)
	maybeSetKubeconfigPath(options)

	// Setup self-signed certs for the apiserver
	// #182087996: this is obviously not ok, need to add real apiserver certs support
	cert, key, err := certutil.GenerateSelfSignedCertKey("xxx", []net.IP{}, []string{m.DNSName})
	if err != nil {
		return nil, fmt.Errorf("generate cert: %w", err)
	}
	dynamicCert, err := dynamiccertificates.NewStaticCertKeyContent("xxx", cert, key)
	if err != nil {
		return nil, fmt.Errorf("create dynamic cert: %w", err)
	}
	options.SecureServing.SecureServingOptions.ServerCert.GeneratedCert = dynamicCert

	// Don't need to talk to etcd - our handling is done in memory
	options.Etcd = nil

	// Apply options to the recommended apiserver config
	config := genericapiserver.NewRecommendedConfig(codecFactory)
	if err := options.ApplyTo(config); err != nil {
		return nil, fmt.Errorf("apply options to apiserver config: %w", err)
	}

	return config, nil
}

func getScheme() *runtime.Scheme {
	scheme := runtime.NewScheme()
	metav1.AddToGroupVersion(scheme, metav1.Unversioned)
	schemeBuilder := runtime.NewSchemeBuilder(
		inputv1alpha1.AddToScheme,
		inputinternal.AddToScheme,
	)
	utilruntime.Must(schemeBuilder.AddToScheme(scheme))
	return scheme
}

func maybeSetKubeconfigPath(options *genericoptions.RecommendedOptions) {
	// Assume that if we can get an in-cluster rest.Config, then we don't need to set a kubeconfig path
	_, err := rest.InClusterConfig()
	if err == nil {
		return
	}

	path := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	klog.InfoS("using kubeconfig path for spirits apiserver", "path", path)
	options.Authentication.RemoteKubeConfigFile = path
	options.Authorization.RemoteKubeConfigFile = path
	options.CoreAPI.CoreAPIKubeconfigPath = path
}

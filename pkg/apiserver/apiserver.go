package apiserver

import (
	"context"
	"fmt"
	"net"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"

	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
)

type ActionSink interface {
	Post(namespace, battleName, battleGeneration, spiritName, spiritGeneration, actionName string) error
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

	actionRequestGVR := spiritsv1alpha1.SchemeGroupVersion.WithResource("actionrequests")
	apiGroup := genericapiserver.NewDefaultAPIGroupInfo(actionRequestGVR.Group, scheme, metav1.ParameterCodec, codecFactory)
	apiGroup.VersionedResourcesStorageMap[actionRequestGVR.Version][actionRequestGVR.Resource] = &actionRequestHandler{ActionSink: m.ActionSink}
	if err := apiServer.InstallAPIGroup(&apiGroup); err != nil {
		return fmt.Errorf("install api group")
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
	options := genericoptions.NewRecommendedOptions("", codecFactory.LegacyCodec(spiritsv1alpha1.SchemeGroupVersion))

	// Set serving port
	options.SecureServing.BindPort = m.Port

	// Setup self-signed certs for the apiserver
	if err := options.SecureServing.MaybeDefaultWithSelfSignedCerts(
		"spirits-apiserver",
		[]string{m.DNSName},
		[]net.IP{},
	); err != nil {
		return nil, fmt.Errorf("default with signed certs: %w", err)
	}

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
	scheme.AddKnownTypes(
		spiritsv1alpha1.SchemeGroupVersion,
		&spiritsv1alpha1.ActionRequest{},
		&spiritsv1alpha1.ActionRequestList{},
	)
	metav1.AddToGroupVersion(scheme, spiritsv1alpha1.SchemeGroupVersion)
	return scheme
}

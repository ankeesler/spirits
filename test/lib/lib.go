package lib

import (
	"context"
	"testing"
	"time"

	spiritsclientset "github.com/ankeesler/spirits/pkg/apis/clientset/versioned"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type TestConfig struct {
	Ctx              context.Context
	CoreClientset    kubernetes.Interface
	SpiritsClientset spiritsclientset.Interface
	Namespace        *corev1.Namespace
}

func NewTestConfig(t *testing.T) *TestConfig {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// Load kubeconfig
	loader := clientcmd.NewDefaultClientConfigLoadingRules()
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loader, &clientcmd.ConfigOverrides{})
	config, err := clientConfig.ClientConfig()
	require.NoError(t, err)

	// Create clients
	coreClientset, err := kubernetes.NewForConfig(config)
	require.NoError(t, err)

	spiritsClientset, err := spiritsclientset.NewForConfig(config)
	require.NoError(t, err)

	// Create namespace
	namespace := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{GenerateName: "spirits-api-test-"}}
	namespace, err = coreClientset.CoreV1().Namespaces().Create(ctx, namespace, metav1.CreateOptions{})
	require.NoError(t, err)

	// Set global config
	return &TestConfig{
		Ctx:             ctx,
		CoreClientset:   coreClientset,
		SpritsClientset: spiritsClientset,
		Namespace:       namespace,
	}
}

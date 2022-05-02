package api

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"

	spiritsclientset "github.com/ankeesler/spirits/pkg/apis/clientset/versioned"
	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
)

var (
	scheme = runtime.NewScheme()
	codecs = serializer.NewCodecFactory(scheme)

	tc = struct {
		ctx              context.Context
		coreClientset    kubernetes.Interface
		spiritsClientset spiritsclientset.Interface
		namespace        *corev1.Namespace
	}{}
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(spiritsv1alpha1.AddToScheme(scheme))
}

func TestMain(m *testing.M) {
	if _, ok := os.LookupEnv("SPIRITS_TEST_INTEGRATION"); !ok {
		alert(false, "skipping integration tests because env var 'SPIRITS_TEST_INTEGRATION' not set")
		os.Exit(0)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	if err := setTestConfig(ctx); err != nil {
		alert(true, err.Error())
		os.Exit(1)
	}

	exitCode := m.Run()

	if err := tc.coreClientset.CoreV1().Namespaces().Delete(ctx, tc.namespace.Name, metav1.DeleteOptions{}); err != nil {
		alert(false, fmt.Errorf("delete namespace: %w", err).Error())
	}

	os.Exit(exitCode)
}

func alert(isError bool, message string) {
	level := "WARNING"
	if isError {
		level = "ERROR"
	}
	fmt.Println(strings.Repeat("!", 80))
	fmt.Printf("  %s\n", level)
	fmt.Printf("    %s\n", message)
	fmt.Printf("  %s\n", level)
	fmt.Println(strings.Repeat("!", 80))
}

func setTestConfig(ctx context.Context) error {
	// Load kubeconfig
	loader := clientcmd.NewDefaultClientConfigLoadingRules()
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loader, &clientcmd.ConfigOverrides{})
	config, err := clientConfig.ClientConfig()
	if err != nil {
		return fmt.Errorf("get client config: %w", err)
	}

	// Create clients
	coreClientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("create kubernetes clientset: %w", err)
	}
	spiritsClientset, err := spiritsclientset.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("create spirits clientset: %w", err)
	}

	// Create namespace
	namespace := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{GenerateName: "spirits-api-test-"}}
	namespace, err = coreClientset.CoreV1().Namespaces().Create(ctx, namespace, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("create test namespace: %w", err)
	}

	// Set global config
	tc.ctx = ctx
	tc.coreClientset = coreClientset
	tc.spiritsClientset = spiritsClientset
	tc.namespace = namespace

	return nil
}

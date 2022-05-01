package api

import (
	"context"
	"os"
	"path/filepath"
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

	spiritsv1alpha1clientset "github.com/ankeesler/spirits/pkg/apis/clientset/versioned"
	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
	"github.com/stretchr/testify/require"
)

var (
	scheme = runtime.NewScheme()
	codecs = serializer.NewCodecFactory(scheme)
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(spiritsv1alpha1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

func TestAPI(t *testing.T) {
	ctx, cancel := context.WithTimeout(time.Minute * 30)
	defer cancel()

	// Read test fixtures
	spiritA := readObject(t, "spirit-a.yaml").(*spiritsv1alpha1.Spirit)
	spiritB := readObject(t, "spirit-a.yaml").(*spiritsv1alpha1.Spirit)
	theBattle := readObject(t, "the-battle.yaml").(*spiritsv1alpha1.Battle)

	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(os.Getenv("HOME"), ".kube", "config"))
	require.NoError(t, err)

	coreClientset, err := kubernetes.NewForConfig(config)
	require.NoError(t, err)

	spiritsClientset, err := spiritsv1alpha1clientset.NewForConfig(config)
	require.NoError(t, err)

	namespace := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{GenerateName: "spirits-api-test-"}}
	namespace, err = coreClientset.CoreV1().Namespaces().Create(ctx, namespace, metav1.CreateOptions{})
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := coreClientset.CoreV1().Namespaces().Delete(ctx, namespace.Name, metav1.DeleteOptions{}); err != nil {
			t.Log("could not delete namespace:", err.Error())
		}
	})

	spiritA, err = spiritsClientset.AnkeeslerV1alpha1().Spirits(namespace.Name).Create(ctx, spiritA, metav1.CreateOptions{})
	require.NoError(t, err)

	spiritB, err = spiritsClientset.AnkeeslerV1alpha1().Spirits(namespace.Name).Create(ctx, spiritB, metav1.CreateOptions{})
	require.NoError(t, err)

	theBattle, err = spiritsClientset.AnkeeslerV1alpha1().Battles(namespace.Name).Create(ctx, theBattle, metav1.CreateOptions{})
	require.NoError(t, err)
}

func readObject(t *testing.T, path string) runtime.Object {
	t.Helper()

	data, err := os.ReadFile(filepath.Join("testdata", path))
	require.NoError(t, err)

	obj, _, err := codecs.UniversalDeserializer().Decode(data, nil, nil)
	require.NoError(t, err)

	return obj
}

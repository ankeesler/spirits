package api

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()

	// Wire clients
	coreClientset, spiritsClientset := createClients(t)

	// Create test namespace
	namespace := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{GenerateName: "spirits-api-test-"}}
	namespace, err := coreClientset.CoreV1().Namespaces().Create(ctx, namespace, metav1.CreateOptions{})
	require.NoError(t, err)
	t.Cleanup(func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := coreClientset.CoreV1().Namespaces().Delete(ctx, namespace.Name, metav1.DeleteOptions{}); err != nil {
			t.Log("could not delete namespace:", err.Error())
		}
	})

	// Read test fixtures
	var spirits []*spiritsv1alpha1.Spirit
	for _, spiritPath := range []string{"spirit-a.yaml", "spirit-b.yaml"} {
		spirits = append(spirits, readObject(t, spiritPath).(*spiritsv1alpha1.Spirit))
	}
	battle := readObject(t, "the-battle.yaml").(*spiritsv1alpha1.Battle)

	// Create test fixtures
	for i := range spirits {
		spirits[i], err = spiritsClientset.SpiritsV1alpha1().Spirits(namespace.Name).Create(ctx, spirits[i], metav1.CreateOptions{})
		require.NoError(t, err)
	}
	battle, err = spiritsClientset.SpiritsV1alpha1().Battles(namespace.Name).Create(ctx, battle, metav1.CreateOptions{})
	require.NoError(t, err)

	// Assert test fixtures are happy
	for i := range spirits {
		requireEventuallyConsistent(t, func() (bool, error) {
			spirit, err := spiritsClientset.SpiritsV1alpha1().Spirits(namespace.Name).Get(ctx, spirits[i].Name, metav1.GetOptions{})
			if err != nil {
				return false, fmt.Errorf("get: %w", err)
			}
			t.Logf("got spirit %q conditions: %#v", spirit.Name, spirit.Status.Conditions)
			return meta.IsStatusConditionTrue(spirit.Status.Conditions, "Ready"), nil
		})
	}
	requireEventuallyConsistent(t, func() (bool, error) {
		battle, err := spiritsClientset.SpiritsV1alpha1().Battles(namespace.Name).Get(ctx, battle.Name, metav1.GetOptions{})
		if err != nil {
			return false, fmt.Errorf("get: %w", err)
		}
		t.Logf("got battle %q conditions: %#v", battle.Name, battle.Status.Conditions)
		return meta.IsStatusConditionTrue(battle.Status.Conditions, "Progressing"), nil
	})

	// Assert battle finished
	requireEventuallyConsistent(t, func() (bool, error) {
		battle, err := spiritsClientset.SpiritsV1alpha1().Battles(namespace.Name).Get(ctx, battle.Name, metav1.GetOptions{})
		if err != nil {
			return false, fmt.Errorf("get: %w", err)
		}
		t.Logf("got battle %q phase: %q", battle.Name, battle.Status.Phase)
		return battle.Status.Phase == spiritsv1alpha1.BattlePhaseFinished, nil
	})

	// Assert spirit stats at end of battle
	battle, err = spiritsClientset.SpiritsV1alpha1().Battles(namespace.Name).Get(ctx, battle.Name, metav1.GetOptions{})
	require.NoError(t, err)
	gotInBattleSpiritsStats := getInBattleSpiritStats(t, ctx, spiritsClientset, battle)
	require.Len(t, gotInBattleSpiritsStats, 2)
	require.Equal(t, 0, gotInBattleSpiritsStats["the-battle-spirit-a-1"].Health)
	require.Equal(t, 1, gotInBattleSpiritsStats["the-battle-spirit-b-1"].Health)
}

func createClients(t *testing.T) (kubernetes.Interface, spiritsv1alpha1clientset.Interface) {
	t.Helper()

	// Load kubeconfig
	loader := clientcmd.NewDefaultClientConfigLoadingRules()
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loader, &clientcmd.ConfigOverrides{})
	config, err := clientConfig.ClientConfig()
	require.NoError(t, err)

	// Really create clients
	coreClientset, err := kubernetes.NewForConfig(config)
	require.NoError(t, err)

	spiritsClientset, err := spiritsv1alpha1clientset.NewForConfig(config)
	require.NoError(t, err)

	return coreClientset, spiritsClientset
}

func readObject(t *testing.T, path string) runtime.Object {
	t.Helper()

	data, err := os.ReadFile(filepath.Join("testdata", path))
	require.NoError(t, err)

	obj, _, err := codecs.UniversalDeserializer().Decode(data, nil, nil)
	require.NoError(t, err)

	return obj
}

func getInBattleSpiritStats(
	t *testing.T,
	ctx context.Context,
	spiritsClientset spiritsv1alpha1clientset.Interface,
	battle *spiritsv1alpha1.Battle,
) map[string]*spiritsv1alpha1.SpiritStats {
	inBattleSpirits := make(map[string]*spiritsv1alpha1.SpiritStats)
	for _, inBattleSpiritRef := range battle.Status.InBattleSpirits {
		inBattleSpirit, err := spiritsClientset.SpiritsV1alpha1().Spirits(battle.Namespace).Get(ctx, inBattleSpiritRef.Name, metav1.GetOptions{})
		require.NoError(t, err)
		inBattleSpirits[inBattleSpirit.Name] = &inBattleSpirit.Spec.Stats
	}
	return inBattleSpirits
}

func requireEventuallyConsistent(t *testing.T, conditionFunc func() (bool, error)) {
	t.Helper()

	// Wait for the condition to be met
	require.NoError(t, wait.PollImmediate(time.Second*1, time.Second*5, func() (bool, error) {
		return conditionFunc()
	}))

	// Make sure the condition stays consistent
	require.Equal(t, wait.ErrWaitTimeout, wait.PollImmediate(time.Second*1, time.Second*5, func() (bool, error) {
		met, err := conditionFunc()
		return !met, err
	}))
}

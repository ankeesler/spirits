package api

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"

	spiritsclientset "github.com/ankeesler/spirits/pkg/apis/clientset/versioned"
	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
)

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
	spiritsClientset spiritsclientset.Interface,
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

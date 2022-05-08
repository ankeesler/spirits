package server

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

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
		inBattleSpirits[inBattleSpirit.Name] = &inBattleSpirit.Spec.Attributes.Stats
	}
	return inBattleSpirits
}

func requireEventuallyConsistent(t *testing.T, conditionFunc func() (bool, error)) {
	t.Helper()

	const interval, duration = time.Second * 1, time.Second * 10

	// Wait until we see the conditionFunc return 3 times in a row with the same result
	deadline := time.Now().Add(duration)
	successesLeft := 3
	for time.Now().Before(deadline) {
		condition, err := conditionFunc()
		require.NoError(t, err)

		if condition {
			successesLeft--
			if successesLeft == 0 {
				return
			}
		} else {
			successesLeft = 3
		}

		time.Sleep(interval)
	}

	require.Fail(t, "condition failed to be eventually consistent")
}

func requireEventuallyConsistentSpirit(
	t *testing.T,
	ctx context.Context,
	name string,
	conditionFunc func(*spiritsv1alpha1.Spirit) bool,
) {
	t.Helper()
	requireEventuallyConsistent(t, func() (bool, error) {
		spirit, err := tc.spiritsClientset.SpiritsV1alpha1().Spirits(tc.namespace.Name).Get(ctx, name, metav1.GetOptions{})
		if err != nil {
			return false, fmt.Errorf("get: %w", err)
		}
		t.Logf("got spirit %q status: %+v", spirit.Name, spirit.Status)
		return conditionFunc(spirit), nil
	})
	t.Log("spirit is consistent")
}

func requireEventuallyConsistentBattle(
	t *testing.T,
	ctx context.Context,
	name string,
	conditionFunc func(*spiritsv1alpha1.Battle) bool,
) {
	t.Helper()
	requireEventuallyConsistent(t, func() (bool, error) {
		battle, err := tc.spiritsClientset.SpiritsV1alpha1().Battles(tc.namespace.Name).Get(ctx, name, metav1.GetOptions{})
		if err != nil {
			return false, fmt.Errorf("get: %w", err)
		}
		t.Logf("got battle %q status: %+v", battle.Name, battle.Status)
		return conditionFunc(battle), nil
	})
	t.Log("battle is consistent")
}

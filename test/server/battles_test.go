package server

import (
	"context"
	"fmt"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
	"github.com/stretchr/testify/require"
)

func TestNonHumanIntelligenceBattle(t *testing.T) {
	ctx, cancel := context.WithTimeout(tc.ctx, time.Minute)
	defer cancel()

	// Read test fixtures
	var spirits []*spiritsv1alpha1.Spirit
	for _, spiritPath := range []string{"spirit-a.yaml", "spirit-b.yaml"} {
		spirits = append(spirits, readObject(t, spiritPath).(*spiritsv1alpha1.Spirit))
	}
	battle := readObject(t, "the-battle.yaml").(*spiritsv1alpha1.Battle)

	// Create test fixtures
	var err error
	for i := range spirits {
		spirits[i], err = tc.spiritsClientset.SpiritsV1alpha1().Spirits(tc.namespace.Name).Create(ctx, spirits[i], metav1.CreateOptions{})
		require.NoError(t, err)
	}
	battle, err = tc.spiritsClientset.SpiritsV1alpha1().Battles(tc.namespace.Name).Create(ctx, battle, metav1.CreateOptions{})
	require.NoError(t, err)

	// Assert battle eventually fininshes
	gotInBattleSpiritsStats := requireBattleFinishes(t, ctx, battle)
	require.Len(t, gotInBattleSpiritsStats, 2)
	require.Equal(t, int64(0), gotInBattleSpiritsStats["the-battle-spirit-a-1"].Health)
	require.Equal(t, int64(1), gotInBattleSpiritsStats["the-battle-spirit-b-1"].Health)

	// Update one of the spirits
	spirits[0], err = tc.spiritsClientset.SpiritsV1alpha1().Spirits(tc.namespace.Name).Get(ctx, spirits[0].Name, metav1.GetOptions{})
	require.NoError(t, err)
	spirits[0].Spec.Stats.Power = 2
	spirits[0], err = tc.spiritsClientset.SpiritsV1alpha1().Spirits(tc.namespace.Name).Update(ctx, spirits[0], metav1.UpdateOptions{})
	require.NoError(t, err)

	// Assert battle eventually finishes with a different ending
	gotInBattleSpiritsStats = requireBattleFinishes(t, ctx, battle)
	t.Log(gotInBattleSpiritsStats)
	require.Len(t, gotInBattleSpiritsStats, 2)
	require.Equal(t, int64(3), gotInBattleSpiritsStats["the-battle-spirit-a-2"].Health)
	require.Equal(t, int64(0), gotInBattleSpiritsStats["the-battle-spirit-b-1"].Health)

	// Make sure there are only 2 spirits for that battle
	spiritsList, err := tc.spiritsClientset.SpiritsV1alpha1().Spirits(tc.namespace.Name).List(ctx, metav1.ListOptions{
		LabelSelector: "spirits.ankeesler.github.io/battle-name=" + battle.Name,
	})
	require.NoError(t, err)
	require.Len(t, spiritsList.Items, 2)
}

func TestInvalidBattles(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()

	// Read test fixtures
	var err error
	for _, path := range []string{
		"battle-bad-spirits.yaml",
	} {
		t.Run(path, func(t *testing.T) {
			battle := readObject(t, path).(*spiritsv1alpha1.Battle)
			battle, err = tc.spiritsClientset.SpiritsV1alpha1().Battles(tc.namespace.Name).Create(ctx, battle, metav1.CreateOptions{})
			require.NoError(t, err)

			// Assert battle is errored
			requireEventuallyConsistent(t, func() (bool, error) {
				battle, err := tc.spiritsClientset.SpiritsV1alpha1().Battles(tc.namespace.Name).Get(ctx, battle.Name, metav1.GetOptions{})
				if err != nil {
					return false, fmt.Errorf("get: %w", err)
				}
				t.Logf("got battle %q conditions: %#v", battle.Name, battle.Status.Conditions)
				return meta.IsStatusConditionFalse(battle.Status.Conditions, "Progressing"), nil
			})
		})
	}
}

func requireBattleFinishes(t *testing.T, ctx context.Context, battle *spiritsv1alpha1.Battle) map[string]*spiritsv1alpha1.SpiritStats {
	t.Helper()

	// Read spirits
	var spirits []*spiritsv1alpha1.Spirit
	for _, spiritRef := range battle.Spec.Spirits {
		spirit, err := tc.spiritsClientset.SpiritsV1alpha1().Spirits(tc.namespace.Name).Get(ctx, spiritRef.Name, metav1.GetOptions{})
		require.NoError(t, err)
		spirits = append(spirits, spirit)
	}

	// Assert spirits and battle are happy
	for i := range spirits {
		requireEventuallyConsistent(t, func() (bool, error) {
			spirit, err := tc.spiritsClientset.SpiritsV1alpha1().Spirits(tc.namespace.Name).Get(ctx, spirits[i].Name, metav1.GetOptions{})
			if err != nil {
				return false, fmt.Errorf("get: %w", err)
			}
			t.Logf("got spirit %q conditions: %#v", spirit.Name, spirit.Status.Conditions)
			return meta.IsStatusConditionTrue(spirit.Status.Conditions, "Ready"), nil
		})
	}
	requireEventuallyConsistent(t, func() (bool, error) {
		battle, err := tc.spiritsClientset.SpiritsV1alpha1().Battles(tc.namespace.Name).Get(ctx, battle.Name, metav1.GetOptions{})
		if err != nil {
			return false, fmt.Errorf("get: %w", err)
		}
		t.Logf("got battle %q conditions: %#v", battle.Name, battle.Status.Conditions)
		return meta.IsStatusConditionTrue(battle.Status.Conditions, "Progressing"), nil
	})

	// Assert battle finished
	requireEventuallyConsistent(t, func() (bool, error) {
		battle, err := tc.spiritsClientset.SpiritsV1alpha1().Battles(tc.namespace.Name).Get(ctx, battle.Name, metav1.GetOptions{})
		if err != nil {
			return false, fmt.Errorf("get: %w", err)
		}
		t.Logf("got battle %q phase: %q", battle.Name, battle.Status.Phase)
		return battle.Status.Phase == spiritsv1alpha1.BattlePhaseFinished, nil
	})

	// Assert spirit stats at end of battle
	battle, err := tc.spiritsClientset.SpiritsV1alpha1().Battles(tc.namespace.Name).Get(ctx, battle.Name, metav1.GetOptions{})
	require.NoError(t, err)
	gotInBattleSpiritsStats := getInBattleSpiritStats(t, ctx, tc.spiritsClientset, battle)
	return gotInBattleSpiritsStats
}

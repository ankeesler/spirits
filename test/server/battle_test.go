package server

import (
	"context"
	"fmt"
	"testing"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	inputv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/input/v1alpha1"
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
	t.Log(gotInBattleSpiritsStats)
	require.Len(t, gotInBattleSpiritsStats, 2)
	require.Equal(t, int64(0), gotInBattleSpiritsStats["the-battle-1-spirit-a-1"].Health)
	require.Equal(t, int64(1), gotInBattleSpiritsStats["the-battle-1-spirit-b-1"].Health)

	// Update one of the spirits
	spirits[0], err = tc.spiritsClientset.SpiritsV1alpha1().Spirits(tc.namespace.Name).Get(ctx, spirits[0].Name, metav1.GetOptions{})
	require.NoError(t, err)
	spirits[0].Spec.Attributes.Stats.Power = 2
	spirits[0], err = tc.spiritsClientset.SpiritsV1alpha1().Spirits(tc.namespace.Name).Update(ctx, spirits[0], metav1.UpdateOptions{})
	require.NoError(t, err)

	// Assert battle eventually finishes with a different ending
	gotInBattleSpiritsStats = requireBattleFinishes(t, ctx, battle)
	t.Log(gotInBattleSpiritsStats)
	require.Len(t, gotInBattleSpiritsStats, 2)
	require.Equal(t, int64(3), gotInBattleSpiritsStats["the-battle-1-spirit-a-2"].Health)
	require.Equal(t, int64(0), gotInBattleSpiritsStats["the-battle-1-spirit-b-1"].Health)

	// Make sure there are only 2 spirits for that battle
	spiritsList, err := tc.spiritsClientset.SpiritsV1alpha1().Spirits(tc.namespace.Name).List(ctx, metav1.ListOptions{
		LabelSelector: "spirits.ankeesler.github.io/battle-name=" + battle.Name,
	})
	require.NoError(t, err)
	require.Len(t, spiritsList.Items, 2)
}

func TestHumanIntelligenceBattle(t *testing.T) {
	ctx, cancel := context.WithTimeout(tc.ctx, time.Minute)
	defer cancel()

	// Read test fixtures
	var spirits []*spiritsv1alpha1.Spirit
	// TODO: maybe we should consider deleting spirits after every test, or using separate namespace and going in parallel...
	for _, spiritPath := range []string{"spirit-human.yaml", "spirit-c.yaml"} {
		spirits = append(spirits, readObject(t, spiritPath).(*spiritsv1alpha1.Spirit))
	}
	battle := readObject(t, "the-human-battle.yaml").(*spiritsv1alpha1.Battle)

	// Create test fixtures
	var err error
	for i := range spirits {
		spirits[i], err = tc.spiritsClientset.SpiritsV1alpha1().Spirits(tc.namespace.Name).Create(ctx, spirits[i], metav1.CreateOptions{})
		require.NoError(t, err)
	}
	battle, err = tc.spiritsClientset.SpiritsV1alpha1().Battles(tc.namespace.Name).Create(ctx, battle, metav1.CreateOptions{})
	require.NoError(t, err)

	// Assert battle is happy
	requireEventuallyConsistentBattle(t, ctx, battle.Name, func(battle *spiritsv1alpha1.Battle) bool {
		return meta.IsStatusConditionTrue(battle.Status.Conditions, "Progressing")
	})

	// Assert battle eventually requests an action
	requireEventuallyConsistentBattle(t, ctx, battle.Name, func(battle *spiritsv1alpha1.Battle) bool {
		return battle.Status.Phase == spiritsv1alpha1.BattlePhaseAwaitingAction
	})

	// Submit action
	actionCall := &inputv1alpha1.ActionCall{
		ObjectMeta: metav1.ObjectMeta{GenerateName: "action-request-"},
		Spec: inputv1alpha1.ActionCallSpec{
			Battle:     corev1.LocalObjectReference{Name: battle.Name},
			Spirit:     corev1.LocalObjectReference{Name: spirits[0].Name},
			ActionName: "attack",
		},
	}
	actionCall, err = tc.spiritsClientset.InputV1alpha1().ActionCalls(tc.namespace.Name).Create(ctx, actionCall, metav1.CreateOptions{})
	require.NoError(t, err)
	require.Equal(t, inputv1alpha1.ActionCallResultAccepted, actionCall.Status.Result, actionCall.Status.Message)

	// Assert battle eventually requests another action
	requireEventuallyConsistentBattle(t, ctx, battle.Name, func(battle *spiritsv1alpha1.Battle) bool {
		return battle.Status.Phase == spiritsv1alpha1.BattlePhaseAwaitingAction
	})

	// Submit the wrong action
	actionCall = &inputv1alpha1.ActionCall{
		ObjectMeta: metav1.ObjectMeta{GenerateName: "action-request-"},
		Spec: inputv1alpha1.ActionCallSpec{
			Battle:     corev1.LocalObjectReference{Name: battle.Name},
			Spirit:     corev1.LocalObjectReference{Name: spirits[1].Name},
			ActionName: "attack",
		},
	}
	actionCall, err = tc.spiritsClientset.InputV1alpha1().ActionCalls(tc.namespace.Name).Create(ctx, actionCall, metav1.CreateOptions{})
	require.NoError(t, err)
	require.Equal(t, inputv1alpha1.ActionCallResultRejected, actionCall.Status.Result, actionCall.Status.Message)

	// Submit another (correct) action
	actionCall = &inputv1alpha1.ActionCall{
		ObjectMeta: metav1.ObjectMeta{GenerateName: "action-request-"},
		Spec: inputv1alpha1.ActionCallSpec{
			Battle:     corev1.LocalObjectReference{Name: battle.Name},
			Spirit:     corev1.LocalObjectReference{Name: spirits[0].Name},
			ActionName: "charge",
		},
	}
	actionCall, err = tc.spiritsClientset.InputV1alpha1().ActionCalls(tc.namespace.Name).Create(ctx, actionCall, metav1.CreateOptions{})
	require.NoError(t, err)
	require.Equal(t, inputv1alpha1.ActionCallResultAccepted, actionCall.Status.Result, actionCall.Status.Message)

	// Assert battle eventually finishes
	gotInBattleSpiritsStats := requireBattleFinishes(t, ctx, battle)
	t.Log(gotInBattleSpiritsStats)
	require.Len(t, gotInBattleSpiritsStats, 2)
	require.Equal(t, int64(1), gotInBattleSpiritsStats["the-human-battle-1-spirit-human-1"].Health)
	require.Equal(t, int64(0), gotInBattleSpiritsStats["the-human-battle-1-spirit-c-1"].Health)
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

func TestUnsolicitedActionCall(t *testing.T) {
	ctx, cancel := context.WithTimeout(tc.ctx, time.Minute)
	defer cancel()

	actionCall := &inputv1alpha1.ActionCall{
		ObjectMeta: metav1.ObjectMeta{GenerateName: "action-request-"},
		Spec: inputv1alpha1.ActionCallSpec{
			Battle:     corev1.LocalObjectReference{Name: "some-battle-name"},
			Spirit:     corev1.LocalObjectReference{Name: "some-spirit-name"},
			ActionName: "attack",
		},
	}
	var err error
	actionCall, err = tc.spiritsClientset.InputV1alpha1().ActionCalls(tc.namespace.Name).Create(ctx, actionCall, metav1.CreateOptions{})
	require.NoError(t, err)
	require.Equal(t, inputv1alpha1.ActionCallResultRejected, actionCall.Status.Result, actionCall.Status.Message)
}

func requireBattleFinishes(t *testing.T, ctx context.Context, battle *spiritsv1alpha1.Battle) map[string]*spiritsv1alpha1.SpiritStats {
	// Read spirits
	var spirits []*spiritsv1alpha1.Spirit
	for _, spiritRef := range battle.Spec.Spirits {
		spirit, err := tc.spiritsClientset.SpiritsV1alpha1().Spirits(tc.namespace.Name).Get(ctx, spiritRef.Name, metav1.GetOptions{})
		require.NoError(t, err)
		spirits = append(spirits, spirit)
	}

	// Assert spirits and battle are happy
	for i := range spirits {
		requireEventuallyConsistentSpirit(t, ctx, spirits[i].Name, func(spirit *spiritsv1alpha1.Spirit) bool {
			return meta.IsStatusConditionTrue(spirit.Status.Conditions, "Ready")
		})
	}
	requireEventuallyConsistentBattle(t, ctx, battle.Name, func(battle *spiritsv1alpha1.Battle) bool {
		return meta.IsStatusConditionTrue(battle.Status.Conditions, "Progressing")
	})

	// Assert battle finished
	requireEventuallyConsistentBattle(t, ctx, battle.Name, func(battle *spiritsv1alpha1.Battle) bool {
		return battle.Status.Phase == spiritsv1alpha1.BattlePhaseFinished
	})

	// Assert spirit stats at end of battle
	battle, err := tc.spiritsClientset.SpiritsV1alpha1().Battles(tc.namespace.Name).Get(ctx, battle.Name, metav1.GetOptions{})
	require.NoError(t, err)
	gotInBattleSpiritsStats := getInBattleSpiritStats(t, ctx, tc.spiritsClientset, battle)
	return gotInBattleSpiritsStats
}

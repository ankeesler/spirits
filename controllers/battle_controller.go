/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"errors"
	"fmt"
	"sync"

	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	spiritsdevv1alpha1 "github.com/ankeesler/spirits/api/v1alpha1"
	spiritpkg "github.com/ankeesler/spirits/internal/spirit"
	"github.com/go-logr/logr"
)

// BattleReconciler reconciles a Battle object
type BattleReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	SpiritsCache *sync.Map
	BattlesCache *sync.Map
}

//+kubebuilder:rbac:groups=spirits.dev,resources=battles,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=spirits.dev,resources=battles/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=spirits.dev,resources=battles/finalizers,verbs=update

//+kubebuilder:rbac:groups=spirits.dev,resources=spirits,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *BattleReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// Get battle - if it doesn't exist, then: stop it, delete it from the cache, and return.
	var battle spiritsdevv1alpha1.Battle
	if err := r.Get(ctx, req.NamespacedName, &battle); err != nil {
		if k8serrors.IsNotFound(err) {
			internalBattle, ok := r.BattlesCache.LoadAndDelete(req.Name)
			if ok {
				internalBattle.Stop()
			}
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, fmt.Errorf("could not get battle: %w", err)
	}

	updateFunc := func() error {
		// Update conditions on current battle status
		battle.Status.Conditions = []metav1.Condition{}
		internalSpirits, condition := r.readyInternalSpirits(ctx, log, &battle)
		battle.Status.Conditions = append(battle.Status.Conditions, condition)
		battle.Status.Conditions = append(battle.Status.Conditions, r.progressBattle(ctx, log, &battle, internalSpirits))

		// Update battle phase
		battle.Status.Phase = getPhase(&battle)

		return nil
	}

	// Update battle
	if _, err := controllerutil.CreateOrPatch(ctx, r.Client, &battle, updateFunc); err != nil {
		log.Error(err, "could not patch spirit")
		return ctrl.Result{}, nil
	}

	log.Info("reconciled battle", "namespace", req.Namespace, "name", req.Name)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *BattleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&spiritsdevv1alpha1.Battle{}).
		Owns(&spiritsdevv1alpha1.Spirit{}).
		// TODO: also need to watch spirits that are used in a battle...
		Complete(r)
}

func (r *BattleReconciler) readyInternalSpirits(
	ctx context.Context,
	log logr.Logger,
	battle *spiritsdevv1alpha1.Battle,
) ([]*spiritpkg.Spirit, metav1.Condition) {
	const conditionType = "Ready"

	// Ensure internal spirits are in the cache.
	internalSpirits, err := r.getInternalSpirits(battle)
	if err != nil {
		return nil, newCondition(battle, conditionType, fmt.Errorf("could not find internal spirits: %w", err))
	}

	return internalSpirits, newCondition(battle, conditionType, nil)
}

func (r *BattleReconciler) getInternalSpirits(battle *spiritsdevv1alpha1.Battle) ([]*spiritpkg.Spirit, error) {
	var internalSpirits []*spiritpkg.Spirit
	for _, spiritName := range battle.Spec.Spirits {
		internalSpirit, ok := r.SpiritsCache.Load(spiritName)
		if !ok {
			return nil, fmt.Errorf("unknown spirit %q", spiritName)
		}
		internalSpirits = append(internalSpirits, internalSpirit)
	}
	return internalSpirits, nil
}

func (r *BattleReconciler) progressBattle(
	ctx context.Context,
	log logr.Logger,
	battle *spiritsdevv1alpha1.Battle,
	internalSpirits []*spiritpkg.Spirit,
) metav1.Condition {
	const conditionType = "Progressing"

	ready := meta.IsStatusConditionTrue(battle.Status.Conditions, "Ready")
	started := meta.IsStatusConditionTrue(battle.Status.Conditions, conditionType)

	if !ready && started {
		// A precondition is wrong, but the battle is running. Stop the battle.
		return newCondition(battle, conditionType, errors.New("battle is not ready"))
	}

	if ready && !started {
		// All preconditions are met, but the battle is not running. Start the battle.
		return newCondition(battle, conditionType, nil)
	}

	// If we get here, then we are already in the correct state. So no updated needed.
	return *meta.FindStatusCondition(battle.Status.Conditions, conditionType)
}

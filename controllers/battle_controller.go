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
	"fmt"
	"strings"
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
	battlepkg "github.com/ankeesler/spirits/internal/battle"
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
			// TODO: cancel battle
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, fmt.Errorf("could not get battle: %w", err)
	}

	// Update battle
	if _, err := controllerutil.CreateOrPatch(ctx, r.Client, &battle, func() error {
		// Update conditions on current battle status
		battle.Status.Conditions = []metav1.Condition{
			newCondition(&battle, "Ready", r.readySpirits(ctx, log, &battle)),
			newCondition(&battle, "Progressing", r.progressBattle(ctx, log, &battle)),
		}

		// Update battle phase
		battle.Status.Phase = getPhase(battle.Status.Conditions)

		return nil
	}); err != nil {
		log.Info("battle", "battle", battle)
		return ctrl.Result{}, fmt.Errorf("could not patch battle: %w", err)
	}

	log.Info("reconciled battle")

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

func (r *BattleReconciler) readySpirits(
	ctx context.Context,
	log logr.Logger,
	battle *spiritsdevv1alpha1.Battle,
) error {
	// TODO: what happens if someone changes the model spirit...
	// TODO: if we change the in-battle spirit, it remains the same...
	// TODO: if someone lists the same spirits in the battle, then there is gonna be an issue

	// Reinitialize the in-battle spirits list
	battle.Status.InBattleSpirits = []string{}

	// Get internal spirits from cache
	internalSpirits, ok := r.getInternalSpirits(battle.Spec.Spirits)
	if !ok {
		return fmt.Errorf("could not find internal spirits %s", battle.Spec.Spirits)
	}

	// Generate the names of the in-battle external spirits for this battle
	for _, internalSpirit := range internalSpirits {
		battle.Status.InBattleSpirits = append(battle.Status.InBattleSpirits, fmt.Sprintf("%s-%s", battle.Name, internalSpirit.Name))
	}

	// Get the internal in-battle spirits, if they exist
	internalInBattleSpirits, _ := r.getInternalSpirits(battle.Status.InBattleSpirits)

	// For each internal spirit in the battle, ensure we have an external in-battle spirit
	for i := range internalSpirits {
		// If the in-battle internal spirit already exists, copy from that
		// Otherwise, copy from the model spirit
		internalInBattleSpirit := internalInBattleSpirits[i]
		if internalInBattleSpirit == nil {
			internalInBattleSpirit = internalSpirits[i]
		}

		// Declare in-battle external spirit
		externalInBattleSpirit := &spiritsdevv1alpha1.Spirit{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: battle.Namespace,
				Name:      battle.Status.InBattleSpirits[i],
			},
		}

		updateFunc := func() error {
			// Convert the internal spirit to an external spirit
			externalInBattleSpiritFromInternalInBattleSpirit, err := fromInternalSpirit(internalInBattleSpirit)
			if err != nil {
				return fmt.Errorf("could not convert internal spirit to in-battle external spirit: %w", err)
			}
			externalInBattleSpirit.Spec = externalInBattleSpiritFromInternalInBattleSpirit.Spec

			// Update the external spirit object so we can track it against this battle
			externalInBattleSpirit.ObjectMeta.Labels = map[string]string{
				"spirits.dev/battle-name": battle.Name,
			}
			externalInBattleSpirit.ObjectMeta.OwnerReferences = []metav1.OwnerReference{
				*metav1.NewControllerRef(battle, spiritsdevv1alpha1.SchemeBuilder.GroupVersion.WithKind("Battle")),
			}

			return nil
		}

		// Create or patch the in-battle external spirit
		if _, err := controllerutil.CreateOrPatch(ctx, r.Client, externalInBattleSpirit, updateFunc); err != nil {
			return fmt.Errorf("cannot create or update external in-battle spirit %s: %w", externalInBattleSpirit.Name, err)
		}
	}

	return nil
}

func (r *BattleReconciler) progressBattle(
	ctx context.Context,
	log logr.Logger,
	battle *spiritsdevv1alpha1.Battle,
) error {
	// TODO: if we run a battle and then delete the in-battle spirits, the battle doesn't get run again
	// Almost like the in-battle internal spirits should be the controller of the in-battle external spirits...

	// Get internal in-battle spirits from cache
	internalInBattleSpirits, ok := r.getInternalSpirits(battle.Status.InBattleSpirits)
	if !ok {
		return fmt.Errorf("could not find internal in-battle spirits %s", battle.Status.InBattleSpirits)
	}

	if r.loadInternalBattleCancel(battle) == nil {
		// No battle exists - let's start it
		ctx, cancel := context.WithCancel(context.Background())
		r.storeInternalBattleCancel(battle, cancel)

		// TODO: after this finishes, we should probably set a status saying the battle is complete
		go battlepkg.Run(ctx, internalInBattleSpirits, func(internalSpirits []*spiritpkg.Spirit, err error) {
			log.V(1).Info("battle callback", "spirits", internalSpirits, "error", err)

			// Redeclare battle so that we don't hold onto old battle
			battle := spiritsdevv1alpha1.Battle{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: battle.Namespace,
					Name:      battle.Name,
				},
			}

			// If there is an error, update the battle condition
			if err != nil {
				r.setBattleError(ctx, log, &battle, err)
				return
			}

			// Otherwise, update the in-battle external spirits
			for _, internalSpirit := range internalSpirits {
				externalSpirit := spiritsdevv1alpha1.Spirit{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: battle.Namespace,
						Name:      internalSpirit.Name,
					},
				}
				if _, err := controllerutil.CreateOrPatch(ctx, r.Client, &externalSpirit, func() error {
					externalSpiritFromInternalSpirit, err := fromInternalSpirit(internalSpirit)
					if err != nil {
						return fmt.Errorf("cannot convert in-battle internal spirit to external spirit: %w", err)
					}
					externalSpirit.Spec = externalSpiritFromInternalSpirit.Spec
					return nil
				}); err != nil {
					r.setBattleError(ctx, log, &battle, err)
					return
				}

				log.V(1).Info("updated external spirit from callback", "spirit", &externalSpirit)
			}
		})
	}

	return nil
}

func (r *BattleReconciler) getInternalSpirits(spiritNames []string) ([]*spiritpkg.Spirit, bool) {
	var internalSpirits []*spiritpkg.Spirit
	found := 0
	for _, spiritName := range spiritNames {
		internalSpirit, ok := r.SpiritsCache.Load(spiritName)
		if ok {
			internalSpirits = append(internalSpirits, internalSpirit.(*spiritpkg.Spirit))
			found++
		} else {
			internalSpirits = append(internalSpirits, nil)
		}
	}
	return internalSpirits, found > 0 && found == len(spiritNames)
}

func (r *BattleReconciler) loadInternalBattleCancel(battle *spiritsdevv1alpha1.Battle) context.CancelFunc {
	cancel, ok := r.BattlesCache.Load(getBattleCacheKey(battle))
	if !ok {
		return nil
	}
	return cancel.(context.CancelFunc)
}

func (r *BattleReconciler) storeInternalBattleCancel(battle *spiritsdevv1alpha1.Battle, cancel context.CancelFunc) {
	r.BattlesCache.Store(getBattleCacheKey(battle), cancel)
}

func (r *BattleReconciler) setBattleError(
	ctx context.Context,
	log logr.Logger,
	battle *spiritsdevv1alpha1.Battle,
	err error,
) {
	log.Error(err, "battle error")

	// TODO: I think this gets overwritten immediately
	if _, err := controllerutil.CreateOrPatch(ctx, r.Client, battle, func() error {
		meta.SetStatusCondition(&battle.Status.Conditions, newCondition(battle, "Progressing", err))
		return nil
	}); err != nil {
		log.Error(err, "could not set battle error")
	}
}

func getBattleCacheKey(battle *spiritsdevv1alpha1.Battle) string {
	return fmt.Sprintf("%s-%s-%s", battle.Namespace, battle.Name, strings.Join(battle.Status.InBattleSpirits, "-"))
}

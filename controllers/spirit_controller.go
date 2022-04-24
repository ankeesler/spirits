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
	"math/rand"
	"sync"

	"github.com/go-logr/logr"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	spiritsdevv1alpha1 "github.com/ankeesler/spirits/api/v1alpha1"
	spiritpkg "github.com/ankeesler/spirits/internal/spirit"
	"github.com/ankeesler/spirits/internal/spirit/action"
)

// SpiritReconciler reconciles a Spirit object
type SpiritReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	SpiritsCache *sync.Map
	Rand         *rand.Rand
}

//+kubebuilder:rbac:groups=spirits.dev,resources=spirits,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=spirits.dev,resources=spirits/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=spirits.dev,resources=spirits/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *SpiritReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// Get spirit - if it doesn't exist, then delete it from the cache and return.
	var spirit spiritsdevv1alpha1.Spirit
	if err := r.Get(ctx, req.NamespacedName, &spirit); err != nil {
		if k8serrors.IsNotFound(err) {
			r.SpiritsCache.Delete(req.Name)
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, fmt.Errorf("could not get spirit: %w", err)
	}

	updateFunc := func() error {
		// Update conditions on current spirit status
		spirit.Status.Conditions = []metav1.Condition{}
		spirit.Status.Conditions = append(spirit.Status.Conditions, r.readyInternalSpirits(ctx, log, &spirit))

		// Update spirit phase
		spirit.Status.Phase = getPhase(&spirit)

		return nil
	}

	// Update spirit
	if _, err := controllerutil.CreateOrPatch(ctx, r.Client, &spirit, updateFunc); err != nil {
		log.Error(err, "could not patch spirit")
		return ctrl.Result{}, nil
	}

	log.Info("reconciled spirit", "namespace", req.Namespace, "name", req.Name)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SpiritReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&spiritsdevv1alpha1.Spirit{}).
		Complete(r)
}

func (r *SpiritReconciler) readyInternalSpirits(
	ctx context.Context,
	log logr.Logger,
	spirit *spiritsdevv1alpha1.Spirit,
) metav1.Condition {
	const conditionType = "Ready"

	// Convert to internal spirit
	internalSpirit, err := toInternalSpirit(
		spirit,
		r.Rand,
		func(ctx context.Context, s *spiritsdevv1alpha1.Spirit) (spiritpkg.Action, error) {
			// TODO: implement human intelligence
			return action.Attack(), nil
		},
	)
	if err != nil {
		return newCondition(spirit, conditionType, fmt.Errorf("could not convert to internal spirit: %w", err))
	}

	// Store in cache
	r.SpiritsCache.Store(spirit.Name, internalSpirit)

	return newCondition(spirit, conditionType, nil)
}

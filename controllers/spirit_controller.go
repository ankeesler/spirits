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
	"time"

	"github.com/go-logr/logr"
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
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Spirit object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *SpiritReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	var spirit spiritsdevv1alpha1.Spirit
	if err := r.Get(ctx, req.NamespacedName, &spirit); err != nil {
		if k8serrors.IsNotFound(err) {
			r.SpiritsCache.Delete(spirit.Name)
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, fmt.Errorf("could not get spirit: %w", err)
	}

	internalSpirit, err := toInternalSpirit(
		&spirit,
		r.Rand,
		func(ctx context.Context, s *spiritsdevv1alpha1.Spirit) (spiritpkg.Action, error) {
			// TODO: implement human intelligence
			return action.Attack(), nil
		},
	)
	r.updateStatus(ctx, log, &spirit, err)
	if err != nil {
		return ctrl.Result{}, nil
	}

	r.SpiritsCache.Store(internalSpirit.Name, internalSpirit)
	log.Info("successfully reconciled spirit", "namespace", req.Namespace, "name", req.Name)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SpiritReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&spiritsdevv1alpha1.Spirit{}).
		Complete(r)
}

func (r *SpiritReconciler) updateStatus(ctx context.Context, log logr.Logger, spirit *spiritsdevv1alpha1.Spirit, err error) {
	condition := metav1.Condition{
		Type:               "Ready",
		Status:             metav1.ConditionTrue,
		Reason:             "Success",
		Message:            "spirit is ready",
		ObservedGeneration: spirit.Generation,
		LastTransitionTime: metav1.NewTime(time.Now()),
	}
	if err != nil {
		condition.Status = metav1.ConditionFalse
		condition.Reason = "Error"
		condition.Message = err.Error()
	}

	if _, err := controllerutil.CreateOrPatch(ctx, r.Client, spirit, func() error {
		meta.SetStatusCondition(&spirit.Status.Conditions, condition)
		return nil
	}); err != nil {
		log.Error(err, "could not patch spirit")
	}
}

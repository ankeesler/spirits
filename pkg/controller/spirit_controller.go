package controller

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

	spiritsdevv1alpha1 "github.com/ankeesler/spirits/pkg/api/v1alpha1"
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
			// TODO: this doesn't work across namespaces
			r.SpiritsCache.Delete(req.Name)
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, fmt.Errorf("could not get spirit: %w", err)
	}

	// Update spirit
	if _, err := controllerutil.CreateOrPatch(ctx, r.Client, &spirit, func() error {
		// Update conditions on current spirit status
		spirit.Status.Conditions = []metav1.Condition{
			newCondition(&spirit, "Ready", r.readySpirit(ctx, log, &spirit)),
		}

		// Update spirit phase
		spirit.Status.Phase = getPhase(spirit.Status.Conditions)

		return nil
	}); err != nil {
		return ctrl.Result{}, fmt.Errorf("could not patch sprit: %w", err)
	}

	log.Info("reconciled spirit")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SpiritReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&spiritsdevv1alpha1.Spirit{}).
		Complete(r)
}

func (r *SpiritReconciler) readySpirit(
	ctx context.Context,
	log logr.Logger,
	spirit *spiritsdevv1alpha1.Spirit,
) error {
	// Convert to internal spirit
	// internalSpirit, err := nil, nil toInternalSpirit(
	// 	spirit,
	// 	r.Rand,
	// 	func(ctx context.Context, s *spiritsdevv1alpha1.Spirit) (spiritpkg.Action, error) {
	// 		// TODO: implement human intelligence
	// 		// battleName := battleNameFromContext(ctx)
	// 		// action, err := r.ActionsCache.Get(battleName, s.Name)
	// 		// if err != nil {
	// 		//   return nil, fmt.Errorf("could not get action from cache: %w", err)
	// 		// }
	// 		// return action, nil
	// 		return action.Attack(), nil
	// 	},
	// )
	// if err != nil {
	// 	return fmt.Errorf("could not convert to internal spirit: %w", err)
	// }

	// // Store in cache
	// r.SpiritsCache.Store(spirit.Name, internalSpirit)

	return nil
}

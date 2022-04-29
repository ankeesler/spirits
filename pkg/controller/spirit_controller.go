package controller

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
)

// SpiritReconciler reconciles a Spirit object
type SpiritReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=spirits.ankeesler.github.com,resources=spirits,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=spirits.ankeesler.github.com,resources=spirits/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=spirits.ankeesler.github.com,resources=spirits/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *SpiritReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// Get spirit - if it doesn't exist, we don't care.
	var spirit spiritsinternal.Spirit
	if err := r.Get(ctx, req.NamespacedName, &spirit); err != nil {
		if k8serrors.IsNotFound(err) {
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
		spirit.Status.Phase = getSpiritPhase(spirit.Status.Conditions)

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
		For(&spiritsv1alpha1.Spirit{}).
		Complete(r)
}

func (r *SpiritReconciler) readySpirit(
	ctx context.Context,
	log logr.Logger,
	spirit *spiritsinternal.Spirit,
) error {
	return nil
}

func getSpiritPhase(conditions []metav1.Condition) spiritsinternal.SpiritPhase {
	for i := range conditions {
		if conditions[i].Status == metav1.ConditionFalse {
			return spiritsinternal.SpiritPhaseError
		}
	}
	return spiritsinternal.SpiritPhaseReady
}

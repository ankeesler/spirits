package controller

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
)

// SpiritReconciler reconciles a Spirit object
type SpiritReconciler struct {
	client.Client
}

// SetupWithManager sets up the controller with the Manager.
func (r *SpiritReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&spiritsv1alpha1.Spirit{}).
		Named("spirit").
		Complete(r)
}

func (r *SpiritReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var spirit spiritsv1alpha1.Spirit
	return reconcile(ctx, req, &reconcileHelper[*spiritsv1alpha1.Spirit]{
		Client:   r.Client,
		Object:   &spirit,
		OnUpsert: r.onUpsert,
	})
}

func (r *SpiritReconciler) onUpsert(
	ctx context.Context,
	log logr.Logger,
	req ctrl.Request,
	spirit *spiritsv1alpha1.Spirit,
) error {
	// Update conditions on current spirit status
	spirit.Status.Conditions = []metav1.Condition{
		newCondition(spirit, readyCondition, r.readySpirit(ctx, log, spirit)),
	}

	// Update spirit phase
	spirit.Status.Phase = getSpiritPhase(spirit.Status.Conditions)

	return nil
}

func (r *SpiritReconciler) readySpirit(
	ctx context.Context,
	log logr.Logger,
	spirit *spiritsv1alpha1.Spirit,
) error {
	if _, err := getAction(&spirit.Spec.Action, func(ctx context.Context) (spiritsinternal.Action, error) {
		return nil, errors.New("this was a placeholder function meant for spirit validation")
	}); err != nil {
		return fmt.Errorf("get action: %w", err)
	}
	return nil
}

func getSpiritPhase(conditions []metav1.Condition) spiritsv1alpha1.SpiritPhase {
	for i := range conditions {
		if conditions[i].Status == metav1.ConditionFalse {
			return spiritsv1alpha1.SpiritPhaseError
		}
	}
	return spiritsv1alpha1.SpiritPhaseReady
}

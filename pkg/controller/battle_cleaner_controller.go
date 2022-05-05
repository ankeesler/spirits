package controller

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
)

// BattleCleanerReconciler reconciles a Spirit object
type BattleCleanerReconciler struct {
	client.Client
}

// SetupWithManager sets up the controller with the Manager.
func (r *BattleCleanerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&spiritsv1alpha1.Spirit{}).
		Complete(r)
}

func (r *BattleCleanerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var spirit spiritsv1alpha1.Spirit
	return reconcile(ctx, req, &reconcileHelper[*spiritsv1alpha1.Spirit]{
		Client:   r.Client,
		Object:   &spirit,
		OnUpsert: r.onUpsert,
	})
}

func (r *BattleCleanerReconciler) onUpsert(
	ctx context.Context,
	log logr.Logger,
	req ctrl.Request,
	spirit *spiritsv1alpha1.Spirit,
) error {
	// Delete all in-battle spirits with this spirit's name that don't match its generation
	inBattleSpiritsList, err := r.listInBattleSpirits(ctx, log, req)
	if err != nil {
		return fmt.Errorf("list in battle spirits: %w", err)
	}
	for i := range inBattleSpiritsList.Items {
		if inBattleSpiritsList.Items[i].Generation == spirit.Generation {
			continue
		}
		if err := r.Delete(ctx, &inBattleSpiritsList.Items[i]); err != nil {
			return fmt.Errorf("delete: %w", err)
		}
	}
	return nil
}

func (r *BattleCleanerReconciler) onDelete(
	ctx context.Context,
	log logr.Logger,
	req ctrl.Request,
	spirit *spiritsv1alpha1.Spirit,
) error {
	// Delete all in-battle spirits with this spirit's name (because it is getting deleted too)
	// TODO: could we do this with an owner reference on the spirit?
	inBattleSpiritsList, err := r.listInBattleSpirits(ctx, log, req)
	if err != nil {
		return fmt.Errorf("list in battle spirits: %w", err)
	}
	for i := range inBattleSpiritsList.Items {
		if err := r.Delete(ctx, &inBattleSpiritsList.Items[i]); err != nil {
			return fmt.Errorf("delete: %w", err)
		}
	}
	return nil
}

func (r *BattleCleanerReconciler) listInBattleSpirits(
	ctx context.Context,
	log logr.Logger,
	req ctrl.Request,
) (*spiritsv1alpha1.SpiritList, error) {
	var inBattleSpiritsList spiritsv1alpha1.SpiritList
	if err := r.List(
		ctx,
		&inBattleSpiritsList,
		client.InNamespace(req.Namespace),
		client.MatchingLabels(map[string]string{
			inBattleSpiritSpiritNameLabel: req.Name,
		}),
	); err != nil {
		return nil, fmt.Errorf("list: %w", err)
	}
	return &inBattleSpiritsList, nil
}

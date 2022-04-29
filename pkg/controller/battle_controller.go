package controller

import (
	"context"
	"fmt"
	"sync"

	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
	"github.com/go-logr/logr"
)

// BattleReconciler reconciles a Battle object
type BattleReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	BattlesCache *sync.Map
}

//+kubebuilder:rbac:groups=ankeesler.github.com,resources=battles,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=ankeesler.github.com,resources=battles/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=ankeesler.github.com,resources=battles/finalizers,verbs=update

//+kubebuilder:rbac:groups=ankeesler.github.com,resources=spirits,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *BattleReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// Get battle - if it doesn't exist, then: stop it, delete it from the cache, and return.
	var battle spiritsinternal.Battle
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
			newCondition(&battle, "Progress", r.progressBattle(ctx, log, &battle)),
		}

		return nil
	}); err != nil {
		log.Info("battle", "battle", battle)
		return ctrl.Result{}, fmt.Errorf("could not patch battle: %w", err)
	}

	log.Info("reconciled battle")

	return ctrl.Result{}, nil
}

func (r *BattleReconciler) progressBattle(
	ctx context.Context,
	log logr.Logger,
	battle *spiritsinternal.Battle,
) error {
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *BattleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&spiritsv1alpha1.Battle{}).
		Owns(&spiritsv1alpha1.Spirit{}).
		// TODO: also need to watch spirits that are used in a battle...
		Complete(r)
}

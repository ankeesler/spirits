package controller

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/go-logr/logr"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/ankeesler/spirits/internal/action"
	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
)

type Actions interface {
	Pend(battleName, spiritName, spiritGeneration string) (string, error)
}

// SpiritReconciler reconciles a Spirit object
type SpiritReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	Actions Actions
}

// SetupWithManager sets up the controller with the Manager.
func (r *SpiritReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&spiritsv1alpha1.Spirit{}).
		Complete(r)
}

//+kubebuilder:rbac:groups=spirits.ankeesler.github.com,resources=spirits,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=spirits.ankeesler.github.com,resources=spirits/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=spirits.ankeesler.github.com,resources=spirits/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *SpiritReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// Get spirit - if it doesn't exist, we don't care.
	var externalSpirit spiritsv1alpha1.Spirit
	if err := r.Get(ctx, req.NamespacedName, &externalSpirit); err != nil {
		if k8serrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, fmt.Errorf("could not get spirit: %w", err)
	}

	var spirit spiritsinternal.Spirit
	if err := r.Scheme.Convert(&externalSpirit, &spirit, nil); err != nil {
		return ctrl.Result{}, fmt.Errorf("convert: %w", err)
	}

	spiritPatch := client.MergeFrom(spirit.DeepCopyObject().(client.Object))

	// Update conditions on current spirit status
	spirit.Status.Conditions = []metav1.Condition{
		newCondition(&spirit, "Ready", r.readySpirit(ctx, log, &spirit)),
	}

	// Update spirit phase
	spirit.Status.Phase = getSpiritPhase(spirit.Status.Conditions)

	if err := r.Patch(ctx, &spirit, spiritPatch); err != nil {
		return ctrl.Result{}, fmt.Errorf("patch spirit: %w", err)
	}

	log.Info("reconciled spirit")

	return ctrl.Result{}, nil
}

func (r *SpiritReconciler) readySpirit(
	ctx context.Context,
	log logr.Logger,
	spirit *spiritsinternal.Spirit,
) error {
	var err error
	spirit.Spec.Internal.Action, err = getAction(
		spirit.Spec.Actions,
		spirit.Spec.Intelligence,
		r.getLazyActionFunc(spirit),
	)
	if err != nil {
		return fmt.Errorf("get action: %w", err)
	}
	return nil
}

func (r *SpiritReconciler) getLazyActionFunc(
	spirit *spiritsinternal.Spirit,
) func(ctx context.Context) (spiritsinternal.Action, error) {
	return func(ctx context.Context) (spiritsinternal.Action, error) {
		battleName, ok := spirit.Labels[inBattleSpiritBattleNameLabel]
		if !ok {
			return nil, errors.New("unknown battle name")
		}

		spiritName, ok := spirit.Labels[inBattleSpiritSpiritGenerationLabel]
		if !ok {
			return nil, errors.New("unknown spirit name")
		}

		spiritGeneration, ok := spirit.Labels[inBattleSpiritSpiritGenerationLabel]
		if !ok {
			return nil, errors.New("unknown spirit generation")
		}

		actionName, err := r.Actions.Pend(battleName, spiritName, spiritGeneration)
		if err != nil {
			return nil, fmt.Errorf("pend action: %w", err)
		}

		action, err := getAction([]string{actionName}, "", nil)
		if err != nil {
			return nil, fmt.Errorf("get action: %w", err)
		}

		return action, nil
	}
}

func getAction(
	actionNames []string,
	intelligence spiritsinternal.SpiritIntelligence,
	lazyActionFunc func(ctx context.Context) (spiritsinternal.Action, error),
) (spiritsinternal.Action, error) {
	// Note: the spirit actions should always be at least of length 1 thanks to defaulting
	var actions []spiritsinternal.Action
	for _, actionName := range actionNames {
		switch actionName {
		case "", "attack":
			actions = append(actions, action.Attack())
		case "bolster":
			actions = append(actions, action.Bolster())
		case "drain":
			actions = append(actions, action.Drain())
		case "charge":
			actions = append(actions, action.Charge())
		default:
			return nil, fmt.Errorf("unrecognized action: %q", actionName)
		}
	}

	var internalAction spiritsinternal.Action
	switch intelligence {
	case "", "roundrobin":
		internalAction = action.RoundRobin(actions)
	case "random":
		internalAction = action.Random(rand.New(rand.NewSource(0)), actions)
	case "human":
		if lazyActionFunc == nil {
			return nil, errors.New("human action is not supported")
		}
		internalAction = action.Lazy(lazyActionFunc)
	default:
		return nil, fmt.Errorf("unrecognized intelligence: %q", intelligence)
	}

	return internalAction, nil
}

func getSpiritPhase(conditions []metav1.Condition) spiritsinternal.SpiritPhase {
	for i := range conditions {
		if conditions[i].Status == metav1.ConditionFalse {
			return spiritsinternal.SpiritPhaseError
		}
	}
	return spiritsinternal.SpiritPhaseReady
}

package controller

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/ankeesler/spirits/internal/action"
	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
)

//+kubebuilder:rbac:groups=spirits.ankeesler.github.com,resources=spirits,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=spirits.ankeesler.github.com,resources=spirits/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=spirits.ankeesler.github.com,resources=spirits/finalizers,verbs=update

// SpiritReconciler reconciles a Spirit object
type SpiritReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	ActionSource ActionSource
}

// SetupWithManager sets up the controller with the Manager.
func (r *SpiritReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&spiritsv1alpha1.Spirit{}).
		Complete(&reconciler[*spiritsv1alpha1.Spirit, *spiritsinternal.Spirit]{
			Client: r.Client,
			Scheme: r.Scheme,
			Handler: &spiritHandler{
				ActionSource: r.ActionSource,
			},
		})
}

type spiritHandler struct {
	ActionSource ActionSource
}

func (h *spiritHandler) NewExternal() *spiritsv1alpha1.Spirit { return &spiritsv1alpha1.Spirit{} }
func (h *spiritHandler) NewInternal() *spiritsinternal.Spirit { return &spiritsinternal.Spirit{} }

func (h *spiritHandler) OnDelete(context.Context, logr.Logger, ctrl.Request) error {
	return nil
}

func (h *spiritHandler) OnUpsert(
	ctx context.Context,
	log logr.Logger,
	req ctrl.Request,
	spirit *spiritsinternal.Spirit,
) error {
	// Update conditions on current spirit status
	spirit.Status.Conditions = []metav1.Condition{
		newCondition(spirit, readyCondition, h.readySpirit(ctx, log, spirit)),
	}

	// Update spirit phase
	spirit.Status.Phase = getSpiritPhase(spirit.Status.Conditions)

	return nil
}

func (h *spiritHandler) readySpirit(
	ctx context.Context,
	log logr.Logger,
	spirit *spiritsinternal.Spirit,
) error {
	if _, err := getAction(
		spirit.Spec.Actions,
		spirit.Spec.Intelligence,
		getLazyActionFunc(spirit, h.ActionSource),
	); err != nil {
		return fmt.Errorf("get action: %w", err)
	}
	return nil
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

	// Note: the spirit intelligence should always default to a non-empty string
	var internalAction spiritsinternal.Action
	switch intelligence {
	case spiritsinternal.SpiritIntelligenceRoundRobin:
		internalAction = action.RoundRobin(actions)
	case spiritsinternal.SpiritIntelligenceRandom:
		internalAction = action.Random(rand.New(rand.NewSource(0)), actions)
	case spiritsinternal.SpiritIntelligenceHuman:
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

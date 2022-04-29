package controller

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
)

const (
	inBattleSpiritBattleNameLabel       = "spirits.ankeesler.github.com/battle-name"
	inBattleSpiritBattleGenerationLabel = "spirits.ankeesler.github.com/battle-generation"
	inBattleSpiritSpiritNameLabel       = "spirits.ankeesler.github.com/spirit-name"
	inBattleSpiritSpiritGenerationLabel = "spirits.ankeesler.github.com/spirit-generation"

	readyCondition       = "Ready"
	progressingCondition = "Progressing"
)

type ActionSource interface {
	Pend(ctx context.Context, battleName, battleGeneration, spiritName, spiritGeneration string) (string, error)
}

type handler[ExternalT, InternalT client.Object] interface {
	NewExternal() ExternalT
	NewInternal() InternalT
	OnUpsert(context.Context, logr.Logger, ctrl.Request, InternalT) error
	OnDelete(context.Context, logr.Logger, ctrl.Request) error
}

type reconciler[ExternalT, InternalT client.Object] struct {
	client.Client
	Scheme  *runtime.Scheme
	Handler handler[ExternalT, InternalT]
}

var _ reconcile.Reconciler = &reconciler[*spiritsv1alpha1.Spirit, *spiritsinternal.Spirit]{}

func (r *reconciler[ExternalT, InternalT]) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	externalT := r.Handler.NewExternal()
	if err := r.Get(ctx, req.NamespacedName, externalT); err != nil {
		if k8serrors.IsNotFound(err) {
			if err := r.Handler.OnDelete(ctx, log, req); err != nil {
				return ctrl.Result{}, fmt.Errorf("handle delete: %w", err)
			}
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, fmt.Errorf("get external object %w", err)
	}

	if _, err := controllerutil.CreateOrPatch(ctx, r.Client, externalT, func() error {
		internalT := r.Handler.NewInternal()
		if err := r.Scheme.Convert(externalT, internalT, nil); err != nil {
			return fmt.Errorf("convert from external object to internal object %w", err)
		}

		if err := r.Handler.OnUpsert(ctx, log, req, internalT); err != nil {
			return fmt.Errorf("handle upsert: %w", err)
		}

		if err := r.Scheme.Convert(internalT, externalT, nil); err != nil {
			return fmt.Errorf("convert from internal object to external object %w", err)
		}

		return nil
	}); err != nil {
		return ctrl.Result{}, fmt.Errorf("create or patch external object %w", err)
	}

	return ctrl.Result{}, nil
}

func newCondition(
	obj metav1.Object,
	teyep string,
	err error,
) metav1.Condition {
	condition := metav1.Condition{
		Type:               teyep,
		Status:             metav1.ConditionTrue,
		Reason:             "Success",
		Message:            "success",
		ObservedGeneration: obj.GetGeneration(),
		LastTransitionTime: metav1.NewTime(time.Now()),
	}
	if err != nil {
		condition.Status = metav1.ConditionFalse
		condition.Reason = "Error"
		condition.Message = err.Error()
	}
	return condition
}

func createOrPatch(
	ctx context.Context,
	client client.Client,
	scheme *runtime.Scheme,
	internalObj, externalObj client.Object,
	mutateFunc func() error,
) error {
	externalObj.SetNamespace(internalObj.GetNamespace())
	externalObj.SetName(internalObj.GetName())
	if _, err := controllerutil.CreateOrPatch(ctx, client, externalObj, func() error {
		if err := scheme.Convert(externalObj, internalObj, nil); err != nil {
			return fmt.Errorf("convert external object to internal object: %w", err)
		}

		if err := mutateFunc(); err != nil {
			return err
		}

		if err := scheme.Convert(internalObj, externalObj, nil); err != nil {
			return fmt.Errorf("convert internal object to external object: %w", err)
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}

func getLazyActionFunc(
	spirit *spiritsinternal.Spirit,
	actionsQueue ActionSource,
) func(ctx context.Context) (spiritsinternal.Action, error) {
	return func(ctx context.Context) (spiritsinternal.Action, error) {
		battleName, ok := spirit.Labels[inBattleSpiritBattleNameLabel]
		if !ok {
			return nil, errors.New("unknown battle name")
		}

		battleGeneration, ok := spirit.Labels[inBattleSpiritBattleGenerationLabel]
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

		actionName, err := actionsQueue.Pend(ctx, battleName, battleGeneration, spiritName, spiritGeneration)
		if err != nil {
			return nil, fmt.Errorf("actions queue pend: %w", err)
		}

		action, err := getAction([]string{actionName}, "", nil)
		if err != nil {
			return nil, fmt.Errorf("get action: %w", err)
		}

		return action, nil
	}
}

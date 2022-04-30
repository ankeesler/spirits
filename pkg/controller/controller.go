package controller

import (
	"context"
	"errors"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
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
	Pend(
		ctx context.Context,
		namespace, battleName, battleGeneration, spiritName, spiritGeneration string,
	) (string, error)
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
	actionSource ActionSource,
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

		actionName, err := actionSource.Pend(
			ctx,
			spirit.Namespace,
			battleName,
			battleGeneration,
			spiritName,
			spiritGeneration,
		)
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

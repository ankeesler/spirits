package controller

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	actioninternal "github.com/ankeesler/spirits/internal/action"
	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
)

type ActionSource interface {
	Pend(
		ctx context.Context,
		namespace, battleName, spiritName string,
	) (string, error)
}

func getAction(
	action *spiritsv1alpha1.SpiritAction,
	lazyActionFunc func(context.Context) (spiritsinternal.Action, error),
	scheme *runtime.Scheme,
) (spiritsinternal.Action, error) {
	// TODO: check to make sure validation of one field happens at admission

	// Well-known actions
	if action.WellKnown != nil {
		switch *action.WellKnown {
		case spiritsv1alpha1.SpiritWellKnownActionAttack:
			return actioninternal.Attack(), nil
		case spiritsv1alpha1.SpiritWellKnownActionNoop:
			return actioninternal.Noop(), nil
		default:
			return nil, fmt.Errorf("unrecognized action: %q", *action.WellKnown)
		}
	}

	// Action choices
	if action.Choices != nil {
		var internalActions []spiritsinternal.Action
		for _, namedAction := range action.Choices.Actions {
			internalAction, err := getAction(&namedAction.Action, lazyActionFunc, scheme)
			if err != nil {
				return nil, fmt.Errorf("invalid action choice %q: %w", namedAction.Name, err)
			}
			internalActions = append(internalActions, internalAction)
		}
		switch action.Choices.Intelligence {
		case spiritsv1alpha1.SpiritActionChoicesIntelligenceRoundRobin:
			return actioninternal.RoundRobin(internalActions), nil
		case spiritsv1alpha1.SpiritActionChoicesIntelligenceRandom:
			return actioninternal.Random(rand.New(rand.NewSource(0)), internalActions), nil
		case spiritsv1alpha1.SpiritActionChoicesIntelligenceHuman:
			if lazyActionFunc == nil {
				return nil, errors.New("human action is not supported")
			}
			return actioninternal.Lazy(lazyActionFunc), nil
		default:
			return nil, fmt.Errorf("unrecognized intelligence: %q", action.Choices.Intelligence)
		}
	}

	if action.Script != nil {
		action, err := actioninternal.Script(action.Script.APIVersion, action.Script.Text, scheme)
		if err != nil {
			return nil, fmt.Errorf("compile action script: %w", err)
		}
		return action, nil
	}

	// TODO: other action types

	return nil, fmt.Errorf("unsupported action")
}

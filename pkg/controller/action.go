package controller

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/ankeesler/spirits/internal/action"
	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
)

type ActionSource interface {
	Pend(
		ctx context.Context,
		namespace, battleName, battleGeneration, spiritName, spiritGeneration string,
	) (string, error)
}

func getAction(
	actionNames []string,
	intelligence spiritsv1alpha1.SpiritIntelligence,
	lazyActionFunc func(context.Context) (spiritsinternal.Action, error),
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
	case spiritsv1alpha1.SpiritIntelligenceRoundRobin:
		internalAction = action.RoundRobin(actions)
	case spiritsv1alpha1.SpiritIntelligenceRandom:
		internalAction = action.Random(rand.New(rand.NewSource(0)), actions)
	case spiritsv1alpha1.SpiritIntelligenceHuman:
		if lazyActionFunc == nil {
			return nil, errors.New("human action is not supported")
		}
		internalAction = action.Lazy(lazyActionFunc)
	default:
		return nil, fmt.Errorf("unrecognized intelligence: %q", intelligence)
	}

	return internalAction, nil
}

package controllers

import (
	"context"
	"fmt"
	"math/rand"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	spiritsdevv1alpha1 "github.com/ankeesler/spirits/api/v1alpha1"
	"github.com/ankeesler/spirits/internal/spirit"
	"github.com/ankeesler/spirits/internal/spirit/action"
)

type actions struct {
	ids []string
	spirit.Action
}

func toInternalSpirits(
	apiSpirits []*spiritsdevv1alpha1.Spirit,
	r *rand.Rand,
	humanActionFunc func(ctx context.Context, s *spiritsdevv1alpha1.Spirit) (spirit.Action, error),
) ([]*spirit.Spirit, error) {
	internalSpirits := make([]*spirit.Spirit, len(apiSpirits))
	var err error
	for i := range apiSpirits {
		internalSpirits[i], err = toInternalSpirit(apiSpirits[i], r, humanActionFunc)
		if err != nil {
			return nil, err
		}
	}
	return internalSpirits, nil
}

func toInternalSpirit(apiSpirit *spiritsdevv1alpha1.Spirit, r *rand.Rand, humanActionFunc func(ctx context.Context, s *spiritsdevv1alpha1.Spirit) (spirit.Action, error)) (*spirit.Spirit, error) {
	s := &spirit.Spirit{
		Name:    apiSpirit.Name,
		Health:  apiSpirit.Spec.Stats.Health,
		Power:   apiSpirit.Spec.Stats.Power,
		Agility: apiSpirit.Spec.Stats.Agility,
		Armor:   apiSpirit.Spec.Stats.Armor,
	}

	var lazyActionFunc func(ctx context.Context) (spirit.Action, error)
	if humanActionFunc != nil {
		lazyActionFunc = func(ctx context.Context) (spirit.Action, error) {
			return humanActionFunc(ctx, apiSpirit)
		}
	}
	var err error
	s.Action, err = toInternalAction(apiSpirit.Spec.Actions, apiSpirit.Spec.Intelligence, r, lazyActionFunc)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func toInternalAction(
	ids []string,
	intelligence spiritsdevv1alpha1.SpiritIntelligence,
	r *rand.Rand,
	lazyActionFunc func(ctx context.Context) (spirit.Action, error),
) (spirit.Action, error) {
	var internalActions []spirit.Action
	if len(ids) == 0 {
		internalActions = []spirit.Action{action.Attack()}
	} else {
		for _, id := range ids {
			switch id {
			case "", "attack":
				internalActions = append(internalActions, action.Attack())
			case "bolster":
				internalActions = append(internalActions, action.Bolster())
			case "drain":
				internalActions = append(internalActions, action.Drain())
			case "charge":
				internalActions = append(internalActions, action.Charge())
			default:
				return nil, fmt.Errorf("unrecognized action: %q", ids[0])
			}
		}
	}

	var internalAction spirit.Action
	switch intelligence {
	case spiritsdevv1alpha1.SpiritIntelligenceRoundRobin:
		internalAction = action.RoundRobin(internalActions)
	case spiritsdevv1alpha1.SpiritIntelligenceRandom:
		internalAction = action.Random(r, internalActions)
	case spiritsdevv1alpha1.SpiritIntelligenceHuman:
		if lazyActionFunc == nil {
			return nil, fmt.Errorf("unsupported intelligence value (hint: you must use websocket API): %q", intelligence)
		}
		internalAction = action.Lazy(lazyActionFunc)
	default:
		return nil, fmt.Errorf("unrecognized intelligence: %q", intelligence)
	}

	return &actions{ids: ids, Action: internalAction}, nil
}

func fromInternalSpirits(internalSpirits []*spirit.Spirit) ([]*spiritsdevv1alpha1.Spirit, error) {
	apiSpirits := make([]*spiritsdevv1alpha1.Spirit, len(internalSpirits))
	for i := range apiSpirits {
		var err error
		apiSpirits[i], err = fromInternalSpirit(internalSpirits[i])
		if err != nil {
			return nil, fmt.Errorf("could not convert from internal spirit %s: %w", internalSpirits[i].Name, err)
		}
	}
	return apiSpirits, nil
}

func fromInternalSpirit(internalSpirit *spirit.Spirit) (*spiritsdevv1alpha1.Spirit, error) {
	apiActions, ok := internalSpirit.Action.(*actions)
	if !ok {
		return nil, fmt.Errorf("invalid internal spirit action type: %T", internalSpirit.Action)
	}

	return &spiritsdevv1alpha1.Spirit{
		ObjectMeta: metav1.ObjectMeta{
			Name: internalSpirit.Name,
		},
		Spec: spiritsdevv1alpha1.SpiritSpec{
			Stats: spiritsdevv1alpha1.SpiritStats{
				Health:  internalSpirit.Health,
				Power:   internalSpirit.Power,
				Agility: internalSpirit.Agility,
				Armor:   internalSpirit.Armor,
			},
			Actions: apiActions.ids,
		},
	}, nil
}

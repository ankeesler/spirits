package battle

import (
	"context"
	"fmt"

	actionpkg "github.com/ankeesler/spirits/internal/action"
	spiritpkg "github.com/ankeesler/spirits/internal/spirit"
)

type SpiritIntelligence string

const (
	SpiritIntelligenceHuman  SpiritIntelligence = "human"
	SpiritIntelligenceRandom SpiritIntelligence = "random"
)

type internalActionSource interface {
	Pend(
		context.Context,
		*Battle,
		*Spirit,
		[]*Spirit,
		[][]*Spirit,
	) (string, []string, error)
}

type Spirit struct {
	*spiritpkg.Spirit

	intelligence SpiritIntelligence
	seed         int64

	battle *Battle

	actionSource internalActionSource
}

func (s *Spirit) Intelligence() SpiritIntelligence { return s.intelligence }

func (s *Spirit) Seed() int64 { return s.seed }

func (s *Spirit) Run(
	ctx context.Context,
	us []*Spirit,
	them [][]*Spirit,
) (context.Context, error) {
	me := s
	actionName, targetIDs, err := s.actionSource.Pend(ctx, s.battle, me, us, them)
	if err != nil {
		return ctx, fmt.Errorf("action source pend: %w", err)
	}

	internalAction := me.Action(actionName)
	if internalAction == nil {
		return nil, fmt.Errorf("spirit %s has no action with name %s", me.ID(), actionName)
	}

	var targets []actionpkg.Spirit
	for _, targetID := range targetIDs {
		target := findTarget(targetID, me, us, them)
		if target == nil {
			return nil, fmt.Errorf("no target spirit with id %s", targetID)
		}
		targets = append(targets, target)
	}

	return internalAction.Call(ctx, me, targets)
}

func findTarget(
	id string, me *Spirit, us []*Spirit, them [][]*Spirit) *Spirit {
	if me.ID() == id {
		return me
	}

	for _, usSpirit := range us {
		if usSpirit.ID() == id {
			return usSpirit
		}
	}

	for _, themSpirits := range them {
		for _, themSpirit := range themSpirits {
			if themSpirit.ID() == id {
				return themSpirit
			}
		}
	}

	return nil
}

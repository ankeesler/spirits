package spirit

import (
	"context"
	"fmt"

	actionpkg "github.com/ankeesler/spirits/internal/action"
	metapkg "github.com/ankeesler/spirits/internal/meta"
	"github.com/ankeesler/spirits/pkg/api"
)

type ActionRepo interface {
	Get(context.Context, string) (*actionpkg.Action, error)
}

type ActionSource interface {
	Pend(context.Context, *Spirit, []*Spirit, [][]*Spirit) (string, []string, error)
}

type runner interface {
	Run(context.Context, *Spirit, []*Spirit, [][]*Spirit) (context.Context, error)
}

type runnerFunc func(context.Context, *Spirit, []*Spirit, [][]*Spirit) (context.Context, error)

func (f runnerFunc) Run(
	ctx context.Context, me *Spirit, us []*Spirit, them [][]*Spirit) (context.Context, error) {
	return f(ctx, me, us, them)
}

type Spirit struct {
	apiSpirit *api.Spirit

	*metapkg.Meta

	name string
	*stats
	actions map[string]*actionpkg.Action

	runner runner
}

func FromAPI(
	ctx context.Context,
	apiSpirit *api.Spirit,
	actionRepo ActionRepo,
	actionSource ActionSource,
) (*Spirit, error) {
	internalSpirit := &Spirit{
		apiSpirit: apiSpirit,

		Meta: metapkg.FromAPI(apiSpirit.GetMeta()),

		name:    apiSpirit.GetName(),
		stats:   statsFromAPI(apiSpirit.GetStats()),
		actions: make(map[string]*actionpkg.Action),
	}

	for _, apiSpiritAction := range apiSpirit.GetActions() {
		actionName := apiSpiritAction.GetName()
		if _, ok := internalSpirit.actions[actionName]; ok {
			return nil, fmt.Errorf("duplicate action name: %s", actionName)
		}

		var internalAction *actionpkg.Action
		var err error

		switch definition := apiSpiritAction.GetDefinition().(type) {
		case *api.SpiritAction_ActionId:
			internalAction, err = actionRepo.Get(ctx, definition.ActionId)
			if err != nil {
				return nil, fmt.Errorf("invalid action ID for %s: %w", apiSpiritAction.GetName(), err)
			}
		case *api.SpiritAction_Inline:
			internalAction, err = actionpkg.FromAPI(definition.Inline)
			if err != nil {
				return nil, fmt.Errorf("invalid action inline for %s: %w",
					apiSpiritAction.GetName(), err)
			}
		}

		internalSpirit.actions[apiSpiritAction.GetName()] = internalAction
	}

	internalSpirit.runner = runnerFunc(func(
		ctx context.Context, me *Spirit, us []*Spirit, them [][]*Spirit) (context.Context, error) {
		actionName, targetIDs, err := actionSource.Pend(ctx, me, us, them)
		if err != nil {
			return ctx, fmt.Errorf("action source pend: %w", err)
		}

		internalAction, ok := internalSpirit.actions[actionName]
		if !ok {
			return nil, fmt.Errorf("spirit %s has no action with name %s", internalSpirit.ID(), actionName)
		}

		var targets []actionpkg.Spirit
		for _, targetID := range targetIDs {
			target := findTarget(targetID, me, us, them)
			if target == nil {
				return nil, fmt.Errorf("no target spirit with id %s", targetID)
			}
			targets = append(targets, target)
		}

		return internalAction.Call(ctx, internalSpirit, targets)
	})

	return internalSpirit, nil
}

func (s *Spirit) Name() string { return s.name }

func (s *Spirit) ActionNames() []string {
	var actionNames []string
	for actionName := range s.actions {
		actionNames = append(actionNames, actionName)
	}
	return actionNames
}

func (s *Spirit) Action(name string) *actionpkg.Action {
	return s.actions[name]
}

func (s *Spirit) Run(ctx context.Context, us []*Spirit, them [][]*Spirit) (context.Context, error) {
	return s.runner.Run(ctx, s, us, them)
}

func (s *Spirit) ToAPI() *api.Spirit {
	apiSpirit := s.apiSpirit
	apiSpirit.Meta = s.Meta.ToAPI()
	return apiSpirit
}

func (s *Spirit) Clone() *Spirit {
	return &Spirit{
		Meta: s.Meta.Clone(),

		name:    s.name,
		stats:   s.stats.Clone(),
		actions: s.actions,

		runner: s.runner,
	}
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

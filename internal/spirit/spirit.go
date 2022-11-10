package spirit

import (
	"context"

	actionpkg "github.com/ankeesler/spirits/internal/action"
	metapkg "github.com/ankeesler/spirits/internal/meta"
)

type ActionRepo interface {
	Get(context.Context, string) (*actionpkg.Action, error)
}

type Spirit struct {
	*metapkg.Meta

	name string
	*stats
	actions map[string]*actionpkg.Action
}

func New(meta *metapkg.Meta) *Spirit {
	return &Spirit{
		Meta:    meta,
		stats:   &stats{},
		actions: make(map[string]*actionpkg.Action),
	}
}

func (s *Spirit) Name() string        { return s.name }
func (s *Spirit) SetName(name string) { s.name = name }

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

func (s *Spirit) AddAction(name string, action *actionpkg.Action) {
	s.actions[name] = action
}

func (s *Spirit) Clone() *Spirit {
	return &Spirit{
		Meta: s.Meta.Clone(),

		name:    s.name,
		stats:   s.stats.Clone(),
		actions: s.actions,
	}
}

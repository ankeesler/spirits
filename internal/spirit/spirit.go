package spirit

import (
	"context"
)

type Action interface {
	Run(context.Context, *Spirit, []*Spirit, [][]*Spirit) (context.Context, error)
}

type Spirit struct {
	id     string
	name   string
	stats  *Stats
	action Action
}

func New(id string, name string, stats *Stats, action Action) *Spirit {
	return &Spirit{
		id:     id,
		name:   name,
		stats:  stats,
		action: action,
	}
}

func (s *Spirit) ID() string    { return s.id }
func (s *Spirit) Name() string  { return s.name }
func (s *Spirit) Stats() *Stats { return s.stats }
func (s *Spirit) Act(ctx context.Context, us []*Spirit, them [][]*Spirit) (context.Context, error) {
	return s.action.Run(ctx, s, us, them)
}

func (s *Spirit) Clone() *Spirit {
	return &Spirit{
		id:     s.id,
		name:   s.name,
		stats:  s.stats.Clone(),
		action: s.action,
	}
}

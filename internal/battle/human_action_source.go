package battle

import (
	"context"
)

type humanActionSource struct {
	actionSource ActionSource
}

func (s *humanActionSource) Pend(
	ctx context.Context,
	battle *Battle,
	me *Spirit,
	us []*Spirit,
	them [][]*Spirit,
) (string, []string, error) {
	return s.actionSource.Pend(ctx, battle.ID(), me.ID(), battle.Turns())
}

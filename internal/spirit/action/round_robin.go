package action

import (
	"context"

	"github.com/ankeesler/spirits/internal/spirit"
)

type roundRobin struct {
	actions []spirit.Action
	next    int
}

func RoundRobin(actions []spirit.Action) spirit.Action {
	return &roundRobin{actions: actions, next: 0}
}

func (rr *roundRobin) Run(ctx context.Context, from, to *spirit.Spirit) error {
	if rr.next >= len(rr.actions) {
		rr.next = 0
	}
	next := rr.actions[rr.next]
	rr.next++
	next.Run(ctx, from, to)
	return nil
}

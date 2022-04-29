package action

import (
	"context"

	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
)

type roundRobin struct {
	actions []spiritsinternal.Action
	next    int
}

func RoundRobin(actions []spiritsinternal.Action) spiritsinternal.Action {
	return &roundRobin{actions: actions, next: 0}
}

func (rr *roundRobin) Run(ctx context.Context, from, to *spiritsinternal.Spirit) error {
	if rr.next >= len(rr.actions) {
		rr.next = 0
	}
	next := rr.actions[rr.next]
	rr.next++
	next.Run(ctx, from, to)
	return nil
}

func (rr *roundRobin) DeepCopyAction() spiritsinternal.Action {
	rrCopy := &roundRobin{
		next: rr.next,
	}
	for _, action := range rr.actions {
		rrCopy.actions = append(rrCopy.actions, action.DeepCopyAction())
	}
	return rrCopy
}

package action

import (
	"context"
	"math/rand"

	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
)

type random struct {
	actions []spiritsinternal.Action
	r       *rand.Rand
}

func Random(r *rand.Rand, actions []spiritsinternal.Action) spiritsinternal.Action {
	return &random{r: r, actions: actions}
}

func (r *random) Run(ctx context.Context, from, to *spiritsinternal.Spirit) error {
	r.actions[r.r.Intn(len(r.actions))].Run(ctx, from, to)
	return nil
}

func (r *random) DeepCopyAction() spiritsinternal.Action {
	rCopy := &random{
		r: r.r,
	}
	for _, action := range r.actions {
		rCopy.actions = append(rCopy.actions, action.DeepCopyAction())
	}
	return rCopy
}

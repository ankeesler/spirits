package action

import (
	"context"
	"math/rand"

	"github.com/ankeesler/spirits/internal/spirit"
)

type random struct {
	r       *rand.Rand
	actions []spirit.Action
}

func Random(r *rand.Rand, actions []spirit.Action) spirit.Action {
	return &random{r: r, actions: actions}
}

func (r *random) Run(ctx context.Context, from, to *spirit.Spirit) error {
	r.actions[r.r.Intn(len(r.actions))].Run(ctx, from, to)
	return nil
}

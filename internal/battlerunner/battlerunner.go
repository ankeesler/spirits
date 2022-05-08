package battlerunner

import (
	"context"
	"errors"
	"fmt"

	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
)

const maxTurns = 100

type Callback func(*spiritsinternal.Battle, []*spiritsinternal.Spirit, error)

func Run(
	ctx context.Context,
	battle *spiritsinternal.Battle,
	spirits []*spiritsinternal.Spirit,
	callback Callback,
) {
	turns := 0

	callback(battle, spirits, nil)

	s := newStrategy(spirits)
	for s.hasNext() {
		turns++
		if turns >= maxTurns {
			callback(battle, spirits, errors.New("too many turns"))
			return
		}
		from, to := s.next()

		if err := from.Spec.Internal.Action.Run(ctx, from, to); err != nil {
			callback(battle, spirits, fmt.Errorf("action errored: %w", err))
			return
		}

		callback(battle, spirits, nil)
	}
}

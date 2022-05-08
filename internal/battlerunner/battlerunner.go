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
	for {
		// Make sure the context has not been cancelled
		select {
		case <-ctx.Done():
			callback(battle, spirits, fmt.Errorf("context canceled: %w", ctx.Err()))
		default:
		}

		// Make sure we aren't infinite-looping
		turns++
		if turns >= maxTurns {
			callback(battle, spirits, errors.New("too many turns"))
			break
		}
		from, to := s.next()

		// Make sure there is a next spirit to run
		if !s.hasNext() {
			break
		}

		// Run the next action
		if err := from.Spec.Internal.Action.Run(ctx, from, to); err != nil {
			callback(battle, spirits, fmt.Errorf("action errored: %w", err))
			break
		}

		// Call the callback to alert the caller of the new turn
		callback(battle, spirits, nil)
	}
}

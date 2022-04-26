package battle

import (
	"context"
	"errors"
	"fmt"

	"github.com/ankeesler/spirits/api/internal/spirit"
)

const maxTurns = 100

func Run(ctx context.Context, spirits []*spirit.Spirit, onSpirits func([]*spirit.Spirit, error)) {
	turns := 0

	onSpirits(spirits, nil)

	s := newStrategy(spirits)
	for s.hasNext() {
		turns++
		if turns >= maxTurns {
			onSpirits(spirits, errors.New("too many turns"))
			return
		}
		from, to := s.next()

		if err := from.Action.Run(ctx, from, to); err != nil {
			onSpirits(spirits, fmt.Errorf("action errored: %w", err))
			return
		}

		onSpirits(spirits, nil)
	}
}

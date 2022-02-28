package battle

import (
	"errors"

	"github.com/ankeesler/spirits/internal/spirit"
)

const maxTurns = 100

func Run(spirits []*spirit.Spirit, onSpirits func([]*spirit.Spirit, error)) {
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
		to.Health -= from.Power
		if to.Health < 0 {
			to.Health = 0
		}
		onSpirits(spirits, nil)
	}
}

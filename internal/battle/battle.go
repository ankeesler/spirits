package battle

import (
	"errors"

	"github.com/ankeesler/spirits/internal/spirit"
)

const maxTurns = 100

type strategy struct {
	spirits []*spirit.Spirit
	first   bool
}

func newStrategy(spirits []*spirit.Spirit) *strategy {
	return &strategy{
		spirits: spirits,
		first:   false,
	}
}

func (s *strategy) hasNext() bool {
	spiritsAlive := 0
	for _, spirit := range s.spirits {
		if spirit.Health > 0 {
			spiritsAlive++
		}
	}
	return spiritsAlive > 1
}

func (s *strategy) next() (*spirit.Spirit, *spirit.Spirit) {
	s.first = !s.first
	if s.first {
		return s.spirits[0], s.spirits[1]
	} else {
		return s.spirits[1], s.spirits[0]
	}
}

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

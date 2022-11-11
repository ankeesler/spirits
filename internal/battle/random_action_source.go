package battle

import (
	"context"
	"math/rand"
)

type randomActionSource struct {
	r *rand.Rand
}

func (s *randomActionSource) Pend(
	ctx context.Context,
	battle *Battle,
	me *Spirit,
	us []*Spirit,
	them [][]*Spirit,
) (string, []string, error) {
	// Select action.
	actionNames := me.ActionNames()
	actionNum := s.r.Intn(len(actionNames))
	actionName := me.ActionNames()[actionNum]

	// Select targets.
	var targets []string
	for _, otherTeam := range them {
		for _, otherSpirit := range otherTeam {
			if s.r.Intn(2) == 0 {
				targets = append(targets, otherSpirit.ID())
			}
		}
	}

	return actionName, targets, nil
}

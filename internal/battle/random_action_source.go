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
	return "", nil, nil
}

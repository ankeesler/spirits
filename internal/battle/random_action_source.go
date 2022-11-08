package battle

import (
	"context"
	"math/rand"

	spiritpkg "github.com/ankeesler/spirits/internal/spirit"
)

type randomActionSource struct {
	r *rand.Rand
}

func (s randomActionSource) Pend(
	ctx context.Context,
	me *spiritpkg.Spirit,
	us []*spiritpkg.Spirit,
	them [][]*spiritpkg.Spirit,
) (string, []string, error) {
	return "", nil, nil
}

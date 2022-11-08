package memory

import (
	"context"

	spiritpkg "github.com/ankeesler/spirits/internal/spirit"
)

type Queue struct {
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) Post(
	ctx context.Context,
	battleID string,
	spiritID string,
	actionName string,
	targetSpiritIDs []string,
) error {
	return nil
}

func (q *Queue) Pend(
	ctx context.Context,
	me *spiritpkg.Spirit,
	us []*spiritpkg.Spirit,
	them [][]*spiritpkg.Spirit,
) (string, []string, error) {
	return "", nil, nil
}

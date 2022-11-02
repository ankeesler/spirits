package memory

import "context"

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
	battleID string,
	spiritID string,
) (string, []string, error) {
	return "", nil, nil
}

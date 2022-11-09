package memory

import (
	"context"
	"fmt"
	"sync"
)

type Queue struct {
	m sync.Map
}

type actionCall struct {
	actionName      string
	targetSpiritIDs []string
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) Pend(
	ctx context.Context,
	battleID string,
	spiritID string,
	turn int64,
) (string, []string, error) {
	select {
	case actionCall, ok := <-q.c(battleID, spiritID, turn):
		if !ok {
			return "", nil, fmt.Errorf(
				"channel closed for battleID %q spiritID %q turn %d actions",
				battleID,
				spiritID,
				turn,
			)
		}
		return actionCall.actionName, actionCall.targetSpiritIDs, nil
	case <-ctx.Done():
		return "", nil, fmt.Errorf(
			"context canceled for battleID %q spiritID %q turn %d actions",
			battleID,
			spiritID,
			turn,
		)
	}
}

func (q *Queue) Post(
	ctx context.Context,
	battleID string,
	spiritID string,
	turn int64,
	actionName string,
	targetSpiritIDs []string,
) error {
	select {
	case q.c(battleID, spiritID, turn) <- &actionCall{actionName: actionName, targetSpiritIDs: targetSpiritIDs}:
	default:
		return fmt.Errorf(
			"no one listening for battleID %q spiritID %q turn %d actions",
			battleID,
			spiritID,
			turn,
		)
	}
	return nil
}

func (q *Queue) c(
	battleID string,
	spiritID string,
	turn int64,
) chan *actionCall {
	key := fmt.Sprintf("%s-%s-%d", battleID, spiritID, turn)
	c := make(chan *actionCall)
	cc, exists := q.m.LoadOrStore(key, c)
	if exists {
		close(c)
		c = cc.(chan *actionCall)
	}
	return c
}

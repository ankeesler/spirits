package memory

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
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
	msg := fmt.Sprintf("action queue: pend %s", q.key(battleID, spiritID, turn))
	log.Print(msg)

	select {
	case actionCall, ok := <-q.c(battleID, spiritID, turn):
		if !ok {
			log.Printf("%s: channel closed", msg)
			return "", nil, fmt.Errorf(
				"channel closed for battleID %q spiritID %q turn %d actions",
				battleID,
				spiritID,
				turn,
			)
		}
		log.Printf("%s: done", msg)
		q.done(battleID, spiritID, turn)
		return actionCall.actionName, actionCall.targetSpiritIDs, nil
	case <-ctx.Done():
		log.Printf("%s: context cancelled", msg)
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
	msg := fmt.Sprintf("action queue: post %s", q.key(battleID, spiritID, turn))
	log.Print(msg)

	timer := time.NewTimer(time.Second * 3)
	select {
	case <-timer.C:
	case q.c(battleID, spiritID, turn) <- &actionCall{actionName: actionName, targetSpiritIDs: targetSpiritIDs}:
		if !timer.Stop() {
			<-timer.C
		}
		log.Printf("%s: done", msg)
		return nil
	}

	log.Printf("%s: no pend", msg)

	return fmt.Errorf(
		"no one listening for battleID %q spiritID %q turn %d actions",
		battleID,
		spiritID,
		turn,
	)
}

func (q *Queue) c(
	battleID string,
	spiritID string,
	turn int64,
) chan *actionCall {
	c := make(chan *actionCall)
	cc, exists := q.m.LoadOrStore(q.key(battleID, spiritID, turn), c)
	if exists {
		close(c)
		c = cc.(chan *actionCall)
	}
	return c
}

func (q *Queue) done(
	battleID string,
	spiritID string,
	turn int64,
) {
	cc, exists := q.m.LoadAndDelete(q.key(battleID, spiritID, turn))
	if exists {
		close(cc.(chan *actionCall))
	}
}

func (q *Queue) key(
	battleID string,
	spiritID string,
	turn int64,
) string {
	return fmt.Sprintf("%s-%s-%d", battleID, spiritID, turn)
}

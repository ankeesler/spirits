package runner

import (
	"context"
	"errors"
	"fmt"

	battlepkg "github.com/ankeesler/spirits0/internal/battle"
	"github.com/ankeesler/spirits0/internal/spirit"
)

const maxTurns = 100

type Queue interface {
	HasNext() bool
	Next() (*spirit.Spirit, []*spirit.Spirit, [][]*spirit.Spirit)
}

type Callback func(int, error)

type Runner struct {
	queue    Queue
	callback Callback

	teams       map[string]*battlepkg.Team // Team name -> team
	spiritTeams map[string]*battlepkg.Team // Spirit ID -> team
}

func New(queue Queue, callback Callback) *Runner {
	return &Runner{
		queue:    queue,
		callback: callback,
	}
}

func (r *Runner) Run(ctx context.Context) {
	turn := 0
	for {
		var err error

		select {
		case <-ctx.Done():
			r.callback(turn, ctx.Err())
		default:
		}

		turn++
		if turn >= maxTurns {
			r.callback(turn, errors.New("too many turns"))
			break
		}

		if !r.queue.HasNext() {
			break
		}

		me, us, them := r.queue.Next()
		ctx, err = me.Act(ctx, us, them)
		if err != nil {
			r.callback(turn, fmt.Errorf("action errored: %w", err))
			break
		}

		r.callback(turn, nil)
	}
}

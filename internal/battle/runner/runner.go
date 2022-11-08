package runner

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	battlepkg "github.com/ankeesler/spirits/internal/battle"
)

const maxTurns = 100

type BattleRepo interface {
	Watch(context.Context, chan<- *battlepkg.Battle) error
	Update(context.Context, *battlepkg.Battle) (*battlepkg.Battle, error)
}

type battleContext struct {
	context.Context

	cancel context.CancelFunc
}

type Runner struct {
	battleRepo BattleRepo

	battles        chan *battlepkg.Battle
	battleContexts map[string]*battleContext
}

func Wire(
	battleRepo BattleRepo,
) (*Runner, error) {
	c := make(chan *battlepkg.Battle)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	if err := battleRepo.Watch(ctx, c); err != nil {
		return nil, fmt.Errorf("start watch: %w", err)
	}

	return &Runner{
		battleRepo: battleRepo,

		battles:        c,
		battleContexts: make(map[string]*battleContext),
	}, nil
}

func (c *Runner) Run(ctx context.Context) error {
	log.Printf("Starting battle runner")

	for {
		select {
		case <-ctx.Done():
			close(c.battles)
			return fmt.Errorf("context cancelled: %w", ctx.Err())
		case battle, ok := <-c.battles:
			if !ok {
				return errors.New("battle runner watch closed")
			}

			needsUpdate, err := c.runBattle(ctx, battle)
			if err != nil {
				log.Printf("error in battle %v: %s", battle, err.Error())
				continue
			}

			if needsUpdate {
				if _, err := c.battleRepo.Update(ctx, battle); err != nil {
					log.Printf("Failed to update battle to %v", battle)
				}
			}
		}
	}
}

func (c *Runner) runBattle(ctx context.Context, battle *battlepkg.Battle) (bool, error) {
	log.Printf("running battle %v", battle)

	if battle.State() != battlepkg.StateStarted {
		return false, nil
	}

	if battle.Turns() > maxTurns {
		battle.SetState(battlepkg.StateError)
		battle.SetError(errors.New("hit max turns"))
	} else if !battle.HasNext() {
		battle.SetState(battlepkg.StateFinished)
	} else {
		battleCtx, ok := c.battleContexts[battle.ID()]
		if !ok {
			var battleCtx battleContext
			battleCtx.Context, battleCtx.cancel = context.WithCancel(ctx)
			c.battleContexts[battle.ID()] = &battleCtx
		}

		me, us, them := battle.Next()
		var err error
		battleCtx.Context, err = me.Run(battleCtx, us, them)
		if err != nil {
			return false, fmt.Errorf("spirit %v run: %w", me, err)
		}
	}

	return true, nil
}

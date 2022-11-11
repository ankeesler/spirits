package runner

import (
	"context"
	"errors"
	"fmt"
	"log"

	battlepkg "github.com/ankeesler/spirits/internal/battle"
)

const maxTurns = 100

type BattleRepo interface {
	Watch(context.Context, *string) (<-chan *battlepkg.Battle, error)
	Update(context.Context, *battlepkg.Battle) (*battlepkg.Battle, error)
}

type battleContext struct {
	ctx context.Context

	cancel context.CancelFunc
}

type Runner struct {
	battleRepo BattleRepo

	battleContexts map[string]*battleContext
}

func New(battleRepo BattleRepo) *Runner {
	return &Runner{
		battleRepo: battleRepo,

		battleContexts: make(map[string]*battleContext),
	}
}

func (r *Runner) Run(ctx context.Context) error {
	log.Printf("starting battle runner")

	battles, err := r.battleRepo.Watch(ctx, nil)
	if err != nil {
		return fmt.Errorf("start watch: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			log.Print("stopping battle runner")
			return nil
		case battle, ok := <-battles:
			if !ok {
				return errors.New("battle runner watch closed")
			}

			needsUpdate, err := r.runBattle(ctx, battle)
			if err != nil {
				log.Printf("error in battle %v: %s", battle, err.Error())

				battle.SetState(battlepkg.StateError)
				battle.SetErr(errors.New("hit max turns"))

				needsUpdate = true
			}

			if needsUpdate {
				if _, err := r.battleRepo.Update(ctx, battle); err != nil {
					log.Printf("Failed to update battle to %v", battle)
				}
			}
		}
	}
}

func (c *Runner) runBattle(ctx context.Context, battle *battlepkg.Battle) (bool, error) {
	if battle.State() != battlepkg.StateStarted {
		return false, nil
	}

	if battle.Turns() > maxTurns {
		log.Printf("hit max turns for battle %+v", battle)
		return false, errors.New("hit max turns for battle")
	} else if !battle.HasNext() {
		log.Printf("finished battle %+v", battle)

		battle.SetState(battlepkg.StateFinished)
	} else {
		log.Printf("running battle %+v", battle)

		battleCtx, ok := c.battleContexts[battle.ID()]
		if !ok {
			battleCtx = &battleContext{}
			battleCtx.ctx, battleCtx.cancel = context.WithCancel(ctx)
			c.battleContexts[battle.ID()] = battleCtx
		}

		me, us, them := battle.Next()
		var err error
		battleCtx.ctx, err = me.Run(battleCtx.ctx, us, them)
		if err != nil {
			return false, fmt.Errorf("spirit %v run: %w", me, err)
		}
	}

	return true, nil
}

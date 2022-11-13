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
			r.processBattle(ctx, battle)
		}
	}
}

func (r *Runner) processBattle(
	ctx context.Context, battle *battlepkg.Battle) {
	battleCtx, exists := r.battleContexts[battle.ID()]

	switch battle.State() {
	case battlepkg.StateStarted:
		if !exists {
			battleCtx = &battleContext{}
			battleCtx.ctx, battleCtx.cancel = context.WithCancel(ctx)
			r.battleContexts[battle.ID()] = battleCtx

			go func() {
				if err := r.runBattle(battleCtx.ctx, battle); err != nil {
					battle.SetState(battlepkg.StateError)
					battle.SetErr(err)
				} else {
					battle.SetState(battlepkg.StateFinished)
				}

				r.updateBattle(ctx, battle)
			}()
		}

	case battlepkg.StateCancelled:
		if exists {
			battleCtx.cancel()
			delete(r.battleContexts, battle.ID())
		}
	}
}

func (r *Runner) runBattle(ctx context.Context, battle *battlepkg.Battle) error {
	for {
		battle.SetState(battlepkg.StateStarted)
		r.updateBattle(ctx, battle)

		select {
		case <-ctx.Done():
			log.Printf("battle context done: %+v", battle)
			return ctx.Err()
		default:
		}

		if battle.Turns() > maxTurns {
			log.Printf("hit max turns for battle %+v", battle)
			return errors.New("hit max turns for battle")
		}

		if !battle.HasNext() {
			log.Printf("finished battle %+v", battle)
			return nil
		}

		log.Printf("really running battle %+v", battle)

		me, us, them, needsWaiting := battle.Next()
		r.handleNeedsWaiting(ctx, battle, needsWaiting)

		var err error
		ctx, err = me.Run(ctx, us, them)
		if err != nil {
			return fmt.Errorf("spirit %v run: %w", me, err)
		}
	}
}

func (r *Runner) handleNeedsWaiting(
	ctx context.Context, battle *battlepkg.Battle, needsWaiting bool) {
	// This is a horrible hack :(
	if needsWaiting {
		battle.SetState(battlepkg.StateWaiting)
		r.updateBattle(ctx, battle)
	}
}

func (r *Runner) updateBattle(ctx context.Context, battle *battlepkg.Battle) {
	if _, err := r.battleRepo.Update(ctx, battle); err != nil {
		log.Printf("failed to update battle to %v", battle)
	}
}

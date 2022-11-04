package controller

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ankeesler/spirits0/internal/api"
	battlepkg "github.com/ankeesler/spirits0/internal/battle"
	runnerpkg "github.com/ankeesler/spirits0/internal/battle/runner"
)

type Repo interface {
	Watch(context.Context, chan<- *api.Battle) error
	Update(context.Context, *api.Battle, func(*api.Battle) error) (*api.Battle, error)
}

type Controller struct {
	repo Repo

	battles       chan *api.Battle
	runnerCancels map[string]context.CancelFunc
}

func Wire(repo Repo) (*Controller, error) {
	c := make(chan *api.Battle)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	if err := repo.Watch(ctx, c); err != nil {
		return nil, fmt.Errorf("start watch: %w", err)
	}

	return &Controller{
		repo: repo,

		battles:       c,
		runnerCancels: make(map[string]context.CancelFunc),
	}, nil
}

func (c *Controller) Run(ctx context.Context) error {
	for battle := range c.battles {
		if battle.GetState() == api.BattleState_BATTLE_STATE_STARTED {
			if _, ok := c.runnerCancels[battle.GetMeta().GetId()]; !ok {
				internalBattle := c.wireBattle(battle)
				runner := runnerpkg.New(internalBattle, c.battleCallbackFunc(ctx, battle, internalBattle))
				ctx, cancel := context.WithCancel(ctx)
				go runner.Run(ctx)
				c.runnerCancels[battle.GetMeta().GetId()] = cancel
			}
		}
		if battle.GetState() == api.BattleState_BATTLE_STATE_CANCELLED {
			if cancel, ok := c.runnerCancels[battle.GetMeta().GetId()]; ok {
				cancel()
				delete(c.runnerCancels, battle.GetMeta().GetId())
			}
		}
	}
	return nil
}

func (c *Controller) wireBattle(apiBattle *api.Battle) *battlepkg.Battle {
	return nil
}

func (c *Controller) battleCallbackFunc(
	ctx context.Context,
	apiBattle *api.Battle,
	internalBattle *battlepkg.Battle,
) func(int, error) {
	return func(turn int, err error) {
		if err != nil {
			errorMessage := err.Error()
			apiBattle.ErrorMessage = &errorMessage
			apiBattle.State = api.BattleState_BATTLE_STATE_ERROR
			if _, err := c.repo.Update(
				ctx,
				apiBattle,
				func(*api.Battle) error { return nil },
			); err != nil {
				log.Printf("could not update battle: %s", err.Error())
			}
			return
		}

		actingSpirit := internalBattle.Queue().Peek()
		_ = actingSpirit
	}
}

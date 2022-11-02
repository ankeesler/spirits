package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/ankeesler/spirits0/internal/api"
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
				runner := runnerpkg.New()
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

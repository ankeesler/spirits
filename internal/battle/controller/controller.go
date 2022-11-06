package controller

import (
	"context"
	"fmt"
	"log"
	"time"

	actionpkg "github.com/ankeesler/spirits0/internal/action"
	"github.com/ankeesler/spirits0/internal/api"
	"github.com/ankeesler/spirits0/internal/battle"
	battlepkg "github.com/ankeesler/spirits0/internal/battle"
	"github.com/ankeesler/spirits0/internal/battle/queue"
	runnerpkg "github.com/ankeesler/spirits0/internal/battle/runner"
	"github.com/ankeesler/spirits0/internal/spirit"
	spiritpkg "github.com/ankeesler/spirits0/internal/spirit"
)

type BattleRepo interface {
	Watch(context.Context, chan<- *api.Battle) error
	Update(context.Context, *api.Battle, func(*api.Battle) error) (*api.Battle, error)
}

type ActionRepo interface {
	Get(context.Context, string) (*api.Action, error)
}

type ActionSource interface {
	Pend(context.Context, string, string) (string, []string, error)
}

type Controller struct {
	battleRepo BattleRepo
	actionRepo ActionRepo

	actionSource ActionSource

	battles       chan *api.Battle
	runnerCancels map[string]context.CancelFunc
}

func Wire(
	battleRepo BattleRepo,
	actionRepo ActionRepo,
	actionSource ActionSource,
) (*Controller, error) {
	c := make(chan *api.Battle)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	if err := battleRepo.Watch(ctx, c); err != nil {
		return nil, fmt.Errorf("start watch: %w", err)
	}

	return &Controller{
		battleRepo: battleRepo,
		actionRepo: actionRepo,

		actionSource: actionSource,

		battles:       c,
		runnerCancels: make(map[string]context.CancelFunc),
	}, nil
}

func (c *Controller) Run(ctx context.Context) error {
	log.Printf("Running battle controller with seen battle: %v", c.runnerCancels)

	for battle := range c.battles {
		log.Printf("Battle controller got battle: %#v", battle)

		switch battle.GetState() {
		case api.BattleState_BATTLE_STATE_STARTED:
			c.onStartedBattle(ctx, battle)
		case api.BattleState_BATTLE_STATE_CANCELLED:
			c.onCancelledBattle(ctx, battle)
		}
	}

	return nil
}

func (c *Controller) onStartedBattle(ctx context.Context, battle *api.Battle) {
	log.Printf("Battle controller sees started battle: %#v", battle)

	if _, ok := c.runnerCancels[battle.GetMeta().GetId()]; ok {
		return
	}

	internalBattle, err := c.apiBattleToInternalBattle(ctx, battle)
	if err != nil {
		log.Printf("Cannot convert api battle %#v to internal battle: %s", battle, err.Error())
		return
	}

	log.Printf("Battle controller starting battle: %#v (with internal battle: %#v)",
		battle, internalBattle)

	runner := runnerpkg.New(internalBattle.Queue(), c.battleCallbackFunc(ctx, battle, internalBattle))
	ctx, cancel := context.WithCancel(ctx)
	c.runnerCancels[battle.GetMeta().GetId()] = cancel
	go runner.Run(ctx)
}

func (c *Controller) onCancelledBattle(ctx context.Context, battle *api.Battle) {
	log.Printf("Battle controller sees cancelled battle: %#v", battle)
	cancel, ok := c.runnerCancels[battle.GetMeta().GetId()]
	if !ok {
		return
	}

	log.Printf("Cancelling battle: %#v", battle)

	cancel()
	delete(c.runnerCancels, battle.GetMeta().GetId())
}

func (c *Controller) apiBattleToInternalBattle(
	ctx context.Context,
	apiBattle *api.Battle,
) (*battlepkg.Battle, error) {
	var internalTeams []*battlepkg.Team
	var allInternalSpirits [][]*spirit.Spirit

	for _, apiTeam := range apiBattle.GetTeams() {
		var internalTeamSpirits []*spirit.Spirit

		for _, apiTeamSpirit := range apiTeam.GetSpirits() {
			internalSpirit, err := c.apiSpiritToInternalSpirit(ctx, apiBattle, apiTeamSpirit)
			if err != nil {
				return nil, fmt.Errorf("convert api spirit %s: %w",
					apiTeamSpirit.GetSpirit().GetMeta().GetId(), err)
			}
			internalTeamSpirits = append(internalTeamSpirits, internalSpirit)
		}

		allInternalSpirits = append(allInternalSpirits, internalTeamSpirits)
		internalTeams = append(internalTeams, battlepkg.NewTeam(apiTeam.GetName(), internalTeamSpirits))
	}

	queue := queue.New(allInternalSpirits)

	return battlepkg.New(apiBattle.GetMeta().GetId(), internalTeams, queue), nil
}

func (c *Controller) apiSpiritToInternalSpirit(
	ctx context.Context,
	apiBattle *api.Battle,
	apiSpirit *api.BattleTeamSpirit,
) (*spiritpkg.Spirit, error) {
	internalStats := spiritpkg.NewStats(
		apiSpirit.GetSpirit().GetStats().GetHealth(),
		apiSpirit.GetSpirit().GetStats().GetPhysicalPower(),
		apiSpirit.GetSpirit().GetStats().GetPhysicalConstitution(),
		apiSpirit.GetSpirit().GetStats().GetMentalPower(),
		apiSpirit.GetSpirit().GetStats().GetMentalConstitution(),
		apiSpirit.GetSpirit().GetStats().GetAgility(),
	)

	var internalAction spiritpkg.Action
	switch apiSpirit.GetIntelligence() {
	case api.BattleTeamSpiritIntelligence_BATTLE_TEAM_SPIRIT_INTELLIGENCE_RANDOM:
		internalAction = nil // actionpkg.Random(apiSpirit.GetSpirit())
	case api.BattleTeamSpiritIntelligence_BATTLE_TEAM_SPIRIT_INTELLIGENCE_HUMAN:
		internalAction = nil // actionpkg.Lazy(
		// c.selectActionFunc(ctx, apiBattle, apiSpirit.GetSpirit()))
	}

	return spiritpkg.New(
		apiSpirit.GetSpirit().GetMeta().GetId(),
		apiSpirit.GetSpirit().GetName(), internalStats, internalAction), nil
}

func (c *Controller) apiSpiritActionToInternalAction(
	ctx context.Context,
	apiSpiritAction *api.SpiritAction,
) (actionpkg.Action, error) {
	var internalAction actionpkg.Action

	switch definition := apiSpiritAction.GetDefinition().(type) {
	case *api.SpiritAction_ActionId:
		apiAction, err := c.actionRepo.Get(ctx, definition.ActionId)
		if err != nil {
			return nil, fmt.Errorf("invalid action ID for %s: %w", apiSpiritAction.GetName(), err)
		}

		internalAction, err = actionpkg.FromAPI(apiAction)
		if err != nil {
			return nil, fmt.Errorf("invalid action inline for %s: %w",
				apiSpiritAction.GetName(), err)
		}
	case *api.SpiritAction_Inline:
		var err error
		internalAction, err = actionpkg.FromAPI(definition.Inline)
		if err != nil {
			return nil, fmt.Errorf("invalid action inline for %s: %w",
				apiSpiritAction.GetName(), err)
		}
	}

	return internalAction, nil
}

func (c *Controller) internalBattleToAPIBattle(internalBattle *battle.Battle) *api.Battle {
	return nil
}

func (c *Controller) internalSpiritToAPISpirit(
	internalSpirit *spirit.Spirit,
) (*api.Spirit, api.BattleTeamSpiritIntelligence) {
	return nil, 0
}

func (c *Controller) selectActionFunc(
	ctx context.Context,
	apiBattle *api.Battle,
	apiSpirit *api.Spirit,
) func(context.Context) (actionpkg.Action, error) {
	return func(ctx context.Context) (actionpkg.Action, error) {
		actionName, _, err := c.actionSource.Pend(
			ctx, apiBattle.GetMeta().GetId(), apiSpirit.GetMeta().GetId())
		if err != nil {
			return nil, fmt.Errorf("action source pend: %w", err)
		}

		for _, apiSpiritAction := range apiSpirit.GetActions() {
			if apiSpiritAction.GetName() == actionName {
				return c.apiSpiritActionToInternalAction(ctx, apiSpiritAction)
			}
		}

		return nil, fmt.Errorf("api spirit %s in api battle %s does not have action with name %s",
			apiSpirit.GetMeta().GetId(), apiBattle.GetMeta().GetId(), actionName)
	}
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
		} else {
			for _, internalTeam := range internalBattle.Teams() {
				apiTeam := &api.BattleTeam{Name: internalTeam.Name()}
				for _, internalSpirit := range internalTeam.Spirits() {
					apiSpirit, intelligence := c.internalSpiritToAPISpirit(internalSpirit)
					apiTeam.Spirits = append(apiTeam.Spirits, &api.BattleTeamSpirit{
						Spirit:       apiSpirit,
						Intelligence: intelligence,
					})
				}
				apiBattle.Teams = append(apiBattle.Teams, apiTeam)
			}

			apiBattle.ActingSpiritId = internalBattle.Queue().Peek().ID()
		}

		if _, err := c.battleRepo.Update(
			ctx,
			apiBattle,
			func(*api.Battle) error { return nil },
		); err != nil {
			log.Printf("could not update battle: %s", err.Error())
		}
	}
}

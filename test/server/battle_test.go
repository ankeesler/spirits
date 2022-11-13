package test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/ankeesler/spirits/pkg/api"
)

func TestAutoBattle(t *testing.T) {
	state := startServer(t)
	clients := state.clients

	// Create battle.
	createBattleRsp, err := clients.battle.CreateBattle(state.ctx, &api.CreateBattleRequest{})
	if err != nil {
		t.Fatal("create battle:", err)
	}
	battleID := createBattleRsp.GetBattle().GetMeta().GetId()

	listSpiritsRsp, err := clients.spirit.ListSpirits(context.Background(), &api.ListSpiritsRequest{
		Name: stringPtr("zombie"),
	})
	if err != nil {
		t.Fatal("list spirits:", err)
	}
	if len(listSpiritsRsp.GetSpirits()) != 1 {
		t.Fatalf("wanted 1 spirit, got %s", listSpiritsRsp.GetSpirits())
	}
	zombieSpirit := listSpiritsRsp.GetSpirits()[0]

	teams := []struct {
		name      string
		spiritIDs []string
	}{
		{
			name:      "a",
			spiritIDs: []string{zombieSpirit.GetMeta().GetId()},
		},
		{
			name:      "b",
			spiritIDs: []string{zombieSpirit.GetMeta().GetId()},
		},
	}
	for _, team := range teams {
		if _, err := clients.battle.AddBattleTeam(state.ctx, &api.AddBattleTeamRequest{
			BattleId: battleID,
			TeamName: team.name,
		}); err != nil {
			t.Fatalf("add battle team %s: %v", team.name, err)
		}

		for _, spiritID := range team.spiritIDs {
			if _, err := clients.battle.AddBattleTeamSpirit(state.ctx, &api.AddBattleTeamSpiritRequest{
				BattleId:     battleID,
				TeamName:     team.name,
				SpiritId:     spiritID,
				Intelligence: api.BattleTeamSpiritIntelligence_BATTLE_TEAM_SPIRIT_INTELLIGENCE_RANDOM,
				Seed:         time.Now().Unix(),
			}); err != nil {
				t.Fatalf("add battle team %s: %v", team.name, err)
			}
		}
	}

	watchStream, err := clients.battle.WatchBattle(state.ctx, &api.WatchBattleRequest{
		Id: battleID,
	})
	if err != nil {
		t.Fatal("watch battle:", err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		if watchBattle(state.ctx, t, watchStream) != nil {
			t.Error("got WAITING battle state")
		}
		wg.Done()
	}()

	if _, err := clients.battle.StartBattle(state.ctx, &api.StartBattleRequest{
		Id: battleID,
	}); err != nil {
		t.Fatal("start battle:", err)
	}

	wg.Wait()
}

func TestManualBattle(t *testing.T) {
	state := startServer(t)
	clients := state.clients

	// Create battle.
	createBattleRsp, err := clients.battle.CreateBattle(state.ctx, &api.CreateBattleRequest{})
	if err != nil {
		t.Fatal("create battle:", err)
	}
	battleID := createBattleRsp.GetBattle().GetMeta().GetId()

	listSpiritsRsp, err := clients.spirit.ListSpirits(context.Background(), &api.ListSpiritsRequest{
		Name: stringPtr("zombie"),
	})
	if err != nil {
		t.Fatal("list spirits:", err)
	}
	if len(listSpiritsRsp.GetSpirits()) != 1 {
		t.Fatalf("wanted 1 spirit, got %s", listSpiritsRsp.GetSpirits())
	}
	zombieSpirit := listSpiritsRsp.GetSpirits()[0]

	teams := []struct {
		name      string
		spiritIDs []string
	}{
		{
			name:      "a",
			spiritIDs: []string{zombieSpirit.GetMeta().GetId()},
		},
		{
			name:      "b",
			spiritIDs: []string{zombieSpirit.GetMeta().GetId()},
		},
	}
	for _, team := range teams {
		if _, err := clients.battle.AddBattleTeam(state.ctx, &api.AddBattleTeamRequest{
			BattleId: battleID,
			TeamName: team.name,
		}); err != nil {
			t.Fatalf("add battle team %s: %v", team.name, err)
		}

		for _, spiritID := range team.spiritIDs {
			if _, err := clients.battle.AddBattleTeamSpirit(state.ctx, &api.AddBattleTeamSpiritRequest{
				BattleId:     battleID,
				TeamName:     team.name,
				SpiritId:     spiritID,
				Intelligence: api.BattleTeamSpiritIntelligence_BATTLE_TEAM_SPIRIT_INTELLIGENCE_HUMAN,
				Seed:         time.Now().Unix(),
			}); err != nil {
				t.Fatalf("add battle team %s: %v", team.name, err)
			}
		}
	}

	watchStream, err := clients.battle.WatchBattle(state.ctx, &api.WatchBattleRequest{
		Id: battleID,
	})
	if err != nil {
		t.Fatal("watch battle:", err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			battle := watchBattle(state.ctx, t, watchStream)
			if battle == nil {
				break
			}

			if _, err := clients.battle.CallAction(state.ctx, &api.CallActionRequest{
				BattleId:        battle.GetMeta().GetId(),
				SpiritId:        battle.GetNextSpiritIds()[0],
				Turn:            battle.GetTurns(),
				ActionName:      "attack",
				TargetSpiritIds: []string{zombieSpirit.Meta.GetId()},
			}); err != nil {
				t.Error("call action:", err)
			}
		}
		wg.Done()
	}()

	if _, err := clients.battle.StartBattle(state.ctx, &api.StartBattleRequest{
		Id: battleID,
	}); err != nil {
		t.Fatal("start battle:", err)
	}

	wg.Wait()
}

func watchBattle(ctx context.Context, t *testing.T, stream api.BattleService_WatchBattleClient) *api.Battle {
	t.Helper()

	c := make(chan *api.Battle)
	var cErr error
	go func() {
		for {
			rsp, err := stream.Recv()
			if err != nil {
				cErr = err
				close(c)
				return
			}
			c <- rsp.GetBattle()
		}
	}()

	for {
		select {
		case <-ctx.Done():
			t.Error("watch battle closed (client):", ctx.Err())
			return nil
		case battle, ok := <-c:
			if !ok {
				t.Error("rx goroutine failed with error:", cErr)
			}
			switch battle.GetState() {
			case api.BattleState_BATTLE_STATE_FINISHED, api.BattleState_BATTLE_STATE_CANCELLED, api.BattleState_BATTLE_STATE_ERROR:
				return nil
			case api.BattleState_BATTLE_STATE_WAITING:
				return battle
			}
		}
	}
}

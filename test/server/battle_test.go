package test

import (
	"context"
	"sync"
	"testing"
	"time"

	spiritsv1 "github.com/ankeesler/spirits/pkg/api/spirits/v1"
)

func TestAutoBattle(t *testing.T) {
	state := startServer(t)
	clients := state.clients

	t.Log("creating battle...")
	createBattleRsp, err := clients.battle.CreateBattle(state.ctx, &spiritsv1.CreateBattleRequest{})
	if err != nil {
		t.Fatal("create battle:", err)
	}
	battleID := createBattleRsp.GetBattle().GetMeta().GetId()

	t.Log("finding spirits...")
	listSpiritsRsp, err := clients.spirit.ListSpirits(context.Background(), &spiritsv1.ListSpiritsRequest{
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
		t.Logf("adding battle team %s...", team.name)
		if _, err := clients.battle.AddBattleTeam(state.ctx, &spiritsv1.AddBattleTeamRequest{
			BattleId: battleID,
			TeamName: team.name,
		}); err != nil {
			t.Fatalf("add battle team %s: %v", team.name, err)
		}

		for _, spiritID := range team.spiritIDs {
			t.Logf("adding battle team spirit to %s...", team.name)
			if _, err := clients.battle.AddBattleTeamSpirit(state.ctx, &spiritsv1.AddBattleTeamSpiritRequest{
				BattleId:     battleID,
				TeamName:     team.name,
				SpiritId:     spiritID,
				Intelligence: spiritsv1.BattleTeamSpiritIntelligence_BATTLE_TEAM_SPIRIT_INTELLIGENCE_RANDOM,
				Seed:         time.Now().Unix(),
			}); err != nil {
				t.Fatalf("add battle team %s: %v", team.name, err)
			}
		}
	}

	t.Log("watching battle...")
	watchStream, err := clients.battle.WatchBattle(state.ctx, &spiritsv1.WatchBattleRequest{
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

	t.Log("starting battle...")
	if _, err := clients.battle.StartBattle(state.ctx, &spiritsv1.StartBattleRequest{
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
	createBattleRsp, err := clients.battle.CreateBattle(state.ctx, &spiritsv1.CreateBattleRequest{})
	if err != nil {
		t.Fatal("create battle:", err)
	}
	battleID := createBattleRsp.GetBattle().GetMeta().GetId()

	listSpiritsRsp, err := clients.spirit.ListSpirits(context.Background(), &spiritsv1.ListSpiritsRequest{
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
		if _, err := clients.battle.AddBattleTeam(state.ctx, &spiritsv1.AddBattleTeamRequest{
			BattleId: battleID,
			TeamName: team.name,
		}); err != nil {
			t.Fatalf("add battle team %s: %v", team.name, err)
		}

		for _, spiritID := range team.spiritIDs {
			if _, err := clients.battle.AddBattleTeamSpirit(state.ctx, &spiritsv1.AddBattleTeamSpiritRequest{
				BattleId:     battleID,
				TeamName:     team.name,
				SpiritId:     spiritID,
				Intelligence: spiritsv1.BattleTeamSpiritIntelligence_BATTLE_TEAM_SPIRIT_INTELLIGENCE_HUMAN,
				Seed:         time.Now().Unix(),
			}); err != nil {
				t.Fatalf("add battle team %s: %v", team.name, err)
			}
		}
	}

	watchStream, err := clients.battle.WatchBattle(state.ctx, &spiritsv1.WatchBattleRequest{
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

			if _, err := clients.battle.CallAction(state.ctx, &spiritsv1.CallActionRequest{
				BattleId:        battle.GetMeta().GetId(),
				SpiritId:        battle.GetNextSpiritIds()[0],
				Turn:            battle.GetTurns(),
				ActionName:      "attack",
				TargetSpiritIds: []string{zombieSpirit.Meta.GetId()},
			}); err != nil {
				t.Error("call action:", err)
				break
			}
		}
		wg.Done()
	}()

	if _, err := clients.battle.StartBattle(state.ctx, &spiritsv1.StartBattleRequest{
		Id: battleID,
	}); err != nil {
		t.Fatal("start battle:", err)
	}

	wg.Wait()
}

func watchBattle(ctx context.Context, t *testing.T, stream spiritsv1.BattleService_WatchBattleClient) *spiritsv1.Battle {
	t.Helper()

	c := make(chan *spiritsv1.Battle)
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
		t.Log("watching...")
		select {
		case <-ctx.Done():
			t.Error("watch battle closed (client):", ctx.Err())
			return nil
		case battle, ok := <-c:
			if !ok {
				t.Error("rx goroutine failed with error:", cErr)
				return nil
			}
			switch battle.GetState() {
			case spiritsv1.BattleState_BATTLE_STATE_FINISHED, spiritsv1.BattleState_BATTLE_STATE_CANCELLED, spiritsv1.BattleState_BATTLE_STATE_ERROR:
				return nil
			case spiritsv1.BattleState_BATTLE_STATE_WAITING:
				return battle
			}
		}
	}
}

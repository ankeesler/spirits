package test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/ankeesler/spirits/pkg/api"
	"google.golang.org/protobuf/encoding/prototext"
)

func TestAutoBattle(t *testing.T) {
	clients := startServer(t)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// Create battle.
	createBattleRsp, err := clients.battle.CreateBattle(ctx, &api.CreateBattleRequest{})
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
		if _, err := clients.battle.AddBattleTeam(ctx, &api.AddBattleTeamRequest{
			BattleId: battleID,
			TeamName: team.name,
		}); err != nil {
			t.Fatalf("add battle team %s: %v", team.name, err)
		}

		for _, spiritID := range team.spiritIDs {
			if _, err := clients.battle.AddBattleTeamSpirit(ctx, &api.AddBattleTeamSpiritRequest{
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

	watchCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	watchStream, err := clients.battle.WatchBattle(watchCtx, &api.WatchBattleRequest{
		Id: battleID,
	})
	if err != nil {
		t.Fatal("watch battle:", err)
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		watchBattle(watchCtx, t, watchStream)
		wg.Done()
	}()

	if _, err := clients.battle.StartBattle(ctx, &api.StartBattleRequest{
		Id: battleID,
	}); err != nil {
		t.Fatal("start battle:", err)
	}

	wg.Wait()
}

func watchBattle(ctx context.Context, t *testing.T, stream api.BattleService_WatchBattleClient) {
	for {
		select {
		case <-ctx.Done():
			t.Errorf("watch battle closed (client): %s\n", ctx.Err().Error())
			return
		default:
		}

		rsp, err := stream.Recv()
		if err != nil {
			t.Errorf("watch battle closed (server): %s", err.Error())
			return
		}

		t.Log("watch battle: ", prototext.MarshalOptions{
			Multiline: true,
		}.Format(rsp))

		switch rsp.GetBattle().GetState() {
		case api.BattleState_BATTLE_STATE_FINISHED, api.BattleState_BATTLE_STATE_CANCELLED, api.BattleState_BATTLE_STATE_ERROR:
			return
		}
	}
}

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/ankeesler/spirits/internal/menu"
	spiritsv1 "github.com/ankeesler/spirits/pkg/api/spirits/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/prototext"
)

type clients struct {
	spirit spiritsv1.SpiritServiceClient
	battle spiritsv1.BattleServiceClient
}

type state struct {
	clients *clients

	battleID string
	teamName string
}

var stateKey struct{}

func setState(ctx context.Context, state *state) context.Context {
	return context.WithValue(ctx, stateKey, state)
}

func getState(ctx context.Context) *state {
	return ctx.Value(stateKey).(*state)
}

var (
	port = flag.Int("port", 12345, "The server port")
)

var battleMenu = menu.Menu{
	{
		Title: "Add Battle Team",
		Runner: menu.RunnerFunc(func(ctx context.Context, io *menu.IO) (context.Context, error) {
			state := getState(ctx)

			teamName, err := menu.Input(io, "Team name: ")
			if err != nil {
				return ctx, err
			}

			rsp, err := state.clients.battle.AddBattleTeam(ctx, &spiritsv1.AddBattleTeamRequest{
				BattleId: &state.battleID,
				TeamName: &teamName,
			})
			if err != nil {
				return ctx, fmt.Errorf("add battle team: %w", err)
			}

			state.teamName = teamName

			fmt.Fprint(io.Out, prototext.MarshalOptions{
				Multiline: true,
			}.Format(rsp))

			return menu.Menu{
				{
					Title: "Add Battle Team Spirit",
					Runner: menu.RunnerFunc(func(ctx context.Context, io *menu.IO) (context.Context, error) {
						state := getState(ctx)

						rsp, err := state.clients.spirit.ListSpirits(ctx, &spiritsv1.ListSpiritsRequest{})
						if err != nil {
							return ctx, fmt.Errorf("list spirits: %w", err)
						}

						var submenu menu.Menu
						for _, spirit := range rsp.GetSpirits() {
							submenu = append(submenu, menu.Item{
								Title: spirit.GetName(),
								Runner: menu.RunnerFunc(func(ctx context.Context, io *menu.IO) (context.Context, error) {
									state := getState(ctx)
									rsp, err := state.clients.battle.AddBattleTeamSpirit(ctx, &spiritsv1.AddBattleTeamSpiritRequest{
										BattleId:     &state.battleID,
										TeamName:     &state.teamName,
										SpiritId:     stringPtr(spirit.GetMeta().GetId()),
										Intelligence: intelligencePtr(spiritsv1.BattleTeamSpiritIntelligence_BATTLE_TEAM_SPIRIT_INTELLIGENCE_RANDOM),
										Seed:         int64Ptr(time.Now().Unix()),
									})
									if err != nil {
										return ctx, fmt.Errorf("add battle team spirit: %w", err)
									}

									fmt.Fprint(io.Out, prototext.MarshalOptions{
										Multiline: true,
									}.Format(rsp))

									return ctx, nil
								}),
							})
						}

						return submenu.Run(ctx, io)
					}),
				},
			}.Run(ctx, io)
		}),
	},
	{
		Title: "Start Battle",
		Runner: menu.RunnerFunc(func(ctx context.Context, io *menu.IO) (context.Context, error) {
			state := getState(ctx)

			watchCtx, cancel := context.WithCancel(ctx)
			defer cancel()
			watchStream, err := state.clients.battle.WatchBattle(watchCtx, &spiritsv1.WatchBattleRequest{
				Id: &state.battleID,
			})
			if err != nil {
				return ctx, fmt.Errorf("watch battle: %w", err)
			}

			wg := sync.WaitGroup{}

			wg.Add(1)
			go func() {
				watchBattle(watchCtx, io, watchStream)
				wg.Done()
			}()

			if _, err := state.clients.battle.StartBattle(ctx, &spiritsv1.StartBattleRequest{
				Id: &state.battleID,
			}); err != nil {
				return ctx, fmt.Errorf("start battle: %w", err)
			}

			wg.Wait()

			return ctx, nil
		}),
	},
}

var m = menu.Menu{
	{
		Title: "List Spirits",
		Runner: menu.RunnerFunc(func(ctx context.Context, io *menu.IO) (context.Context, error) {
			state := getState(ctx)

			rsp, err := state.clients.spirit.ListSpirits(ctx, &spiritsv1.ListSpiritsRequest{})
			if err != nil {
				return ctx, fmt.Errorf("list spirits: %w", err)
			}

			fmt.Fprint(io.Out, prototext.MarshalOptions{
				Multiline: true,
			}.Format(rsp))

			return ctx, nil
		}),
	},
	{
		Title: "List Battles",
		Runner: menu.RunnerFunc(func(ctx context.Context, io *menu.IO) (context.Context, error) {
			state := getState(ctx)

			rsp, err := state.clients.battle.ListBattles(ctx, &spiritsv1.ListBattlesRequest{})
			if err != nil {
				return ctx, fmt.Errorf("list battles: %w", err)
			}

			fmt.Fprint(io.Out, prototext.MarshalOptions{
				Multiline: true,
			}.Format(rsp))

			return ctx, nil
		}),
	},
	{
		Title: "Create Battle",
		Runner: menu.RunnerFunc(func(ctx context.Context, io *menu.IO) (context.Context, error) {
			state := getState(ctx)

			rsp, err := state.clients.battle.CreateBattle(ctx, &spiritsv1.CreateBattleRequest{})
			if err != nil {
				return ctx, fmt.Errorf("create battle: %w", err)
			}

			state.battleID = rsp.GetBattle().GetMeta().GetId()

			fmt.Fprint(io.Out, prototext.MarshalOptions{
				Multiline: true,
			}.Format(rsp))

			return battleMenu.Run(ctx, io)
		}),
	},
	{
		Title: "Update Battle",
		Runner: menu.RunnerFunc(func(ctx context.Context, io *menu.IO) (context.Context, error) {
			state := getState(ctx)

			var err error
			state.battleID, err = menu.Input(io, "Battle ID: ")
			if err != nil {
				return ctx, fmt.Errorf("input: %w", err)
			}

			return battleMenu.Run(ctx, io)
		}),
	},
	{
		Title: "Start Demo Battle",
		Runner: menu.RunnerFunc(func(ctx context.Context, io *menu.IO) (context.Context, error) {
			state := getState(ctx)

			createBattleRsp, err := state.clients.battle.CreateBattle(ctx, &spiritsv1.CreateBattleRequest{})
			if err != nil {
				return ctx, fmt.Errorf("create battle: %w", err)
			}
			battleID := createBattleRsp.GetBattle().GetMeta().GetId()

			listSpiritsRsp, err := state.clients.spirit.ListSpirits(context.Background(), &spiritsv1.ListSpiritsRequest{
				Name: stringPtr("zombie"),
			})
			if err != nil {
				return ctx, fmt.Errorf("list spirits: %w", err)
			}
			if len(listSpiritsRsp.GetSpirits()) != 1 {
				return ctx, fmt.Errorf("wanted 1 spirit, got %s", listSpiritsRsp.GetSpirits())
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
				if _, err := state.clients.battle.AddBattleTeam(ctx, &spiritsv1.AddBattleTeamRequest{
					BattleId: &battleID,
					TeamName: &team.name,
				}); err != nil {
					return ctx, fmt.Errorf("add battle team %s: %w", team.name, err)
				}

				for _, spiritID := range team.spiritIDs {
					if _, err := state.clients.battle.AddBattleTeamSpirit(ctx, &spiritsv1.AddBattleTeamSpiritRequest{
						BattleId:     &battleID,
						TeamName:     &team.name,
						SpiritId:     &spiritID,
						Intelligence: intelligencePtr(spiritsv1.BattleTeamSpiritIntelligence_BATTLE_TEAM_SPIRIT_INTELLIGENCE_RANDOM),
						Seed:         int64Ptr(time.Now().Unix()),
					}); err != nil {
						return ctx, fmt.Errorf("add battle team %s: %w", team.name, err)
					}
				}
			}

			watchCtx, cancel := context.WithCancel(ctx)
			defer cancel()
			watchStream, err := state.clients.battle.WatchBattle(watchCtx, &spiritsv1.WatchBattleRequest{
				Id: &battleID,
			})
			if err != nil {
				return ctx, fmt.Errorf("watch battle: %w", err)
			}

			wg := sync.WaitGroup{}

			wg.Add(1)
			go func() {
				watchBattle(watchCtx, io, watchStream)
				wg.Done()
			}()

			if _, err := state.clients.battle.StartBattle(ctx, &spiritsv1.StartBattleRequest{
				Id: &battleID,
			}); err != nil {
				return ctx, fmt.Errorf("start battle: %w", err)
			}

			wg.Wait()

			return ctx, nil
		}),
	},
	{
		Title: "Start Demo Human Battle",
		Runner: menu.RunnerFunc(func(ctx context.Context, io *menu.IO) (context.Context, error) {
			state := getState(ctx)

			createBattleRsp, err := state.clients.battle.CreateBattle(ctx, &spiritsv1.CreateBattleRequest{})
			if err != nil {
				return ctx, fmt.Errorf("create battle: %w", err)
			}
			battleID := createBattleRsp.GetBattle().GetMeta().GetId()
			fmt.Fprintln(io.Out, "Created battle")
			time.Sleep(time.Second * 4)

			listSpiritsRsp, err := state.clients.spirit.ListSpirits(context.Background(), &spiritsv1.ListSpiritsRequest{
				Name: stringPtr("zombie"),
			})
			if err != nil {
				return ctx, fmt.Errorf("list spirits: %w", err)
			}
			if len(listSpiritsRsp.GetSpirits()) != 1 {
				return ctx, fmt.Errorf("wanted 1 spirit, got %s", listSpiritsRsp.GetSpirits())
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
				if _, err := state.clients.battle.AddBattleTeam(ctx, &spiritsv1.AddBattleTeamRequest{
					BattleId: &battleID,
					TeamName: &team.name,
				}); err != nil {
					return ctx, fmt.Errorf("add battle team %s: %w", team.name, err)
				}
				fmt.Fprintf(io.Out, "Added %s battle team\n", team.name)
				time.Sleep(time.Second * 4)

				for _, spiritID := range team.spiritIDs {
					if _, err := state.clients.battle.AddBattleTeamSpirit(ctx, &spiritsv1.AddBattleTeamSpiritRequest{
						BattleId:     &battleID,
						TeamName:     &team.name,
						SpiritId:     &spiritID,
						Intelligence: intelligencePtr(spiritsv1.BattleTeamSpiritIntelligence_BATTLE_TEAM_SPIRIT_INTELLIGENCE_HUMAN),
						Seed:         int64Ptr(time.Now().Unix()),
					}); err != nil {
						return ctx, fmt.Errorf("add battle team %s: %w", team.name, err)
					}

					fmt.Fprintf(io.Out, "Added %s battle team spirit\n", team.name)
					time.Sleep(time.Second * 4)
				}
			}

			watchCtx, cancel := context.WithCancel(ctx)
			defer cancel()
			watchStream, err := state.clients.battle.WatchBattle(watchCtx, &spiritsv1.WatchBattleRequest{
				Id: &battleID,
			})
			if err != nil {
				return ctx, fmt.Errorf("watch battle: %w", err)
			}
			fmt.Fprintf(io.Out, "Watched battle\n")
			time.Sleep(time.Second * 4)

			wg := sync.WaitGroup{}

			wg.Add(1)
			go func() {
				for {
					battle := watchBattle(watchCtx, io, watchStream)
					if battle == nil {
						break
					}

					if _, err := menu.Input(io, "Press any button to call next action"); err != nil {
						fmt.Fprintln(io.Out, "input:", err)
						break
					}

					if _, err := state.clients.battle.CallAction(watchCtx, &spiritsv1.CallActionRequest{
						BattleId:        stringPtr(battle.GetMeta().GetId()),
						SpiritId:        stringPtr(battle.GetNextSpiritIds()[0]),
						Turn:            int64Ptr(battle.GetTurns()),
						ActionName:      stringPtr("attack"),
						TargetSpiritIds: []string{zombieSpirit.Meta.GetId()},
					}); err != nil {
						fmt.Fprintln(io.Out, "call action:", err)
						break
					}

					fmt.Fprintf(io.Out, "Called action")
					time.Sleep(time.Second * 4)
				}
				wg.Done()
			}()

			if _, err := state.clients.battle.StartBattle(ctx, &spiritsv1.StartBattleRequest{
				Id: &battleID,
			}); err != nil {
				return ctx, fmt.Errorf("start battle: %w", err)
			}
			fmt.Fprintf(io.Out, "Started battle\n")
			time.Sleep(time.Second * 4)

			wg.Wait()

			return ctx, nil
		}),
	},
}

func watchBattle(ctx context.Context, io *menu.IO, stream spiritsv1.BattleService_WatchBattleClient) *spiritsv1.Battle {
	c := make(chan *spiritsv1.Battle)
	go func() {
		for {
			rsp, err := stream.Recv()
			if err != nil {
				fmt.Fprintln(io.Out, "watch battle error:", err)
				close(c)
				return
			}
			c <- rsp.GetBattle()
		}
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Fprintln(io.Out, "watch battle closed (client):", ctx.Err())
			return nil
		case battle, ok := <-c:
			if !ok {
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

func main() {
	flag.Parse()

	conn, err := grpc.Dial(
		fmt.Sprintf(":%d", *port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := setState(context.Background(), &state{
		clients: &clients{
			spirit: spiritsv1.NewSpiritServiceClient(conn),
			battle: spiritsv1.NewBattleServiceClient(conn),
		},
	})
	if _, err := m.Run(ctx, &menu.IO{In: os.Stdin, Out: os.Stdout}); err != nil {
		log.Fatal(err)
	}
}

func stringPtr(s string) *string { return &s }
func int64Ptr(i int64) *int64    { return &i }
func intelligencePtr(intelligence spiritsv1.BattleTeamSpiritIntelligence) *spiritsv1.BattleTeamSpiritIntelligence {
	return &intelligence
}

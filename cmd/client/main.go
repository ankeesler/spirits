package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ankeesler/spirits0/internal/api"
	"github.com/ankeesler/spirits0/internal/menu"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/prototext"
)

type clients struct {
	spirit api.SpiritServiceClient
	battle api.BattleServiceClient
}

type state struct {
	clients *clients

	battleID string
}

var (
	port = flag.Int("port", 50051, "The server port")
)

var m = menu.Menu[*state]{
	{
		Title: "List Spirits",
		Runner: menu.RunnerFunc[*state](func(
			ctx context.Context,
			out io.Writer,
			in io.Reader,
			state *state,
		) error {
			rsp, err := state.clients.spirit.ListSpirits(ctx, &api.ListSpiritsRequest{})
			if err != nil {
				return fmt.Errorf("list spirits: %w", err)
			}

			fmt.Fprint(out, prototext.MarshalOptions{
				Multiline: true,
			}.Format(rsp))

			return nil
		}),
	},
	{
		Title: "List Battles",
		Runner: menu.RunnerFunc[*state](func(
			ctx context.Context,
			out io.Writer,
			in io.Reader,
			state *state,
		) error {
			rsp, err := state.clients.battle.ListBattles(ctx, &api.ListBattlesRequest{})
			if err != nil {
				return fmt.Errorf("list battles: %w", err)
			}

			fmt.Fprint(out, prototext.MarshalOptions{
				Multiline: true,
			}.Format(rsp))

			return nil
		}),
	},
	{
		Title: "Create Battle",
		Runner: menu.RunnerFunc[*state](func(
			ctx context.Context,
			out io.Writer,
			in io.Reader,
			s *state,
		) error {
			rsp, err := s.clients.battle.CreateBattle(ctx, &api.CreateBattleRequest{})
			if err != nil {
				return fmt.Errorf("create battle: %w", err)
			}

			s.battleID = rsp.GetBattle().GetMeta().GetId()

			fmt.Fprint(out, prototext.MarshalOptions{
				Multiline: true,
			}.Format(rsp))

			return menu.Menu[*state]{
				{
					Title: "Add Battle Team",
					Runner: menu.RunnerFunc[*state](func(
						ctx context.Context,
						out io.Writer,
						in io.Reader,
						state *state,
					) error {
						teamName, err := menu.Input(out, in, "Team name: ")
						if err != nil {
							return fmt.Errorf("input: %w", err)
						}

						rsp, err := state.clients.battle.AddBattleTeam(ctx, &api.AddBattleTeamRequest{
							BattleId: state.battleID,
							TeamName: teamName,
						})
						if err != nil {
							return fmt.Errorf("add battle team: %w", err)
						}

						fmt.Fprint(out, prototext.MarshalOptions{
							Multiline: true,
						}.Format(rsp))

						return nil
					}),
				},
			}.Run(ctx, out, in, s)
		}),
	},
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

	if err := m.Run(context.Background(), os.Stdout, os.Stdin, &state{
		clients: &clients{
			spirit: api.NewSpiritServiceClient(conn),
			battle: api.NewBattleServiceClient(conn),
		},
	}); err != nil {
		log.Fatal(err)
	}
}

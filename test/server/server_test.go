package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ankeesler/spirits/internal/server"
	"github.com/ankeesler/spirits/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const defaultTestServerPort = 12345

type clients struct {
	spirit api.SpiritServiceClient
	battle api.BattleServiceClient
}

func startServer(t *testing.T) *clients {
	port, ok := os.LookupEnv("SPIRITS_TEST_PORT")
	if !ok {
		port = fmt.Sprintf("%d", defaultTestServerPort)
		server, err := server.Wire(&server.Config{
			Port: defaultTestServerPort,

			SpiritBuiltinDir: os.DirFS("../../api/builtin/spirit"),
			ActionBuiltinDir: os.DirFS("../../api/builtin/action"),
		})
		if err != nil {
			t.Fatal(err)
		}

		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			if err := server.Serve(ctx); err != nil {
				t.Errorf("server exited with error: %v", err)
			}
		}()

		t.Cleanup(cancel)
	}

	conn, err := grpc.Dial(
		fmt.Sprintf(":%s", port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatal(err)
	}

	return &clients{
		spirit: api.NewSpiritServiceClient(conn),
		battle: api.NewBattleServiceClient(conn),
	}
}

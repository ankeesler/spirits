package test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ankeesler/spirits/internal/server"
	spiritsv1 "github.com/ankeesler/spirits/pkg/api/spirits/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const defaultTestServerPort = 12345

type state struct {
	ctx     context.Context
	clients *clients
}

type clients struct {
	spirit spiritsv1.SpiritServiceClient
	battle spiritsv1.BattleServiceClient
}

func startServer(t *testing.T) *state {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	// cancel called below in t.Cleanup()

	port, ok := os.LookupEnv("SPIRITS_TEST_PORT")
	if !ok {
		port = fmt.Sprintf("%d", defaultTestServerPort)
		server, err := server.Wire(&server.Config{
			Port: defaultTestServerPort,

			SpiritBuiltinDir: os.DirFS("../../api/builtin/v1/spirit"),
			ActionBuiltinDir: os.DirFS("../../api/builtin/v1/action"),
		})
		if err != nil {
			t.Fatal(err)
		}

		serverErrC := make(chan error, 1)
		t.Cleanup(func() {
			if serverErr := <-serverErrC; serverErr != nil {
				t.Errorf("server exited with error: %v", serverErr)
			}
		})

		go func() {
			serverErrC <- server.Serve(ctx)
		}()
	}
	t.Cleanup(cancel)

	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf(":%s", port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatal(err)
	}

	return &state{
		ctx: ctx,
		clients: &clients{
			spirit: spiritsv1.NewSpiritServiceClient(conn),
			battle: spiritsv1.NewBattleServiceClient(conn),
		},
	}
}

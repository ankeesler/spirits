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

		serverErrC := make(chan error, 1)
		ctx, cancel := context.WithCancel(context.Background())
		t.Cleanup(func() {
			cancel()
			if serverErr := <-serverErrC; serverErr != nil {
				t.Errorf("server exited with error: %v", serverErr)
			}
		})

		go func() {
			serverErrC <- server.Serve(ctx)
		}()
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

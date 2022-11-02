package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ankeesler/spirits0/internal/api"
	"github.com/ankeesler/spirits0/internal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type clients struct {
	spirit api.SpiritServiceClient
	battle api.BattleServiceClient
}

func startServer(t *testing.T) *clients {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "12345"
		os.Setenv("PORT", port)

		server, err := server.Wire(&server.Config{
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

		t.Cleanup(func() {
			cancel()

			if !ok {
				os.Unsetenv("PORT")
			}
		})
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

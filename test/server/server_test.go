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
	const port = 12345
	server, err := server.Wire(&server.Config{
		Port:             port,
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

	conn, err := grpc.Dial(
		fmt.Sprintf(":%d", port),
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

package controller

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/ankeesler/spirits/pkg/api"
	"github.com/ankeesler/spirits/internal/battle/storage/memory"
)

var r = rand.New(rand.NewSource(0))

func TestController(t *testing.T) {
	t.Skip()

	tests := []struct {
		name   string
		battle *api.Battle
	}{
		{
			name: "smoke",
			battle: &api.Battle{
				State: api.BattleState_BATTLE_STATE_STARTED,
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
			defer cancel()

			repo := memory.New(r)
			if _, err := repo.Create(ctx, test.battle, func(_ *api.Battle) error {
				return nil
			}); err != nil {
				t.Fatal("repo create:", err)
			}

			controller, err := Wire(repo, nil, nil)
			if err != nil {
				t.Fatal("wire:", err)
			}

			if err := controller.Run(ctx); err != nil {
				t.Fatal("run:", err)
			}
		})
	}
}

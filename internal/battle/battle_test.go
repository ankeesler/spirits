package battle_test

import (
	"testing"

	"github.com/ankeesler/spirits/internal/battle"
	"github.com/ankeesler/spirits/internal/spirit"
	"github.com/ankeesler/spirits/internal/team"
	"github.com/stretchr/testify/require"
)

func TestBattle(t *testing.T) {
	tests := []struct {
		name          string
		teams         []*team.Team
		wantCallbacks []interface{}
	}{}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			callbacks := callbacks{}
			b := battle.New(test.teams...)
			b.Callback = callbacks
			b.Run()
			require.Equal(t, test.wantCallbacks, callbacks)
		})
	}
}

// callbacks is a test double for battle.Battle.Callback.
type callbacks []interface{}

func (c callbacks) OnBegin(ts []*team.Team) { c = append(c, ts) }
func (c callbacks) OnTurn(s *spirit.Spirit) { c = append(c, s) }
func (c callbacks) OnEnd(t *team.Team)      { c = append(c, t) }

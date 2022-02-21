package ui_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/ankeesler/spirits/internal/spirit"
	"github.com/ankeesler/spirits/internal/tbass"
	"github.com/ankeesler/spirits/internal/tbass/activity"
	"github.com/ankeesler/spirits/internal/tbass/activity/strategy"
	"github.com/ankeesler/spirits/internal/tbass/actor/action"
	"github.com/ankeesler/spirits/internal/tbass/team"
	"github.com/ankeesler/spirits/internal/ui"
	"github.com/stretchr/testify/require"
)

func TestText(t *testing.T) {
	tests := []struct {
		name string
		tt   tbass.Teams
	}{
		{
			name: "one-on-one",
			tt: tbass.Teams{
				team.New(
					"team a",
					spirit.New("a1", 5, 1, 1, 2, action.Debuf(spirit.StatPower, spirit.StatHealth, false)),
				),
				team.New(
					"team b",
					spirit.New("b1", 5, 2, 1, 1, action.Debuf(spirit.StatPower, spirit.StatHealth, false)),
				),
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			gotOut := bytes.NewBuffer([]byte{})
			activity.New(&activity.Config{
				GetStrategy: func(tt tbass.Teams) activity.Strategy {
					return strategy.HealthAndSpeed(spirit.StatHealth, spirit.StatAgility, tt)
				},
				Listener: ui.NewText(gotOut),
			}).Play(test.tt)

			wantOut, err := os.ReadFile("testdata/" + test.name + ".txt")
			require.NoError(t, err)
			require.Equal(t, string(wantOut), gotOut.String())
		})
	}
}

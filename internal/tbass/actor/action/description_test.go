package action_test

import (
	"testing"

	"github.com/ankeesler/spirits/internal/tbass"
	"github.com/ankeesler/spirits/internal/tbass/actor"
	"github.com/ankeesler/spirits/internal/tbass/actor/action"
	"github.com/stretchr/testify/require"
)

type stat tbass.StatValue

func (s *stat) Name() string          { return "whatever" }
func (s *stat) Get() tbass.StatValue  { return tbass.StatValue(*s) }
func (s *stat) Set(v tbass.StatValue) { *s = stat(v) }

type stats map[string]*stat

func (s stats) Stat(name string) tbass.Stat { return s[name] }

func TestDescription(t *testing.T) {
	tests := []struct {
		description     string
		meIn, meOut     stats
		usIn, usOut     []stats
		themIn, themOut [][]stats
		wantError       string
	}{
		{
			description: "buf health on allied spirit with minimum health using power",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.description, func(t *testing.T) {
			a, err := action.Description(test.description)
			if test.wantError != "" {
				require.EqualError(t, err, test.wantError)
			}
			meIn := actor.New("whatever", test.meIn, a)
			usIn := team.New("whatever", meIn, 
			a.Act(meIn, usIn, themIn)
			require.Equal(t, test.meOut, test.meIn)
			require.Equal(t, test.usIn, test.usOut)
			require.Equal(t, test.themIn, test.themOut)
		})
	}
}

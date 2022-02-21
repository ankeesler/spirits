package strategy_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ankeesler/spirits/internal/tbass"
	"github.com/ankeesler/spirits/internal/tbass/activity/strategy"
	"github.com/ankeesler/spirits/internal/tbass/actor"
	"github.com/ankeesler/spirits/internal/tbass/actor/stats"
	"github.com/ankeesler/spirits/internal/tbass/team"
)

const (
	healthStat = "hp"
	speedStat  = "speed"
)

func TestHealthAndSpeed(t *testing.T) {
	t.Run("Next()", testHealthAndSpeedNext)
	t.Run("Winner()", testHealthAndSpeedWinner)
}

func testHealthAndSpeedNext(t *testing.T) {
	newActor := func(name string, speedVal tbass.StatValue) tbass.Actor {
		return actor.New(
			name,
			stats.New().With(speedStat, speedVal, 0, 0),
			nil,
		)
	}

	type next struct {
		a string
		t string
	}

	tests := []struct {
		name      string
		tt        tbass.Teams
		wantNexts []next
	}{
		{
			name: "1 vs 2",
			tt: tbass.Teams{
				team.New("a", newActor("a1", 1)),
				team.New("b", newActor("b1", 2)),
			},
			wantNexts: []next{
				{a: "b1", t: "b"},
				{a: "b1", t: "b"},
				{a: "a1", t: "a"},
				{a: "b1", t: "b"},
				{a: "b1", t: "b"},
				{a: "a1", t: "a"},
			},
		},
		{
			name: "2 vs 1",
			tt: tbass.Teams{
				team.New("a", newActor("a1", 2)),
				team.New("b", newActor("b1", 1)),
			},
			wantNexts: []next{
				{a: "a1", t: "a"},
				{a: "a1", t: "a"},
				{a: "b1", t: "b"},
				{a: "a1", t: "a"},
				{a: "a1", t: "a"},
				{a: "b1", t: "b"},
			},
		},
		{
			name: "2,1 vs 1,2",
			tt: tbass.Teams{
				team.New("a", newActor("a1", 2), newActor("a2", 1)),
				team.New("b", newActor("b1", 1), newActor("b2", 2)),
			},
			wantNexts: []next{
				{a: "a1", t: "a"},
				{a: "b2", t: "b"},
				{a: "a1", t: "a"},
				{a: "b2", t: "b"},
				{a: "a2", t: "a"},
				{a: "b1", t: "b"},
				{a: "a1", t: "a"},
				{a: "b2", t: "b"},
				{a: "a1", t: "a"},
				{a: "b2", t: "b"},
				{a: "a2", t: "a"},
				{a: "b1", t: "b"},
			},
		},
		{
			name: "1 vs 2 vs 3",
			tt: tbass.Teams{
				team.New("a", newActor("a1", 1)),
				team.New("b", newActor("b1", 2)),
				team.New("c", newActor("c1", 3)),
			},
			wantNexts: []next{
				{a: "c1", t: "c"},
				{a: "b1", t: "b"},
				{a: "c1", t: "c"},
				{a: "c1", t: "c"},
				{a: "b1", t: "b"},
				{a: "a1", t: "a"},
				{a: "c1", t: "c"},
				{a: "b1", t: "b"},
				{a: "c1", t: "c"},
				{a: "c1", t: "c"},
				{a: "b1", t: "b"},
				{a: "a1", t: "a"},
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			s := strategy.HealthAndSpeed(healthStat, speedStat, test.tt)
			for i, wantNext := range test.wantNexts {
				a, team := s.Next()
				gotNext := next{a: a.Name(), t: team.Name()}
				require.Equalf(t, wantNext, gotNext, "at index %d", i)
			}
		})
	}
}

func testHealthAndSpeedWinner(t *testing.T) {
	newActor := func(name string, healthVal tbass.StatValue) tbass.Actor {
		return actor.New(
			name,
			stats.New().With(healthStat, healthVal, 0, 0).With(speedStat, 0, 0, 0),
			nil,
		)
	}

	tests := []struct {
		name       string
		tt         tbass.Teams
		wantWinner tbass.Teams
	}{
		{
			name: "no winner",
			tt: tbass.Teams{
				team.New("a", newActor("a1", 1), newActor("a2", 2)),
				team.New("b", newActor("b1", 1), newActor("b2", 2)),
			},
		},
		{
			name: "no winner with one unhealthy actor on one team",
			tt: tbass.Teams{
				team.New("a", newActor("a1", 1), newActor("a2", 2)),
				team.New("b", newActor("b1", 1), newActor("b2", 0)),
			},
		},
		{
			name: "no winner with one unhealthy actor on each team",
			tt: tbass.Teams{
				team.New("a", newActor("a1", 0), newActor("a2", 2)),
				team.New("b", newActor("b1", 1), newActor("b2", 0)),
			},
		},
		{
			name: "winner with fully healthy team",
			tt: tbass.Teams{
				team.New("a", newActor("a1", 0), newActor("a2", 0)),
				team.New("b", newActor("b1", 1), newActor("b2", 2)),
			},
			wantWinner: tbass.Teams{
				team.New("b", newActor("b1", 1), newActor("b2", 2)),
			},
		},
		{
			name: "winner with half healthy team",
			tt: tbass.Teams{
				team.New("a", newActor("a1", 0), newActor("a2", 0)),
				team.New("b", newActor("b1", 0), newActor("b2", 2)),
			},
			wantWinner: tbass.Teams{
				team.New("b", newActor("b1", 0), newActor("b2", 2)),
			},
		},
		{
			name: "winner with no healthy teams",
			tt: tbass.Teams{
				team.New("a", newActor("a1", 0), newActor("a2", 0)),
				team.New("b", newActor("b1", 0), newActor("b2", 0)),
			},
			wantWinner: tbass.Teams{},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			s := strategy.HealthAndSpeed(healthStat, speedStat, test.tt)
			gotWinner := s.Winner()
			require.Equal(t, test.wantWinner, gotWinner)
		})
	}
}

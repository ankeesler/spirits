package test

import (
	"reflect"
	"testing"

	"github.com/ankeesler/spirits/internal/spirit"
	"github.com/ankeesler/spirits/internal/tbass"
	"github.com/ankeesler/spirits/internal/tbass/activity"
	"github.com/ankeesler/spirits/internal/tbass/activity/strategy"
	"github.com/ankeesler/spirits/internal/tbass/team"
	"github.com/google/go-cmp/cmp"
)

func TestBattle(t *testing.T) {
	cmpopts := cmp.Options{
		// cmp.Comparer(func(a, b actor.Action) bool { return true }),
		cmp.Comparer(func(a, b *spirit.Spirit) bool {
			stats := []string{
				spirit.StatHealth,
				spirit.StatPower,
				spirit.StatArmour,
				spirit.StatAgility,
			}
			for _, stat := range stats {
				if a.Stat(stat).Get() != b.Stat(stat).Get() {
					return false
				}
			}
			return true
		}),
		cmp.Exporter(func(_ reflect.Type) bool { return true }),
	}

	// t5211 := newTeam(t, "5.2.1.1")
	// t5112 := newTeam(t, "5.1.1.2")

	tests := []struct {
		name         string
		tt           tbass.Teams
		wantOnBegins []tbass.Teams
		wantOnTurns  []tbass.Actor
		wantOnEnds   []tbass.Teams
	}{
		// {
		// 	name:         "one on one",
		// 	tt:           tbass.Teams{t5211, t5112},
		// 	wantOnBegins: []tbass.Teams{{t5211, t5112}},
		// 	wantOnTurns: []tbass.Actor{
		// 		newSpirit(t, "5.1.1.2"),
		// 		newSpirit(t, "5.1.1.2"),
		// 		newSpirit(t, "3.2.1.1"),
		// 		newSpirit(t, "3.1.1.2"),
		// 		newSpirit(t, "3.1.1.2"),
		// 		newSpirit(t, "1.2.1.1"),
		// 		newSpirit(t, "1.1.1.2"),
		// 	},
		// 	wantOnEnds: []tbass.Teams{{t5112}},
		// },
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			callbacks := callbacks{}
			activity.New(&activity.Config{
				GetStrategy: func(tt tbass.Teams) activity.Strategy {
					return strategy.HealthAndSpeed(spirit.StatHealth, spirit.StatAgility, tt)
				},
				Listener: &callbacks,
			}).Play(test.tt)
			if diff := cmp.Diff(test.wantOnBegins, callbacks.onBegins, cmpopts); diff != "" {
				t.Errorf("unexpected onBegin calls; -want, +got:\n%s", diff)
			}
			if diff := cmp.Diff(test.wantOnTurns, callbacks.onTurns, cmpopts); diff != "" {
				t.Errorf("unexpected onTurns calls; -want, +got:\n%s", diff)
			}
			if diff := cmp.Diff(test.wantOnEnds, callbacks.onEnds, cmpopts); diff != "" {
				t.Errorf("unexpected onEnds calls; -want, +got:\n%s", diff)
			}
		})
	}
}

type callbacks struct {
	onBegins []tbass.Teams
	onTurns  []tbass.Actor
	onEnds   []tbass.Teams
}

func (c *callbacks) OnBegin(tt tbass.Teams) {
	c.onBegins = append(c.onBegins, deepCopyTeams(tt))
}

func (c *callbacks) OnTurn(a tbass.Actor) {
	c.onTurns = append(c.onTurns, deepCopyActor(a))
}

func (c *callbacks) OnEnd(tt tbass.Teams) {
	c.onEnds = append(c.onEnds, deepCopyTeams(tt))
}

func deepCopyTeams(tt tbass.Teams) tbass.Teams {
	ttCopy := tbass.Teams{}
	for _, t := range tt {
		ttCopy = append(ttCopy, deepCopyTeam(t))
	}
	return ttCopy
}

func deepCopyTeam(t tbass.Team) tbass.Team {
	aaCopy := []tbass.Actor{}
	for _, a := range t.Actors() {
		aaCopy = append(aaCopy, deepCopyActor(a))
	}
	return team.New(t.Name(), aaCopy...)
}

func deepCopyActor(a tbass.Actor) tbass.Actor {
	return &spirit.Spirit{Actor: a.Clone()}
}

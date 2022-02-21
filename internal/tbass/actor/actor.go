// Package actor provides a simple implementation of tbass.Actor.
package actor

import (
	"fmt"

	"github.com/ankeesler/spirits/internal/tbass"
)

// Stats are a set of tbass.Stat that can be retrieved by name.
type Stats interface {
	// Stat returns the tbass.Stat associated with the provided name, or nil if no such tbass.Stat
	// exists.
	Stat(string) tbass.Stat

	// Clone returns a deep-copy of these Stats.
	Clone() Stats
}

// Action performs a tbass.Actor's act.
type Action interface {
	// Act performs a tbass.Actor's act. The arguments are:
	//   - the tbass.Actor performing this Action
	//   - the acting tbass.Actor's tbass.Team
	//   - the opposing the tbass.Teams on which this Action is being performed
	Act(tbass.Actor, tbass.Team, tbass.Teams)
}

type actor struct {
	name   string
	stats  Stats
	action Action
}

// New creates a new tbass.Actor with the provided name, stats, and action.
func New(name string, stats Stats, action Action) tbass.Actor {
	return &actor{name: name, stats: stats, action: action}
}

func (a *actor) Name() string { return a.name }

func (a *actor) Act(tt tbass.Teams) {
	us := tt[0]
	them := make(tbass.Teams, 0, len(tt)-1)
	for i := range tt {
		if tt[i].Name() != us.Name() {
			them = append(them, tt[i])
		}
	}
	a.action.Act(a, us, them)
}

func (a *actor) Stat(name string) tbass.Stat { return a.stats.Stat(name) }

func (a *actor) Clone() tbass.Actor {
	return New(a.name, a.stats.Clone(), a.action)
}

func (a *actor) String() string {
	return fmt.Sprintf("%s: %s", a.name, a.stats)
}

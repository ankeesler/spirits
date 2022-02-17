// Package battle provides the Battle logic for the spirits project.
package battle

import (
	"github.com/ankeesler/spirits/pkg/spirit"
	"github.com/ankeesler/spirits/pkg/team"
)

// Callback is a type that wants to handle specific events that happen during the course of a
// Battle.
type Callback interface {
	// OnBegin is called before the Battle begins. A list of the team.Team's in the Battle are
	// provided.
	OnBegin([]*team.Team)
	// OnTurn is called after a turn runs, i.e., after a single spirit.Spirit performs a
	// spirit.Action. The spirit.Spirit that just went is provided.
	OnTurn(*spirit.Spirit)
	// OnEnd is called after the Battle finishes. The winning team.Team is provided.
	OnEnd(*team.Team)
}

// Battle orchestrates a Battle amongst team.Team's.
type Battle struct {
	// Teams are the teams.Team's involved in the Battle.
	Teams []*team.Team

	// Callback is an optional field that can be used to notify an object of events that happen over
	// the course of a Battle.
	Callback Callback
}

// New instantiates a new Battle amongst the provided team.Team's.
func New(teams ...*team.Team) *Battle {
	return &Battle{
		Teams: teams,
	}
}

// Run actually runs the battle. Once this function returns, the Battle has ended.
//
// It is recommended that Battle.Callback is utilized to determine the winner of the Battle.
func (b *Battle) Run() {
	for {
		break
	}
}

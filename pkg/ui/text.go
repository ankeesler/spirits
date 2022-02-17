package ui

import (
	"io"

	"github.com/ankeesler/spirits/pkg/battle"
	"github.com/ankeesler/spirits/pkg/spirit"
	"github.com/ankeesler/spirits/pkg/team"
)

type text struct {
	w io.Writer

	teams []*team.Team
	turns int
}

// Text returns a text-based UI. This UI is good for command line applications, and simple streaming
// applications.
func Text(w io.Writer) battle.Callback {
	return &text{w: w}
}

func (t *text) OnBegin(teams []*team.Team) {
}

func (t *text) OnTurn(spirit *spirit.Spirit) {
}

func (t *text) OnEnd(team *team.Team) {
}

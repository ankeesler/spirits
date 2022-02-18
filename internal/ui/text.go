package ui

import (
	"io"

	"github.com/ankeesler/spirits/internal/spirit"
	"github.com/ankeesler/spirits/internal/team"
)

// Text is a text-based UI.
type Text struct {
	w io.Writer

	teams []*team.Team
	turns int
}

// Text returns a text-based UI. This UI is good for command line applications, and simple streaming
// applications.
func NewText(w io.Writer) *Text {
	return &Text{w: w}
}

func (t *Text) OnBegin(teams []*team.Team) {
}

func (t *Text) OnTurn(spirit *spirit.Spirit) {
}

func (t *Text) OnEnd(team *team.Team) {
}

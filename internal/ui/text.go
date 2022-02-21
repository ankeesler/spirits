package ui

import (
	"fmt"
	"io"
	"strings"

	"github.com/ankeesler/spirits/internal/spirit"
	"github.com/ankeesler/spirits/internal/tbass"
)

// Text is a text-based UI.
type Text struct {
	w io.Writer

	tt tbass.Teams

	maxHealths map[tbass.Actor]tbass.StatValue
}

// Text returns a text-based UI. This UI is good for command line applications, and simple streaming
// applications.
func NewText(w io.Writer) *Text {
	return &Text{w: w, maxHealths: make(map[tbass.Actor]tbass.StatValue)}
}

func (t *Text) OnBegin(tt tbass.Teams) {
	t.tt = tt

	for _, a := range tt.Actors() {
		t.maxHealths[a] = a.Stat(spirit.StatHealth).Get()
	}

	fmt.Fprint(t.w, "battle\n\n")
	t.printSummary()
}

func (t *Text) OnTurn(a tbass.Actor) {
	fmt.Fprintf(t.w, "> %s attacked\n\n", a.Name())
	t.printSummary()
}

func (t *Text) OnEnd(tt tbass.Teams) {
	fmt.Fprintf(t.w, "> %s wins\n", tt[0].Name())
}

func (t *Text) printSummary() {
	fmt.Fprintf(t.w, "> summary\n")

	for i := range t.tt {
		fmt.Fprintf(t.w, "\t\t%s\n", t.tt[i].Name())

		for j := range t.tt[i].Actors() {
			a := t.tt[i].Actors()[j]

			const healthWidth = 50
			currentHealth := a.Stat(spirit.StatHealth).Get()
			maxHealth := t.maxHealths[a]
			healthBars := currentHealth * healthWidth / maxHealth
			fmt.Fprintf(
				t.w,
				"\t\t\t%s\t\t[%s%s]\n",
				a.Name(),
				strings.Repeat("=", int(healthBars)),
				strings.Repeat(" ", int(healthWidth-healthBars)),
			)
		}
	}

	fmt.Fprintln(t.w)
}

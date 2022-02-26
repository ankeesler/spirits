package ui

import (
	"fmt"
	"io"

	"github.com/ankeesler/spirits/internal/spirit"
)

type UI struct {
	out io.Writer
}

func New(out io.Writer) *UI {
	return &UI{
		out: out,
	}
}
func (u *UI) OnSpirits(spirits []*spirit.Spirit) {
	fmt.Fprintln(u.out, "> summary")
	for _, spirit := range spirits {
		fmt.Fprintf(u.out, "  %s: %d\n", spirit.Name, spirit.Health)
	}
}

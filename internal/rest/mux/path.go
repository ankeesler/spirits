package mux

import (
	"net/http"

	"github.com/ankeesler/spirits/internal/rest"
)

type Path struct {
	Name     *string
	Handler  http.Handler
	Children []*Path
}

var _ rest.Handler = &Path{}

func (p *Path) Handle(w http.ResponseWriter, r *http.Request) error {
	return nil
}

package mux

import (
	"net/http"

	"github.com/ankeesler/spirits/internal/rest"
)

type Method map[string]rest.Handler

var _ rest.Handler = Method{}

func (m Method) Handle(w http.ResponseWriter, r *http.Request) error {
	handler, ok := m[r.Method]
	if !ok {
		return &rest.Err{Code: http.StatusMethodNotAllowed}
	}
	return handler.Handle(w, r)
}

package auth

import (
	"net/http"

	"github.com/ankeesler/spirits/internal/rest"
)

type OIDC struct {
	next rest.Handler
}

var _ rest.Handler = &OIDC{}

func NewOIDC(next rest.Handler) *OIDC {
	return &OIDC{
		next: next,
	}
}

func (a *OIDC) Handle(w http.ResponseWriter, r *http.Request) error {
	return a.next.Handle(w, r)
}

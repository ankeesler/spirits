package domain

import (
	"github.com/ankeesler/spirits/internal/domain/session"
	"github.com/ankeesler/spirits/internal/store"
)

type Domain struct {
	Sessions *store.Store[session.Session]
}

func New() *Domain {
	return &Domain{
		Sessions: store.New(func(session *session.Session) string { return session.Name }),
	}
}

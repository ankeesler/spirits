package session

import (
	"github.com/ankeesler/spirits/internal/domain/team"
	"github.com/ankeesler/spirits/internal/store"
)

type Session struct {
	Name string

	Teams *store.Store[team.Team]
	// Sessions *store.Store[battle.Battle]
}

func New(name string) *Session {
	return &Session{
		Name:  name,
		Teams: store.New(func(team *team.Team) string { return team.Name }),
	}
}

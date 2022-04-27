package session

import (
	"github.com/ankeesler/spirits/internal/domain/battle"
	"github.com/ankeesler/spirits/internal/domain/team"
	"github.com/ankeesler/spirits/internal/store"
)

type Session struct {
	Name string

	Teams   *store.Store[team.Team]
	Battles *store.Store[battle.Battle]
}

func New(name string) *Session {
	return &Session{
		Name:    name,
		Teams:   store.New(func(team *team.Team) string { return team.Name }),
		Battles: store.New(func(battle *battle.Battle) string { return battle.Name }),
	}
}

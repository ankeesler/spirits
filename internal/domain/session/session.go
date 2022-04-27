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

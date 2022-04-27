package team

import (
	"github.com/ankeesler/spirits/internal/domain/spirit"
	"github.com/ankeesler/spirits/internal/store"
)

type Team struct {
	Name    string
	Spirits *store.Store[spirit.Spirit]
}

func New(name string) *Team {
	return &Team{
		Name:    name,
		Spirits: store.New(func(spirit *spirit.Spirit) string { return spirit.Name }),
	}
}

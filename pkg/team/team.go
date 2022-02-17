// Package team provides the concept of a group of spirit.Spirit's working together to win a
// battle.Battle.
package team

import (
	"github.com/ankeesler/spirits/pkg/spirit"
)

type Team struct {
	Name    string
	Spirits []*spirit.Spirit
}

func New(name string, spirits []*spirit.Spirit) *Team {
	return &Team{
		Name:    name,
		Spirits: spirits,
	}
}

// Package team provides a simple implementation of a tbass.Team.
package team

import (
	"github.com/ankeesler/spirits/internal/tbass"
)

type team struct {
	name string
	aa   []tbass.Actor
}

func New(name string, aa ...tbass.Actor) tbass.Team {
	return &team{
		name: name,
		aa:   aa,
	}
}

func (t *team) Name() string          { return t.name }
func (t *team) Actors() []tbass.Actor { return t.aa }

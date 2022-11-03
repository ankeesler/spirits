package action

import (
	"context"
	"errors"

	"github.com/ankeesler/spirits0/internal/api"
	"github.com/ankeesler/spirits0/internal/spirit"
)

type SpiritStats interface {
	Health() int64
	SetHealth(int64)

	PhysicalPower() int64
	SetPhysicalPower(int64)
	PhysicalConstitution() int64
	SetPhysicalConstitution(int64)

	MentalPower() int64
	SetMentalPower(int64)
	MentalConstitution() int64
	SetMentalConstitution(int64)

	Agility() int64
	SetAgility(int64)
}

type Spirit interface {
	ID() string
	Name() string
	Stats() SpiritStats
}

type Action interface {
	Run(context.Context, *spirit.Spirit, []*spirit.Spirit) (context.Context, error)
}

func FromAPI(apiAction *api.Action) (Action, error) {
	switch definition := apiAction.Definition.(type) {
	case *api.Action_Script:
		return compile(definition.Script)
	default:
		return nil, errors.New("definition not set")
	}
}

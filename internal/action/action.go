package action

import (
	"context"
	"errors"
	"fmt"

	"github.com/ankeesler/spirits/internal/meta"
	metapkg "github.com/ankeesler/spirits/internal/meta"
	"github.com/ankeesler/spirits/pkg/api"
)

type Spirit interface {
	ID() string
	Name() string

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

type caller interface {
	Call(context.Context, Spirit, []Spirit) (context.Context, error)
}

type Action struct {
	apiAction *api.Action

	*meta.Meta

	description string
	script      *string

	caller
}

func FromAPI(apiAction *api.Action) (*Action, error) {
	internalAction := &Action{
		apiAction: apiAction,

		Meta: metapkg.FromAPI(apiAction.GetMeta()),

		description: apiAction.GetDescription(),
	}

	switch definition := apiAction.Definition.(type) {
	case *api.Action_Script:
		internalAction.script = &definition.Script
		compiledScript, err := compile(definition.Script)
		if err != nil {
			return nil, fmt.Errorf("compile action script: %w", err)
		}
		internalAction.caller = compiledScript
	default:
		return nil, errors.New("definition not set")
	}

	return internalAction, nil
}

func (a *Action) ToAPI() *api.Action {
	return a.apiAction
}

func (a *Action) Clone() *Action {
	return &Action{
		Meta: a.Meta.Clone(),

		description: a.description,
		script:      a.script,

		caller: a.caller,
	}
}

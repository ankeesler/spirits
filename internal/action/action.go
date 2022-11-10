package action

import (
	"context"

	metapkg "github.com/ankeesler/spirits/internal/meta"
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
	*metapkg.Meta

	description string
	script      *string

	caller
}

func New(meta *metapkg.Meta) *Action {
	return &Action{Meta: meta}
}

func (a *Action) Description() string               { return a.description }
func (a *Action) SetDescription(description string) { a.description = description }

func (a *Action) Script() *string { return a.script }
func (a *Action) SetScript(script string) error {
	a.script = &script

	compiledScript, err := compile(script)
	if err != nil {
		return err
	}
	a.caller = compiledScript

	return nil
}

func (a *Action) Clone() *Action {
	return &Action{
		Meta: a.Meta.Clone(),

		description: a.description,
		script:      a.script,

		caller: a.caller,
	}
}

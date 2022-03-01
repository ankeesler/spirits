package action

import "github.com/ankeesler/spirits/internal/spirit"

type action struct {
	name string
	run  func(from, to *spirit.Spirit)
}

func (a *action) Name() string                { return a.name }
func (a *action) Run(from, to *spirit.Spirit) { a.run(from, to) }

func Attack() spirit.Action {
	return &action{
		name: "attack",
		run: func(from, to *spirit.Spirit) {
			netPower := (from.Power - to.Armour)
			if netPower < 0 {
				netPower = 0
			}

			to.Health -= netPower
			if to.Health < 0 {
				to.Health = 0
			}
		},
	}
}

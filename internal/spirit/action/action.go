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

func Bolster() spirit.Action {
	return &action{
		name: "bolster",
		run: func(from, to *spirit.Spirit) {
			netPower := ((from.Power / 2) - to.Armour)
			if netPower < 0 {
				netPower = 0
			}

			to.Health -= netPower
			if to.Health < 0 {
				to.Health = 0
			}

			from.Armour += (from.Power / 2)
		},
	}
}

func Drain() spirit.Action {
	return &action{
		name: "drain",
		run: func(from, to *spirit.Spirit) {
			netPower := ((from.Power / 2) - to.Armour)
			if netPower < 0 {
				netPower = 0
			}

			to.Health -= netPower
			if to.Health < 0 {
				to.Health = 0
			}

			from.Health += (netPower / 2)
		},
	}
}

func Charge() spirit.Action {
	return &action{
		name: "charge",
		run: func(from, to *spirit.Spirit) {
			netPower := ((from.Power * 2) - to.Armour)
			if netPower < 0 {
				netPower = 0
			}

			to.Health -= netPower
			if to.Health < 0 {
				to.Health = 0
			}

			from.Health -= (netPower / 2)
			if from.Health < 0 {
				from.Health = 0
			}
		},
	}
}

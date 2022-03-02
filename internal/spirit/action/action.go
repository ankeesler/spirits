package action

import "github.com/ankeesler/spirits/internal/spirit"

type actionFunc func(from, to *spirit.Spirit)

func (f actionFunc) Run(from, to *spirit.Spirit) { f(from, to) }

func Attack() spirit.Action {
	return actionFunc(func(from, to *spirit.Spirit) {
		netPower := (from.Power - to.Armour)
		if netPower < 0 {
			netPower = 0
		}

		to.Health -= netPower
		if to.Health < 0 {
			to.Health = 0
		}
	})
}

func Bolster() spirit.Action {
	return actionFunc(func(from, to *spirit.Spirit) {
		netPower := ((from.Power / 2) - to.Armour)
		if netPower < 0 {
			netPower = 0
		}

		to.Health -= netPower
		if to.Health < 0 {
			to.Health = 0
		}

		from.Armour += (from.Power / 2)
	})
}

func Drain() spirit.Action {
	return actionFunc(func(from, to *spirit.Spirit) {
		netPower := ((from.Power / 2) - to.Armour)
		if netPower < 0 {
			netPower = 0
		}

		to.Health -= netPower
		if to.Health < 0 {
			to.Health = 0
		}

		from.Health += (netPower / 2)
	})
}

func Charge() spirit.Action {
	return actionFunc(func(from, to *spirit.Spirit) {
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
	})
}

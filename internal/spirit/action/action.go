package action

import (
	"context"

	"github.com/ankeesler/spirits/api/internal/spirit"
)

type actionFunc func(ctx context.Context, from, to *spirit.Spirit) error

func (f actionFunc) Run(ctx context.Context, from, to *spirit.Spirit) error {
	return f(ctx, from, to)
}

func Attack() spirit.Action {
	return actionFunc(func(ctx context.Context, from, to *spirit.Spirit) error {
		netPower := (from.Power - to.Armor)
		if netPower < 0 {
			netPower = 0
		}

		to.Health -= netPower
		if to.Health < 0 {
			to.Health = 0
		}

		return nil
	})
}

func Bolster() spirit.Action {
	return actionFunc(func(ctx context.Context, from, to *spirit.Spirit) error {
		netPower := ((from.Power / 2) - to.Armor)
		if netPower < 0 {
			netPower = 0
		}

		to.Health -= netPower
		if to.Health < 0 {
			to.Health = 0
		}

		from.Armor += (from.Power / 2)

		return nil
	})
}

func Drain() spirit.Action {
	return actionFunc(func(ctx context.Context, from, to *spirit.Spirit) error {
		netPower := ((from.Power / 2) - to.Armor)
		if netPower < 0 {
			netPower = 0
		}

		to.Health -= netPower
		if to.Health < 0 {
			to.Health = 0
		}

		from.Health += (netPower / 2)

		return nil
	})
}

func Charge() spirit.Action {
	return actionFunc(func(ctx context.Context, from, to *spirit.Spirit) error {
		netPower := ((from.Power * 2) - to.Armor)
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

		return nil
	})
}

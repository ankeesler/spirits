package action

import (
	"context"

	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
)

type actionFunc func(ctx context.Context, from, to *spiritsinternal.Spirit) error

func (f actionFunc) Run(ctx context.Context, from, to *spiritsinternal.Spirit) error {
	return f(ctx, from, to)
}

func (f actionFunc) DeepCopyAction() spiritsinternal.Action { return f }

func Attack() spiritsinternal.Action {
	return actionFunc(func(ctx context.Context, from, to *spiritsinternal.Spirit) error {
		netPower := (from.Spec.Stats.Power - to.Spec.Stats.Armor)
		if netPower < 0 {
			netPower = 0
		}

		to.Spec.Stats.Health -= netPower
		if to.Spec.Stats.Health < 0 {
			to.Spec.Stats.Health = 0
		}

		return nil
	})
}

func Bolster() spiritsinternal.Action {
	return actionFunc(func(ctx context.Context, from, to *spiritsinternal.Spirit) error {
		netPower := ((from.Spec.Stats.Power / 2) - to.Spec.Stats.Armor)
		if netPower < 0 {
			netPower = 0
		}

		to.Spec.Stats.Health -= netPower
		if to.Spec.Stats.Health < 0 {
			to.Spec.Stats.Health = 0
		}

		from.Spec.Stats.Armor += (from.Spec.Stats.Power / 2)

		return nil
	})
}

func Drain() spiritsinternal.Action {
	return actionFunc(func(ctx context.Context, from, to *spiritsinternal.Spirit) error {
		netPower := ((from.Spec.Stats.Power / 2) - to.Spec.Stats.Armor)
		if netPower < 0 {
			netPower = 0
		}

		to.Spec.Stats.Health -= netPower
		if to.Spec.Stats.Health < 0 {
			to.Spec.Stats.Health = 0
		}

		from.Spec.Stats.Health += (netPower / 2)

		return nil
	})
}

func Charge() spiritsinternal.Action {
	return actionFunc(func(ctx context.Context, from, to *spiritsinternal.Spirit) error {
		netPower := ((from.Spec.Stats.Power * 2) - to.Spec.Stats.Armor)
		if netPower < 0 {
			netPower = 0
		}

		to.Spec.Stats.Health -= netPower
		if to.Spec.Stats.Health < 0 {
			to.Spec.Stats.Health = 0
		}

		from.Spec.Stats.Health -= (netPower / 2)
		if from.Spec.Stats.Health < 0 {
			from.Spec.Stats.Health = 0
		}

		return nil
	})
}

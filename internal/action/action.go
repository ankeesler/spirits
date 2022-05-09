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
		netPower := (from.Spec.Attributes.Stats.Power - to.Spec.Attributes.Stats.Armor)
		if netPower < 0 {
			netPower = 0
		}

		to.Spec.Attributes.Stats.Health -= netPower
		if to.Spec.Attributes.Stats.Health < 0 {
			to.Spec.Attributes.Stats.Health = 0
		}

		return nil
	})
}

func Noop() spiritsinternal.Action {
	return actionFunc(func(ctx context.Context, from, to *spiritsinternal.Spirit) error { return nil })
}

func Bolster() spiritsinternal.Action {
	return actionFunc(func(ctx context.Context, from, to *spiritsinternal.Spirit) error {
		netPower := ((from.Spec.Attributes.Stats.Power / 2) - to.Spec.Attributes.Stats.Armor)
		if netPower < 0 {
			netPower = 0
		}

		to.Spec.Attributes.Stats.Health -= netPower
		if to.Spec.Attributes.Stats.Health < 0 {
			to.Spec.Attributes.Stats.Health = 0
		}

		from.Spec.Attributes.Stats.Armor += (from.Spec.Attributes.Stats.Power / 2)

		return nil
	})
}

func Drain() spiritsinternal.Action {
	return actionFunc(func(ctx context.Context, from, to *spiritsinternal.Spirit) error {
		netPower := ((from.Spec.Attributes.Stats.Power / 2) - to.Spec.Attributes.Stats.Armor)
		if netPower < 0 {
			netPower = 0
		}

		to.Spec.Attributes.Stats.Health -= netPower
		if to.Spec.Attributes.Stats.Health < 0 {
			to.Spec.Attributes.Stats.Health = 0
		}

		from.Spec.Attributes.Stats.Health += (netPower / 2)

		return nil
	})
}

func Charge() spiritsinternal.Action {
	return actionFunc(func(ctx context.Context, from, to *spiritsinternal.Spirit) error {
		netPower := ((from.Spec.Attributes.Stats.Power * 2) - to.Spec.Attributes.Stats.Armor)
		if netPower < 0 {
			netPower = 0
		}

		to.Spec.Attributes.Stats.Health -= netPower
		if to.Spec.Attributes.Stats.Health < 0 {
			to.Spec.Attributes.Stats.Health = 0
		}

		from.Spec.Attributes.Stats.Health -= (netPower / 2)
		if from.Spec.Attributes.Stats.Health < 0 {
			from.Spec.Attributes.Stats.Health = 0
		}

		return nil
	})
}

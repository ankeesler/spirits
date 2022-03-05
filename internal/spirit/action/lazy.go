package action

import (
	"context"

	"github.com/ankeesler/spirits/internal/spirit"
)

func Lazy(f func(ctx context.Context) (spirit.Action, error)) spirit.Action {
	return actionFunc(func(ctx context.Context, from, to *spirit.Spirit) error {
		action, err := f(ctx)
		if err != nil {
			return err
		}
		return action.Run(ctx, from, to)
	})
}

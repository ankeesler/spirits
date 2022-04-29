package action

import (
	"context"

	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
)

func Lazy(f func(ctx context.Context) (spiritsinternal.Action, error)) spiritsinternal.Action {
	return actionFunc(func(ctx context.Context, from, to *spiritsinternal.Spirit) error {
		action, err := f(ctx)
		if err != nil {
			return err
		}
		return action.Run(ctx, from, to)
	})
}

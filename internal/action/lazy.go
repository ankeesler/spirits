package action

import (
	"context"
	"fmt"

	"github.com/ankeesler/spirits/internal/spirit"
)

type lazy struct {
	f func(context.Context) (Action, error)
}

func Lazy(f func(context.Context) (Action, error)) Action {
	return &lazy{f: f}
}

func (l *lazy) Run(
	ctx context.Context,
	source *spirit.Spirit,
	targets []*spirit.Spirit,
) (context.Context, error) {
	action, err := l.f(ctx)
	if err != nil {
		return ctx, fmt.Errorf("lazy action: %w", err)
	}
	return action.Run(ctx, source, targets)
}

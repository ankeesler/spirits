package runner

import "context"

type Runner struct {
}

func New() *Runner {
	return &Runner{}
}

func (r *Runner) Run(ctx context.Context) error {
	return nil
}

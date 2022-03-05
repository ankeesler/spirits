package spirit

import "context"

type Action interface {
	Run(ctx context.Context, to, from *Spirit) error
}

type Spirit struct {
	Name    string
	Health  int
	Power   int
	Agility int
	Armor   int
	Action  Action
}

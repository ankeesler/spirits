package spirits

import "context"

type Action interface {
	Run(context.Context, *Spirit, *Spirit) error
	DeepCopyAction() Action
}

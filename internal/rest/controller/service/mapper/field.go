package mapper

import (
	"context"

	"github.com/ankeesler/spirits/internal/rest/controller/service"
)

type Field[A, B any] struct {
}

func NewField[A, B any]() *Field[A, B] {
	return &Field[A, B]{}
}

var _ service.Mapper[struct{}, struct{}] = &Field[struct{}, struct{}]{}

func (c *Field[A, B]) AToB(ctx context.Context, a *A) (*B, error) {
	return nil, nil
}

func (c *Field[A, B]) BToA(ctx context.Context, b *B) (*A, error) {
	return nil, nil
}

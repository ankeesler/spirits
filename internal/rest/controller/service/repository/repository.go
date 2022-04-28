package repository

import (
	"context"

	"github.com/ankeesler/spirits/internal/rest/controller/service"
)

type Repository[T any] struct {
}

var _ service.Repository[struct{}] = &Repository[struct{}]{}

func New[T any]() *Repository[T] {
	return &Repository[T]{}
}

func (r *Repository[T]) Create(context.Context, string, *T) (*T, error) { return nil, nil }
func (r *Repository[T]) Update(context.Context, string, *T) (*T, error) { return nil, nil }
func (r *Repository[T]) List(context.Context, string) ([]*T, error)     { return nil, nil }
func (r *Repository[T]) Get(context.Context, string) (*T, error)        { return nil, nil }
func (r *Repository[T]) Delete(context.Context, string) (*T, error)     { return nil, nil }

package validater

import (
	"errors"
	"reflect"

	"github.com/ankeesler/spirits/internal/rest/controller/serializer"
)

type Constraint interface {
	Fits(reflect.Value) bool
}

type Field[T any] struct {
	path       string
	constraint Constraint
}

var _ serializer.Validater[struct{}] = &Field[struct{}]{}

func (f *Field[T]) Validate(t *T) error {
	value, err := f.findField(t)
	if err != nil {
		return err
	}

	if !f.constraint.Fits(value) {
		return errors.New("constraint does not fit value")
	}

	return nil
}

func (f *Field[T]) findField(t *T) (reflect.Value, error) {
	// TODO: walk type and find field
	return reflect.ValueOf(t), nil
}

func NewField[T any](path string, constraint Constraint) *Field[T] {
	return &Field[T]{
		path:       path,
		constraint: constraint,
	}
}

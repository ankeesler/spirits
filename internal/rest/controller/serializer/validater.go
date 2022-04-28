package serializer

import "github.com/ankeesler/spirits/internal/rest"

type Validater[T any] interface {
	Validate(*T) error
}

func WithValidater[T any](validater Validater[T]) rest.Option[Serializer[T]] {
	return rest.Option[Serializer[T]](func(s *Serializer[T]) {
		s.validaters = append(s.validaters, validater)
	})
}

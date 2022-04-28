package rest

type Option[T any] func(*T)

type Options[T any] []Option[T]

func (oo Options[T]) Apply(t *T) {
	for _, o := range oo {
		o(t)
	}
}

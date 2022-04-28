package serializer

import (
	"io"
	"net/http"

	"github.com/ankeesler/spirits/internal/rest"
	"github.com/ankeesler/spirits/internal/rest/controller"
)

type serializer[T any] interface {
	Write(io.Writer, *T) error
	Read(io.Reader, *T) error
}

type Serializer[T any] struct {
	mimeTypes  map[string]serializer[T]
	validaters []Validater[T]
}

var _ controller.Serializer[struct{}] = &Serializer[struct{}]{}

func New[T any](options rest.Options[Serializer[T]]) *Serializer[T] {
	s := &Serializer[T]{
		mimeTypes: make(map[string]serializer[T]),
	}
	options.Apply(s)
	return s
}

func (s *Serializer[T]) Write(w http.ResponseWriter, r *http.Request, t *T) error {
	if err := s.validate(t); err != nil {
		return err
	}

	if err := s.write(w, r, t); err != nil {
		return err
	}

	return nil
}

func (s *Serializer[T]) Read(w http.ResponseWriter, r *http.Request, t *T) error {
	if err := s.read(w, r, t); err != nil {
		return err
	}

	if err := s.validate(t); err != nil {
		return err
	}

	return nil
}

func (s *Serializer[T]) write(w http.ResponseWriter, r *http.Request, t *T) error {
	return nil
}

func (s *Serializer[T]) read(w http.ResponseWriter, r *http.Request, t *T) error {
	return nil
}

func (s *Serializer[T]) validate(t *T) error {
	for _, validater := range s.validaters {
		if err := validater.Validate(t); err != nil {
			return err
		}
	}
	return nil
}

package serializer

import (
	"encoding/json"
	"io"

	"github.com/ankeesler/spirits/internal/rest"
)

type jsonSerializer[T any] struct {
}

func (j *jsonSerializer[T]) Write(w io.Writer, t *T) error {
	return json.NewEncoder(w).Encode(t)
}

func (j *jsonSerializer[T]) Read(r io.Reader, t *T) error {
	return json.NewDecoder(r).Decode(t)
}

func WithJSON[T any]() rest.Option[Serializer[T]] {
	return rest.Option[Serializer[T]](func(s *Serializer[T]) {
		s.mimeTypes["application/json"] = &jsonSerializer[T]{}
	})
}

package controller

import (
	"context"
	"net/http"

	"github.com/ankeesler/spirits/internal/rest"
)

type Serializer[T any] interface {
	Write(http.ResponseWriter, *http.Request, *T) error
	Read(http.ResponseWriter, *http.Request, *T) error
}

type Service[ExternalT, InternalT any] interface {
	Create(context.Context, *ExternalT) (*ExternalT, error)
	Update(context.Context, *ExternalT) (*ExternalT, error)
	List(context.Context) ([]*ExternalT, error)
	Get(context.Context) (*ExternalT, error)
	// TODO: watch
	Delete(context.Context) (*ExternalT, error)
}

type Controller[ExternalT, InternalT any] struct {
	serializer Serializer[ExternalT]
	service    Service[ExternalT, InternalT]
}

func New[ExternalT, InternalT any](
	serializer Serializer[ExternalT],
	service Service[ExternalT, InternalT],
) *Controller[ExternalT, InternalT] {
	return &Controller[ExternalT, InternalT]{
		serializer: serializer,
		service:    service,
	}
}

var _ rest.Handler = rest.HandlerFunc((&Controller[struct{}, struct{}]{}).Create)

func (c *Controller[ExternalT, InternalT]) Create(w http.ResponseWriter, r *http.Request) error {
	r = xxx(r)

	var reqT ExternalT
	if err := c.serializer.Read(w, r, &reqT); err != nil {
		return err
	}

	rspT, err := c.service.Create(r.Context(), &reqT)
	if err != nil {
		return err
	}

	if err := c.serializer.Write(w, r, rspT); err != nil {
		return err
	}

	return nil
}

var _ rest.Handler = rest.HandlerFunc((&Controller[struct{}, struct{}]{}).Update)

func (c *Controller[ExternalT, InternalT]) Update(w http.ResponseWriter, r *http.Request) error {
	r = xxx(r)

	var reqT ExternalT
	if err := c.serializer.Read(w, r, &reqT); err != nil {
		return err
	}

	rspT, err := c.service.Create(r.Context(), &reqT)
	if err != nil {
		return err
	}

	if err := c.serializer.Write(w, r, rspT); err != nil {
		return err
	}

	return nil
}

var _ rest.Handler = rest.HandlerFunc((&Controller[struct{}, struct{}]{}).List)

func (c *Controller[ExternalT, InternalT]) List(w http.ResponseWriter, r *http.Request) error {
	r = xxx(r)

	rspTs, err := c.service.List(r.Context())
	if err != nil {
		return err
	}

	// TODO: this is wrong...
	for _, rspT := range rspTs {
		if err := c.serializer.Write(w, r, rspT); err != nil {
			return err
		}
	}

	return nil
}

var _ rest.Handler = rest.HandlerFunc((&Controller[struct{}, struct{}]{}).Get)

func (c *Controller[ExternalT, InternalT]) Get(w http.ResponseWriter, r *http.Request) error {
	r = xxx(r)

	rspT, err := c.service.Get(r.Context())
	if err != nil {
		return err
	}

	if err := c.serializer.Write(w, r, rspT); err != nil {
		return err
	}

	return nil
}

var _ rest.Handler = rest.HandlerFunc((&Controller[struct{}, struct{}]{}).Delete)

func (c *Controller[ExternalT, InternalT]) Delete(w http.ResponseWriter, r *http.Request) error {
	r = xxx(r)

	rspT, err := c.service.Delete(r.Context())
	if err != nil {
		return err
	}

	if err := c.serializer.Write(w, r, rspT); err != nil {
		return err
	}

	return nil
}

func xxx(r *http.Request) *http.Request {
	r = r.WithContext(rest.WithPath(r.Context(), r.URL.Path))
	return r
}

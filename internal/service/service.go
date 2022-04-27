package service

import (
	"context"
	"errors"
	"net/http"

	storepkg "github.com/ankeesler/spirits/internal/store"
	server "github.com/ankeesler/spirits/pkg/api/generated/server/api"
)

type ConverterFuncs[FromT any, ToT any] struct {
	From func(from *FromT) (*ToT, error)
	To   func(to *ToT) (*FromT, error)
}

func Create[FromT any, ToT any](
	ctx context.Context,
	from *FromT,
	store *storepkg.Store[ToT],
	converter *ConverterFuncs[FromT, ToT],
) (server.ImplResponse, error) {
	to, err := converter.From(from)
	if err != nil {
		return server.ImplResponse{
			Code: http.StatusBadRequest,
			Body: err.Error(),
		}, nil
	}

	to, err = store.Create(ctx, to)
	if err != nil {
		if errors.Is(err, &storepkg.ErrAlreadyExists{}) {
			return server.ImplResponse{
				Code: http.StatusConflict,
			}, nil
		}
		return server.ImplResponse{}, err
	}

	from, err = converter.To(to)
	if err != nil {
		return server.ImplResponse{}, err
	}

	return server.ImplResponse{
		Code: http.StatusCreated,
		Body: from,
	}, nil
}

func Update[FromT any, ToT any](
	ctx context.Context,
	from *FromT,
	store *storepkg.Store[ToT],
	converter *ConverterFuncs[FromT, ToT],
) (server.ImplResponse, error) {
	to, err := converter.From(from)
	if err != nil {
		return server.ImplResponse{
			Code: http.StatusBadRequest,
			Body: err.Error(),
		}, nil
	}

	to, err = store.Update(ctx, to)
	if err != nil {
		if errors.Is(err, &storepkg.ErrNotFound{}) {
			return server.ImplResponse{
				Code: http.StatusNotFound,
			}, nil
		}
		return server.ImplResponse{}, err
	}

	from, err = converter.To(to)
	if err != nil {
		return server.ImplResponse{}, err
	}

	return server.ImplResponse{
		Code: http.StatusOK,
		Body: from,
	}, nil
}

func List[FromT any, ToT any](
	ctx context.Context,
	store *storepkg.Store[ToT],
	converter *ConverterFuncs[FromT, ToT],
) (server.ImplResponse, error) {
	tos, err := store.List(ctx)
	if err != nil {
		return server.ImplResponse{}, err
	}

	froms := []*FromT{}
	for _, to := range tos {
		from, err := converter.To(to)
		if err != nil {
			return server.ImplResponse{}, err
		}
		froms = append(froms, from)
	}

	return server.ImplResponse{
		Code: http.StatusOK,
		Body: froms,
	}, nil
}

func Get[FromT any, ToT any](
	ctx context.Context,
	name string,
	store *storepkg.Store[ToT],
	converter *ConverterFuncs[FromT, ToT],
) (server.ImplResponse, error) {
	to, err := store.Get(ctx, name)
	if err != nil {
		if errors.Is(err, &storepkg.ErrNotFound{}) {
			return server.ImplResponse{
				Code: http.StatusNotFound,
			}, nil
		}
		return server.ImplResponse{}, err
	}

	from, err := converter.To(to)
	if err != nil {
		return server.ImplResponse{}, err
	}

	return server.ImplResponse{
		Code: http.StatusOK,
		Body: from,
	}, nil
}

func Delete[FromT any, ToT any](
	ctx context.Context,
	name string,
	store *storepkg.Store[ToT],
	converter *ConverterFuncs[FromT, ToT],
) (server.ImplResponse, error) {
	to, err := store.Delete(ctx, name)
	if err != nil {
		if errors.Is(err, &storepkg.ErrNotFound{}) {
			return server.ImplResponse{
				Code: http.StatusNotFound,
			}, nil
		}
		return server.ImplResponse{}, err
	}

	from, err := converter.To(to)
	if err != nil {
		return server.ImplResponse{}, err
	}

	return server.ImplResponse{
		Code: http.StatusOK,
		Body: from,
	}, nil
}

func Find[T any](ctx context.Context, name string, store *storepkg.Store[T]) (*T, server.ImplResponse, error) {
	t, err := store.Get(ctx, name)
	if err != nil {
		if errors.Is(err, &storepkg.ErrNotFound{}) {
			return nil, server.ImplResponse{
				Code: http.StatusNotFound,
			}, nil
		}
		return nil, server.ImplResponse{}, err
	}
	return t, server.ImplResponse{}, nil
}

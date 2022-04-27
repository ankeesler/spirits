package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/ankeesler/spirits/internal/log"
	storepkg "github.com/ankeesler/spirits/internal/store"
	server "github.com/ankeesler/spirits/pkg/api/generated/server/api"
)

type ConverterFuncs[A any, B any] struct {
	AToB func(from *A) (*B, error)
	BToA func(to *B) (*A, error)
}

func Create[FromT any, ToT any](
	ctx context.Context,
	from *FromT,
	store *storepkg.Store[ToT],
	converter *ConverterFuncs[FromT, ToT],
) (server.ImplResponse, error) {
	log.Debug(fmt.Sprintf("service: create %#v: begin", from))
	defer log.Debug(fmt.Sprintf("service: create %#v: end", from))

	to, err := converter.AToB(from)
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
				Body: err.Error(),
			}, nil
		}
		return server.ImplResponse{}, err
	}

	from, err = converter.BToA(to)
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
	log.Debug(fmt.Sprintf("service: update %#v: begin", from))
	defer log.Debug(fmt.Sprintf("service: update %#v: end", from))

	to, err := converter.AToB(from)
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
				Body: err.Error(),
			}, nil
		}
		return server.ImplResponse{}, err
	}

	from, err = converter.BToA(to)
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
	log.Debug(fmt.Sprintf("service: list: begin"))
	defer log.Debug(fmt.Sprintf("service: update: end"))

	tos, err := store.List(ctx)
	if err != nil {
		return server.ImplResponse{}, err
	}

	froms := []*FromT{}
	for _, to := range tos {
		from, err := converter.BToA(to)
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
	log.Debug(fmt.Sprintf("service: get %q: begin", name))
	defer log.Debug(fmt.Sprintf("service: get %q: end", name))

	to, err := store.Get(ctx, name)
	if err != nil {
		if errors.Is(err, &storepkg.ErrNotFound{}) {
			return server.ImplResponse{
				Code: http.StatusNotFound,
				Body: err.Error(),
			}, nil
		}
		return server.ImplResponse{}, err
	}

	from, err := converter.BToA(to)
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
	log.Debug(fmt.Sprintf("service: get %q: begin", name))
	defer log.Debug(fmt.Sprintf("service: get %q: end", name))

	to, err := store.Delete(ctx, name)
	if err != nil {
		if errors.Is(err, &storepkg.ErrNotFound{}) {
			return server.ImplResponse{
				Code: http.StatusNotFound,
				Body: err.Error(),
			}, nil
		}
		return server.ImplResponse{}, err
	}

	from, err := converter.BToA(to)
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

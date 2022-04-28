package service

import (
	"context"

	"github.com/ankeesler/spirits/internal/rest"
	"github.com/ankeesler/spirits/internal/rest/controller"
)

type Mapper[A, B any] interface {
	AToB(context.Context, *A) (*B, error)
	BToA(context.Context, *B) (*A, error)
}

type Repository[T any] interface {
	Create(context.Context, string, *T) (*T, error)
	Update(context.Context, string, *T) (*T, error)
	List(context.Context, string) ([]*T, error)
	Get(context.Context, string) (*T, error)
	// TODO: watchw
	Delete(context.Context, string) (*T, error)
}

type Service[ExternalT, InternalT any] struct {
	mapper     Mapper[ExternalT, InternalT]
	repository Repository[InternalT]
}

var _ controller.Service[struct{}, struct{}] = &Service[struct{}, struct{}]{}

func New[ExternalT, InternalT any](
	mapper Mapper[ExternalT, InternalT],
	repository Repository[InternalT],
) *Service[ExternalT, InternalT] {
	return &Service[ExternalT, InternalT]{
		mapper:     mapper,
		repository: repository,
	}
}

func (s *Service[ExternalT, Internal]) Create(ctx context.Context, externalT *ExternalT) (*ExternalT, error) {
	internalT, err := s.mapper.AToB(ctx, externalT)
	if err != nil {
		return nil, &rest.Err{ /* TODO: bad request */ }
	}

	internalT, err = s.repository.Create(ctx, rest.Path(ctx), internalT)
	if err != nil {
		return nil, err
	}

	externalT, err = s.mapper.BToA(ctx, internalT)
	if err != nil {
		return nil, err
	}

	return externalT, nil
}

func (s *Service[ExternalT, Internal]) Update(ctx context.Context, externalT *ExternalT) (*ExternalT, error) {
	internalT, err := s.mapper.AToB(ctx, externalT)
	if err != nil {
		return nil, &rest.Err{ /* TODO: bad request */ }
	}

	internalT, err = s.repository.Update(ctx, rest.Path(ctx), internalT)
	if err != nil {
		return nil, err
	}

	externalT, err = s.mapper.BToA(ctx, internalT)
	if err != nil {
		return nil, err
	}

	return externalT, nil
}

func (s *Service[ExternalT, InternalT]) List(ctx context.Context) ([]*ExternalT, error) {
	internalTs, err := s.repository.List(ctx, rest.Path(ctx))
	if err != nil {
		return nil, err
	}

	externalTs := []*ExternalT{}
	for _, internalT := range internalTs {
		externalT, err := s.mapper.BToA(ctx, internalT)
		if err != nil {
			return nil, err
		}
		externalTs = append(externalTs, externalT)
	}

	return externalTs, nil
}

func (s *Service[ExternalT, InternalT]) Get(ctx context.Context) (*ExternalT, error) {
	internalT, err := s.repository.Get(ctx, rest.Path(ctx))
	if err != nil {
		return nil, err
	}

	externalT, err := s.mapper.BToA(ctx, internalT)
	if err != nil {
		return nil, err
	}

	return externalT, nil
}

func (s *Service[ExternalT, InternalT]) Delete(ctx context.Context) (*ExternalT, error) {
	internalT, err := s.repository.Delete(ctx, rest.Path(ctx))
	if err != nil {
		return nil, err
	}

	externalT, err := s.mapper.BToA(ctx, internalT)
	if err != nil {
		return nil, err
	}

	return externalT, nil
}

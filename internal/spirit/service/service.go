package service

import (
	"context"

	spiritpkg "github.com/ankeesler/spirits/internal/spirit"
	convertspirit "github.com/ankeesler/spirits/internal/spirit/convert"
	"github.com/ankeesler/spirits/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SpiritRepo interface {
	Create(context.Context, *spiritpkg.Spirit) (*spiritpkg.Spirit, error)
	Get(context.Context, string) (*spiritpkg.Spirit, error)
	Update(context.Context, *spiritpkg.Spirit) (*spiritpkg.Spirit, error)
	Delete(context.Context, string) (*spiritpkg.Spirit, error)

	ListSpirits(context.Context, *string) ([]*spiritpkg.Spirit, error)
}

type Service struct {
	spiritRepo SpiritRepo

	api.UnimplementedSpiritServiceServer
}

var _ api.SpiritServiceServer = &Service{}

func New(spiritRepo SpiritRepo) *Service {
	return &Service{
		spiritRepo: spiritRepo,
	}
}

func (s *Service) CreateSpirit(
	ctx context.Context,
	req *api.CreateSpiritRequest,
) (*api.CreateSpiritResponse, error) {
	internalSpirit, err := convertspirit.FromAPI(req.GetSpirit())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	internalSpirit, err = s.spiritRepo.Create(ctx, internalSpirit)
	if err != nil {
		return nil, err
	}

	return &api.CreateSpiritResponse{Spirit: convertspirit.ToAPI(internalSpirit)}, nil
}

func (s *Service) GetSpirit(
	ctx context.Context,
	req *api.GetSpiritRequest,
) (*api.GetSpiritResponse, error) {
	internalSpirit, err := s.spiritRepo.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &api.GetSpiritResponse{Spirit: convertspirit.ToAPI(internalSpirit)}, nil
}

func (s *Service) ListSpirits(
	ctx context.Context,
	req *api.ListSpiritsRequest,
) (*api.ListSpiritsResponse, error) {
	internalSpirits, err := s.spiritRepo.ListSpirits(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	var apiSpirits []*api.Spirit
	for _, internalSpirit := range internalSpirits {
		apiSpirits = append(apiSpirits, convertspirit.ToAPI(internalSpirit))
	}

	return &api.ListSpiritsResponse{Spirits: apiSpirits}, nil
}

func (s *Service) UpdateSpirit(
	ctx context.Context,
	req *api.UpdateSpiritRequest,
) (*api.UpdateSpiritResponse, error) {
	internalSpirit, err := convertspirit.FromAPI(req.GetSpirit())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	internalSpirit, err = s.spiritRepo.Update(ctx, internalSpirit)
	if err != nil {
		return nil, err
	}

	return &api.UpdateSpiritResponse{Spirit: convertspirit.ToAPI(internalSpirit)}, nil
}

func (s *Service) DeleteSpirit(
	ctx context.Context,
	req *api.DeleteSpiritRequest,
) (*api.DeleteSpiritResponse, error) {
	internalSpirit, err := s.spiritRepo.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &api.DeleteSpiritResponse{Spirit: convertspirit.ToAPI(internalSpirit)}, nil
}

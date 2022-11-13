package service

import (
	"context"

	spiritpkg "github.com/ankeesler/spirits/internal/spirit"
	convertspirit "github.com/ankeesler/spirits/internal/spirit/convert"
	"github.com/ankeesler/spirits/pkg/api/spirits/v1"
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

	spiritsv1.UnimplementedSpiritServiceServer
}

var _ spiritsv1.SpiritServiceServer = &Service{}

func New(spiritRepo SpiritRepo) *Service {
	return &Service{
		spiritRepo: spiritRepo,
	}
}

func (s *Service) CreateSpirit(
	ctx context.Context,
	req *spiritsv1.CreateSpiritRequest,
) (*spiritsv1.CreateSpiritResponse, error) {
	internalSpirit, err := convertspirit.FromAPI(req.GetSpirit())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	internalSpirit, err = s.spiritRepo.Create(ctx, internalSpirit)
	if err != nil {
		return nil, err
	}

	return &spiritsv1.CreateSpiritResponse{Spirit: convertspirit.ToAPI(internalSpirit)}, nil
}

func (s *Service) GetSpirit(
	ctx context.Context,
	req *spiritsv1.GetSpiritRequest,
) (*spiritsv1.GetSpiritResponse, error) {
	internalSpirit, err := s.spiritRepo.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &spiritsv1.GetSpiritResponse{Spirit: convertspirit.ToAPI(internalSpirit)}, nil
}

func (s *Service) ListSpirits(
	ctx context.Context,
	req *spiritsv1.ListSpiritsRequest,
) (*spiritsv1.ListSpiritsResponse, error) {
	internalSpirits, err := s.spiritRepo.ListSpirits(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	var apiSpirits []*spiritsv1.Spirit
	for _, internalSpirit := range internalSpirits {
		apiSpirits = append(apiSpirits, convertspirit.ToAPI(internalSpirit))
	}

	return &spiritsv1.ListSpiritsResponse{Spirits: apiSpirits}, nil
}

func (s *Service) UpdateSpirit(
	ctx context.Context,
	req *spiritsv1.UpdateSpiritRequest,
) (*spiritsv1.UpdateSpiritResponse, error) {
	internalSpirit, err := convertspirit.FromAPI(req.GetSpirit())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	internalSpirit, err = s.spiritRepo.Update(ctx, internalSpirit)
	if err != nil {
		return nil, err
	}

	return &spiritsv1.UpdateSpiritResponse{Spirit: convertspirit.ToAPI(internalSpirit)}, nil
}

func (s *Service) DeleteSpirit(
	ctx context.Context,
	req *spiritsv1.DeleteSpiritRequest,
) (*spiritsv1.DeleteSpiritResponse, error) {
	internalSpirit, err := s.spiritRepo.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &spiritsv1.DeleteSpiritResponse{Spirit: convertspirit.ToAPI(internalSpirit)}, nil
}

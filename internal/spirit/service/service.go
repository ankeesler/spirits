package service

import (
	"context"

	actionpkg "github.com/ankeesler/spirits/internal/action"
	spiritpkg "github.com/ankeesler/spirits/internal/spirit"
	"github.com/ankeesler/spirits/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ActionRepo interface {
	Get(context.Context, string) (*actionpkg.Action, error)
}

type SpiritRepo interface {
	Create(context.Context, *spiritpkg.Spirit) (*spiritpkg.Spirit, error)
	Get(context.Context, string) (*spiritpkg.Spirit, error)
	Update(context.Context, *spiritpkg.Spirit) (*spiritpkg.Spirit, error)
	Delete(context.Context, string) (*spiritpkg.Spirit, error)

	ListSpirits(context.Context, *string) ([]*spiritpkg.Spirit, error)
}

type fakeActionSource struct{}

func (s fakeActionSource) Pend(
	context.Context, *spiritpkg.Spirit,
	[]*spiritpkg.Spirit, [][]*spiritpkg.Spirit) (string, []string, error) {
	return "", nil, nil
}

type Service struct {
	spiritRepo SpiritRepo
	actionRepo ActionRepo

	api.UnimplementedSpiritServiceServer
}

var _ api.SpiritServiceServer = &Service{}

func New(spiritRepo SpiritRepo, actionRepo ActionRepo) *Service {
	return &Service{
		spiritRepo: spiritRepo,
		actionRepo: actionRepo,
	}
}

func (s *Service) CreateSpirit(
	ctx context.Context,
	req *api.CreateSpiritRequest,
) (*api.CreateSpiritResponse, error) {
	internalSpirit, err := spiritpkg.FromAPI(ctx, req.GetSpirit(), s.actionRepo, fakeActionSource{})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	internalSpirit, err = s.spiritRepo.Create(ctx, internalSpirit)
	if err != nil {
		return nil, err
	}

	return &api.CreateSpiritResponse{Spirit: internalSpirit.ToAPI()}, nil
}

func (s *Service) GetSpirit(
	ctx context.Context,
	req *api.GetSpiritRequest,
) (*api.GetSpiritResponse, error) {
	internalSpirit, err := s.spiritRepo.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &api.GetSpiritResponse{Spirit: internalSpirit.ToAPI()}, nil
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
		apiSpirits = append(apiSpirits, internalSpirit.ToAPI())
	}

	return &api.ListSpiritsResponse{Spirits: apiSpirits}, nil
}

func (s *Service) UpdateSpirit(
	ctx context.Context,
	req *api.UpdateSpiritRequest,
) (*api.UpdateSpiritResponse, error) {
	internalSpirit, err := spiritpkg.FromAPI(ctx, req.GetSpirit(), s.actionRepo, fakeActionSource{})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	internalSpirit, err = s.spiritRepo.Update(ctx, internalSpirit)
	if err != nil {
		return nil, err
	}

	return &api.UpdateSpiritResponse{Spirit: internalSpirit.ToAPI()}, nil
}

func (s *Service) DeleteSpirit(
	ctx context.Context,
	req *api.DeleteSpiritRequest,
) (*api.DeleteSpiritResponse, error) {
	internalSpirit, err := s.spiritRepo.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &api.DeleteSpiritResponse{Spirit: internalSpirit.ToAPI()}, nil
}

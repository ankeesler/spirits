package service

import (
	"context"

	"github.com/ankeesler/spirits0/internal/api"
	spiritpkg "github.com/ankeesler/spirits0/internal/spirit"
)

type Repo interface {
	Create(context.Context, *api.Spirit, func(*api.Spirit) error) (*api.Spirit, error)
	Get(context.Context, string) (*api.Spirit, error)
	List(context.Context) ([]*api.Spirit, error)
	Update(context.Context, *api.Spirit, func(*api.Spirit) error) (*api.Spirit, error)
	Delete(context.Context, string) (*api.Spirit, error)
}

type Service struct {
	repo Repo

	api.UnimplementedSpiritServiceServer
}

var _ api.SpiritServiceServer = &Service{}

func New(repo Repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateSpirit(
	ctx context.Context,
	req *api.CreateSpiritRequest,
) (*api.CreateSpiritResponse, error) {
	spirit := req.GetSpirit()
	spirit.Meta = &api.Meta{}
	spirit, err := s.repo.Create(ctx, spirit, func(spirit *api.Spirit) error {
		_, err := spiritpkg.FromAPI(spirit)
		return err
	})
	if err != nil {
		return nil, err
	}
	return &api.CreateSpiritResponse{Spirit: spirit}, nil
}

func (s *Service) GetSpirit(
	ctx context.Context,
	req *api.GetSpiritRequest,
) (*api.GetSpiritResponse, error) {
	spirit, err := s.repo.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &api.GetSpiritResponse{Spirit: spirit}, nil
}

func (s *Service) ListSpirits(
	ctx context.Context,
	req *api.ListSpiritsRequest,
) (*api.ListSpiritsResponse, error) {
	spirits, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return &api.ListSpiritsResponse{Spirits: spirits}, nil
}

func (s *Service) UpdateSpirit(
	ctx context.Context,
	req *api.UpdateSpiritRequest,
) (*api.UpdateSpiritResponse, error) {
	spirit, err := s.repo.Update(ctx, req.GetSpirit(), func(spirit *api.Spirit) error {
		_, err := spiritpkg.FromAPI(spirit)
		return err
	})
	if err != nil {
		return nil, err
	}
	return &api.UpdateSpiritResponse{Spirit: spirit}, nil
}

func (s *Service) DeleteSpirit(
	ctx context.Context,
	req *api.DeleteSpiritRequest,
) (*api.DeleteSpiritResponse, error) {
	spirit, err := s.repo.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &api.DeleteSpiritResponse{Spirit: spirit}, nil
}

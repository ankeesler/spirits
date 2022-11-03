package service

import (
	"context"
	"errors"
	"fmt"

	actionpkg "github.com/ankeesler/spirits0/internal/action"
	"github.com/ankeesler/spirits0/internal/api"
)

type ActionRepo interface {
	Get(context.Context, string) (*api.Action, error)
}

type SpiritRepo interface {
	Create(context.Context, *api.Spirit, func(*api.Spirit) error) (*api.Spirit, error)
	Get(context.Context, string) (*api.Spirit, error)
	Update(context.Context, *api.Spirit, func(*api.Spirit) error) (*api.Spirit, error)
	Delete(context.Context, string) (*api.Spirit, error)

	ListSpirits(context.Context, *string) ([]*api.Spirit, error)
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
	spirit := req.GetSpirit()
	spirit.Meta = &api.Meta{}
	spirit, err := s.spiritRepo.Create(ctx, spirit, func(spirit *api.Spirit) error {
		return s.validateSpirit(ctx, spirit)
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
	spirit, err := s.spiritRepo.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &api.GetSpiritResponse{Spirit: spirit}, nil
}

func (s *Service) ListSpirits(
	ctx context.Context,
	req *api.ListSpiritsRequest,
) (*api.ListSpiritsResponse, error) {
	var name *string
	if req.Name != nil {
		name = stringPtr(req.GetName())
	}
	spirits, err := s.spiritRepo.ListSpirits(ctx, name)
	if err != nil {
		return nil, err
	}
	return &api.ListSpiritsResponse{Spirits: spirits}, nil
}

func (s *Service) UpdateSpirit(
	ctx context.Context,
	req *api.UpdateSpiritRequest,
) (*api.UpdateSpiritResponse, error) {
	spirit, err := s.spiritRepo.Update(ctx, req.GetSpirit(), func(spirit *api.Spirit) error {
		return s.validateSpirit(ctx, spirit)
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
	spirit, err := s.spiritRepo.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &api.DeleteSpiritResponse{Spirit: spirit}, nil
}

func stringPtr(s string) *string { return &s }

func (s *Service) validateSpirit(ctx context.Context, spirit *api.Spirit) error {
	if len(spirit.GetName()) == 0 {
		return errors.New("name cannot be empty")
	}

	if spirit.GetStats().GetHealth() <= 0 {
		return errors.New("health must be greater than 0")
	}

	if spirit.GetStats().GetAgility() <= 0 {
		return errors.New("agility must be greater than 0")
	}

	actionNames := make(map[string]struct{})
	for _, action := range spirit.GetActions() {
		if _, ok := actionNames[action.GetName()]; ok {
			return fmt.Errorf("duplicate action name: %s", action.GetName())
		}
		actionNames[action.GetName()] = struct{}{}

		switch definition := action.Definition.(type) {
		case *api.SpiritAction_ActionId:
			if _, err := s.actionRepo.Get(ctx, definition.ActionId); err != nil {
				return fmt.Errorf("invalid action ID for %s: %w", action.GetName(), err)
			}
		case *api.SpiritAction_Inline:
			if _, err := actionpkg.FromAPI(definition.Inline); err != nil {
				return fmt.Errorf("invalid action inline for %s: %w", action.GetName(), err)
			}
		default:
			return fmt.Errorf("must provide definition for action: %s", action.GetName())
		}
	}

	return nil
}

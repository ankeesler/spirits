package service

import (
	"context"

	actionpkg "github.com/ankeesler/spirits0/internal/action"
	"github.com/ankeesler/spirits0/internal/api"
)

type Repo interface {
	Create(context.Context, *api.Action, func(*api.Action) error) (*api.Action, error)
	Get(context.Context, string) (*api.Action, error)
	List(context.Context) ([]*api.Action, error)
	Update(context.Context, *api.Action, func(*api.Action) error) (*api.Action, error)
	Delete(context.Context, string) (*api.Action, error)
}

type Sink interface {
	Post(context.Context, string, string, string, []string) error
}

type Service struct {
	repo Repo
	sink Sink

	api.UnimplementedActionServiceServer
}

var _ api.ActionServiceServer = &Service{}

func New(repo Repo, sink Sink) *Service {
	return &Service{
		repo: repo,
		sink: sink,
	}
}

func (s *Service) CreateAction(
	ctx context.Context,
	req *api.CreateActionRequest,
) (*api.CreateActionResponse, error) {
	action, err := s.repo.Create(ctx, req.GetAction(), func(action *api.Action) error {
		_, err := actionpkg.FromAPI(action)
		return err
	})
	if err != nil {
		return nil, err
	}
	return &api.CreateActionResponse{Action: action}, nil
}

func (s *Service) GetAction(
	ctx context.Context,
	req *api.GetActionRequest,
) (*api.GetActionResponse, error) {
	action, err := s.repo.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &api.GetActionResponse{Action: action}, nil
}

func (s *Service) ListActions(
	ctx context.Context,
	req *api.ListActionsRequest,
) (*api.ListActionsResponse, error) {
	actions, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return &api.ListActionsResponse{Actions: actions}, nil
}

func (s *Service) UpdateAction(
	ctx context.Context,
	req *api.UpdateActionRequest,
) (*api.UpdateActionResponse, error) {
	action, err := s.repo.Update(ctx, req.GetAction(), func(action *api.Action) error {
		_, err := actionpkg.FromAPI(action)
		return err
	})
	if err != nil {
		return nil, err
	}
	return &api.UpdateActionResponse{Action: action}, nil
}

func (s *Service) DeleteAction(
	ctx context.Context,
	req *api.DeleteActionRequest,
) (*api.DeleteActionResponse, error) {
	action, err := s.repo.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &api.DeleteActionResponse{Action: action}, nil
}

func (s *Service) CallAction(
	ctx context.Context,
	req *api.CallActionRequest,
) (*api.CallActionResponse, error) {
	if err := s.sink.Post(
		ctx,
		req.GetBattleId(),
		req.GetSpiritId(),
		req.GetActionName(),
		req.GetTargetSpiritIds(),
	); err != nil {
		return nil, err
	}
	return &api.CallActionResponse{}, nil
}

package service

import (
	"context"

	actionpkg "github.com/ankeesler/spirits/internal/action"
	"github.com/ankeesler/spirits/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Repo interface {
	Create(context.Context, *actionpkg.Action) (*actionpkg.Action, error)
	Get(context.Context, string) (*actionpkg.Action, error)
	List(context.Context) ([]*actionpkg.Action, error)
	Update(context.Context, *actionpkg.Action) (*actionpkg.Action, error)
	Delete(context.Context, string) (*actionpkg.Action, error)
}

type Service struct {
	repo Repo

	api.UnimplementedActionServiceServer
}

var _ api.ActionServiceServer = &Service{}

func New(repo Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateAction(
	ctx context.Context,
	req *api.CreateActionRequest,
) (*api.CreateActionResponse, error) {
	internalAction, err := actionpkg.FromAPI(req.GetAction())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	internalAction, err = s.repo.Create(ctx, internalAction)
	if err != nil {
		return nil, err
	}

	return &api.CreateActionResponse{Action: internalAction.ToAPI()}, nil
}

func (s *Service) GetAction(
	ctx context.Context,
	req *api.GetActionRequest,
) (*api.GetActionResponse, error) {
	internalAction, err := s.repo.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &api.GetActionResponse{Action: internalAction.ToAPI()}, nil
}

func (s *Service) ListActions(
	ctx context.Context,
	req *api.ListActionsRequest,
) (*api.ListActionsResponse, error) {
	internalActions, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	var apiActions []*api.Action
	for _, internalAction := range internalActions {
		apiActions = append(apiActions, internalAction.ToAPI())
	}

	return &api.ListActionsResponse{Actions: apiActions}, nil
}

func (s *Service) UpdateAction(
	ctx context.Context,
	req *api.UpdateActionRequest,
) (*api.UpdateActionResponse, error) {
	internalAction, err := actionpkg.FromAPI(req.GetAction())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	internalAction, err = s.repo.Update(ctx, internalAction)
	if err != nil {
		return nil, err
	}

	return &api.UpdateActionResponse{Action: internalAction.ToAPI()}, nil
}

func (s *Service) DeleteAction(
	ctx context.Context,
	req *api.DeleteActionRequest,
) (*api.DeleteActionResponse, error) {
	internalAction, err := s.repo.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &api.DeleteActionResponse{Action: internalAction.ToAPI()}, nil
}

package service

import (
	"context"

	actionpkg "github.com/ankeesler/spirits/internal/action"
	convertaction "github.com/ankeesler/spirits/internal/action/convert"
	"github.com/ankeesler/spirits/pkg/api/spirits/v1"
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

	spiritsv1.UnimplementedActionServiceServer
}

var _ spiritsv1.ActionServiceServer = &Service{}

func New(repo Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateAction(
	ctx context.Context,
	req *spiritsv1.CreateActionRequest,
) (*spiritsv1.CreateActionResponse, error) {
	internalAction, err := convertaction.FromAPI(req.GetAction())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	internalAction, err = s.repo.Create(ctx, internalAction)
	if err != nil {
		return nil, err
	}

	return &spiritsv1.CreateActionResponse{Action: convertaction.ToAPI(internalAction)}, nil
}

func (s *Service) GetAction(
	ctx context.Context,
	req *spiritsv1.GetActionRequest,
) (*spiritsv1.GetActionResponse, error) {
	internalAction, err := s.repo.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &spiritsv1.GetActionResponse{Action: convertaction.ToAPI(internalAction)}, nil
}

func (s *Service) ListActions(
	ctx context.Context,
	req *spiritsv1.ListActionsRequest,
) (*spiritsv1.ListActionsResponse, error) {
	internalActions, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	var apiActions []*spiritsv1.Action
	for _, internalAction := range internalActions {
		apiActions = append(apiActions, convertaction.ToAPI(internalAction))
	}

	return &spiritsv1.ListActionsResponse{Actions: apiActions}, nil
}

func (s *Service) UpdateAction(
	ctx context.Context,
	req *spiritsv1.UpdateActionRequest,
) (*spiritsv1.UpdateActionResponse, error) {
	internalAction, err := convertaction.FromAPI(req.GetAction())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	internalAction, err = s.repo.Update(ctx, internalAction)
	if err != nil {
		return nil, err
	}

	return &spiritsv1.UpdateActionResponse{Action: convertaction.ToAPI(internalAction)}, nil
}

func (s *Service) DeleteAction(
	ctx context.Context,
	req *spiritsv1.DeleteActionRequest,
) (*spiritsv1.DeleteActionResponse, error) {
	internalAction, err := s.repo.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &spiritsv1.DeleteActionResponse{Action: convertaction.ToAPI(internalAction)}, nil
}

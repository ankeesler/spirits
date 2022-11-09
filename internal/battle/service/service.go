package service

import (
	"context"

	battlepkg "github.com/ankeesler/spirits/internal/battle"
	spiritpkg "github.com/ankeesler/spirits/internal/spirit"
	"github.com/ankeesler/spirits/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SpiritRepo interface {
	Get(context.Context, string) (*spiritpkg.Spirit, error)
}

type BattleRepo interface {
	Create(context.Context, *battlepkg.Battle) (*battlepkg.Battle, error)
	Watch(context.Context, chan<- *battlepkg.Battle) error
	List(context.Context) ([]*battlepkg.Battle, error)

	AddBattleTeam(context.Context, string, string) (*battlepkg.Battle, error)
	AddBattleTeamSpirit(context.Context, string, string, *spiritpkg.Spirit) (*battlepkg.Battle, error)
	UpdateBattleState(context.Context, string, []battlepkg.State, battlepkg.State) (*battlepkg.Battle, error)
}

type Service struct {
	battleRepo   BattleRepo
	spiritRepo   SpiritRepo
	actionSource spiritpkg.ActionSource

	api.UnimplementedBattleServiceServer
}

var _ api.BattleServiceServer = &Service{}

func New(
	battleRepo BattleRepo, spiritRepo SpiritRepo, actionSource spiritpkg.ActionSource) *Service {
	return &Service{
		battleRepo:   battleRepo,
		spiritRepo:   spiritRepo,
		actionSource: actionSource,
	}
}

func (s *Service) CreateBattle(
	ctx context.Context,
	req *api.CreateBattleRequest,
) (*api.CreateBattleResponse, error) {
	apiBattle := &api.Battle{
		State: api.BattleState_BATTLE_STATE_PENDING,
	}
	internalBattle, err := battlepkg.FromAPI(
		ctx, apiBattle, nil, s.actionSource)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	internalBattle, err = s.battleRepo.Create(ctx, internalBattle)
	if err != nil {
		return nil, err
	}

	return &api.CreateBattleResponse{Battle: internalBattle.ToAPI()}, nil
}

func (s *Service) WatchBattle(
	req *api.WatchBattleRequest,
	watch api.BattleService_WatchBattleServer,
) error {
	c := make(chan *battlepkg.Battle)
	defer close(c)

	if err := s.battleRepo.Watch(watch.Context(), c); err != nil {
		return err
	}

	for battle := range c {
		if battle.ID() == req.GetId() {
			if err := watch.Send(&api.WatchBattleResponse{Battle: battle.ToAPI()}); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Service) ListBattles(
	ctx context.Context,
	req *api.ListBattlesRequest,
) (*api.ListBattlesResponse, error) {
	internalBattles, err := s.battleRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	var apiBattles []*api.Battle
	for _, internalBattle := range internalBattles {
		apiBattles = append(apiBattles, internalBattle.ToAPI())
	}
	return &api.ListBattlesResponse{Battles: apiBattles}, nil
}

func (s *Service) AddBattleTeam(
	ctx context.Context,
	req *api.AddBattleTeamRequest,
) (*api.AddBattleTeamResponse, error) {
	internalBattle, err := s.battleRepo.AddBattleTeam(ctx, req.GetBattleId(), req.GetTeamName())
	if err != nil {
		return nil, err
	}
	return &api.AddBattleTeamResponse{Battle: internalBattle.ToAPI()}, nil
}

func (s *Service) AddBattleTeamSpirit(
	ctx context.Context,
	req *api.AddBattleTeamSpiritRequest,
) (*api.AddBattleTeamSpiritResponse, error) {
	internalSpirit, err := s.spiritRepo.Get(ctx, req.GetSpiritId())
	if err != nil {
		return nil, err
	}

	internalBattle, err := s.battleRepo.AddBattleTeamSpirit(
		ctx, req.GetBattleId(), req.GetTeamName(), internalSpirit)
	if err != nil {
		return nil, err
	}

	return &api.AddBattleTeamSpiritResponse{Battle: internalBattle.ToAPI()}, nil
}

func (s *Service) StartBattle(
	ctx context.Context,
	req *api.StartBattleRequest,
) (*api.StartBattleResponse, error) {
	internalBattle, err := s.battleRepo.UpdateBattleState(
		ctx,
		req.GetId(),
		[]battlepkg.State{battlepkg.StatePending},
		battlepkg.StateStarted,
	)
	if err != nil {
		return nil, err
	}
	return &api.StartBattleResponse{Battle: internalBattle.ToAPI()}, nil
}

func (s *Service) CancelBattle(
	ctx context.Context,
	req *api.CancelBattleRequest,
) (*api.CancelBattleResponse, error) {
	internalBattle, err := s.battleRepo.UpdateBattleState(
		ctx,
		req.GetId(),
		[]battlepkg.State{battlepkg.StateStarted, battlepkg.StateWaiting},
		battlepkg.StateCancelled,
	)
	if err != nil {
		return nil, err
	}
	return &api.CancelBattleResponse{Battle: internalBattle.ToAPI()}, nil
}

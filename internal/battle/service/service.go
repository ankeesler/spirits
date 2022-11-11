package service

import (
	"context"

	battlepkg "github.com/ankeesler/spirits/internal/battle"
	convertbattle "github.com/ankeesler/spirits/internal/battle/convert"
	"github.com/ankeesler/spirits/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BattleRepo interface {
	Create(context.Context, *battlepkg.Battle) (*battlepkg.Battle, error)
	Watch(context.Context) (<-chan *battlepkg.Battle, error)
	List(context.Context) ([]*battlepkg.Battle, error)

	AddBattleTeam(context.Context, string, string) (*battlepkg.Battle, error)
	AddBattleTeamSpirit(
		context.Context,
		string,
		string,
		string,
		battlepkg.SpiritIntelligence,
		int64,
		battlepkg.ActionSource) (*battlepkg.Battle, error)
	UpdateBattleState(context.Context, string, []battlepkg.State, battlepkg.State) (*battlepkg.Battle, error)
}

type Service struct {
	battleRepo   BattleRepo
	actionSource battlepkg.ActionSource

	api.UnimplementedBattleServiceServer
}

var _ api.BattleServiceServer = &Service{}

func New(battleRepo BattleRepo, actionSource battlepkg.ActionSource) *Service {
	return &Service{
		battleRepo:   battleRepo,
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
	internalBattle, err := convertbattle.FromAPI(apiBattle, s.actionSource)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	internalBattle, err = s.battleRepo.Create(ctx, internalBattle)
	if err != nil {
		return nil, err
	}

	return &api.CreateBattleResponse{Battle: convertbattle.ToAPI(internalBattle)}, nil
}

func (s *Service) WatchBattle(
	req *api.WatchBattleRequest,
	watch api.BattleService_WatchBattleServer,
) error {
	c, err := s.battleRepo.Watch(watch.Context())
	if err != nil {
		return err
	}

	for battle := range c {
		if battle.ID() == req.GetId() {
			if err := watch.Send(&api.WatchBattleResponse{
				Battle: convertbattle.ToAPI(battle),
			}); err != nil {
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
		apiBattles = append(apiBattles, convertbattle.ToAPI(internalBattle))
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
	return &api.AddBattleTeamResponse{Battle: convertbattle.ToAPI(internalBattle)}, nil
}

func (s *Service) AddBattleTeamSpirit(
	ctx context.Context,
	req *api.AddBattleTeamSpiritRequest,
) (*api.AddBattleTeamSpiritResponse, error) {
	internalBattle, err := s.battleRepo.AddBattleTeamSpirit(
		ctx, req.GetBattleId(), req.GetTeamName(), req.GetSpiritId(),
		convertbattle.SpiritIntelligenceFromAPI(req.GetIntelligence()), req.GetSeed(),
		s.actionSource)
	if err != nil {
		return nil, err
	}

	return &api.AddBattleTeamSpiritResponse{Battle: convertbattle.ToAPI(internalBattle)}, nil
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
	return &api.StartBattleResponse{Battle: convertbattle.ToAPI(internalBattle)}, nil
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
	return &api.CancelBattleResponse{Battle: convertbattle.ToAPI(internalBattle)}, nil
}

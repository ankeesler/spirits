package service

import (
	"context"

	battlepkg "github.com/ankeesler/spirits/internal/battle"
	convertbattle "github.com/ankeesler/spirits/internal/battle/convert"
	spiritsv1 "github.com/ankeesler/spirits/pkg/api/spirits/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BattleRepo interface {
	Create(context.Context, *battlepkg.Battle) (*battlepkg.Battle, error)
	Watch(context.Context, *string) (<-chan *battlepkg.Battle, error)
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

type ActionSink interface {
	Post(context.Context, string, string, int64, string, []string) error
}

type Service struct {
	battleRepo   BattleRepo
	actionSink   ActionSink
	actionSource battlepkg.ActionSource

	spiritsv1.UnimplementedBattleServiceServer
}

var _ spiritsv1.BattleServiceServer = &Service{}

func New(battleRepo BattleRepo, actionSink ActionSink, actionSource battlepkg.ActionSource) *Service {
	return &Service{
		battleRepo:   battleRepo,
		actionSink:   actionSink,
		actionSource: actionSource,
	}
}

func (s *Service) CreateBattle(
	ctx context.Context,
	req *spiritsv1.CreateBattleRequest,
) (*spiritsv1.CreateBattleResponse, error) {
	state := spiritsv1.BattleState_BATTLE_STATE_PENDING
	apiBattle := &spiritsv1.Battle{
		State: &state,
	}
	internalBattle, err := convertbattle.FromAPI(apiBattle, s.actionSource)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	internalBattle, err = s.battleRepo.Create(ctx, internalBattle)
	if err != nil {
		return nil, err
	}

	return &spiritsv1.CreateBattleResponse{Battle: convertbattle.ToAPI(internalBattle)}, nil
}

func (s *Service) WatchBattle(
	req *spiritsv1.WatchBattleRequest,
	watch spiritsv1.BattleService_WatchBattleServer,
) error {
	id := req.GetId()
	c, err := s.battleRepo.Watch(watch.Context(), &id)
	if err != nil {
		return err
	}

	for {
		select {
		case <-watch.Context().Done():
			return nil
		case battle, ok := <-c:
			if !ok {
				break
			}
			if err := watch.Send(&spiritsv1.WatchBattleResponse{
				Battle: convertbattle.ToAPI(battle),
			}); err != nil {
				return err
			}
		}
	}
}

func (s *Service) ListBattles(
	ctx context.Context,
	req *spiritsv1.ListBattlesRequest,
) (*spiritsv1.ListBattlesResponse, error) {
	internalBattles, err := s.battleRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	var apiBattles []*spiritsv1.Battle
	for _, internalBattle := range internalBattles {
		apiBattles = append(apiBattles, convertbattle.ToAPI(internalBattle))
	}
	return &spiritsv1.ListBattlesResponse{Battles: apiBattles}, nil
}

func (s *Service) AddBattleTeam(
	ctx context.Context,
	req *spiritsv1.AddBattleTeamRequest,
) (*spiritsv1.AddBattleTeamResponse, error) {
	internalBattle, err := s.battleRepo.AddBattleTeam(ctx, req.GetBattleId(), req.GetTeamName())
	if err != nil {
		return nil, err
	}
	return &spiritsv1.AddBattleTeamResponse{Battle: convertbattle.ToAPI(internalBattle)}, nil
}

func (s *Service) AddBattleTeamSpirit(
	ctx context.Context,
	req *spiritsv1.AddBattleTeamSpiritRequest,
) (*spiritsv1.AddBattleTeamSpiritResponse, error) {
	internalBattle, err := s.battleRepo.AddBattleTeamSpirit(
		ctx, req.GetBattleId(), req.GetTeamName(), req.GetSpiritId(),
		convertbattle.SpiritIntelligenceFromAPI(req.GetIntelligence()), req.GetSeed(),
		s.actionSource)
	if err != nil {
		return nil, err
	}

	return &spiritsv1.AddBattleTeamSpiritResponse{Battle: convertbattle.ToAPI(internalBattle)}, nil
}

func (s *Service) StartBattle(
	ctx context.Context,
	req *spiritsv1.StartBattleRequest,
) (*spiritsv1.StartBattleResponse, error) {
	internalBattle, err := s.battleRepo.UpdateBattleState(
		ctx,
		req.GetId(),
		[]battlepkg.State{battlepkg.StatePending},
		battlepkg.StateStarted,
	)
	if err != nil {
		return nil, err
	}
	return &spiritsv1.StartBattleResponse{Battle: convertbattle.ToAPI(internalBattle)}, nil
}

func (s *Service) CancelBattle(
	ctx context.Context,
	req *spiritsv1.CancelBattleRequest,
) (*spiritsv1.CancelBattleResponse, error) {
	internalBattle, err := s.battleRepo.UpdateBattleState(
		ctx,
		req.GetId(),
		[]battlepkg.State{battlepkg.StateStarted, battlepkg.StateRunning, battlepkg.StateWaiting},
		battlepkg.StateCancelled,
	)
	if err != nil {
		return nil, err
	}
	return &spiritsv1.CancelBattleResponse{Battle: convertbattle.ToAPI(internalBattle)}, nil
}

func (s *Service) CallAction(
	ctx context.Context,
	req *spiritsv1.CallActionRequest,
) (*spiritsv1.CallActionResponse, error) {
	if err := s.actionSink.Post(
		ctx, req.GetBattleId(), req.GetSpiritId(), req.GetTurn(),
		req.GetActionName(), req.GetTargetSpiritIds()); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &spiritsv1.CallActionResponse{}, nil
}

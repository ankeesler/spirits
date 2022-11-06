package service

import (
	"context"

	"github.com/ankeesler/spirits0/internal/api"
)

type SpiritRepo interface {
	Get(context.Context, string) (*api.Spirit, error)
}

type BattleRepo interface {
	Create(context.Context, *api.Battle, func(*api.Battle) error) (*api.Battle, error)
	Watch(context.Context, chan<- *api.Battle) error
	List(context.Context) ([]*api.Battle, error)

	AddBattleTeam(context.Context, string, string) (*api.Battle, error)
	AddBattleTeamSpirit(context.Context, string, string, *api.BattleTeamSpirit) (*api.Battle, error)
	UpdateBattleState(context.Context, string, []api.BattleState, api.BattleState) (*api.Battle, error)
}

type ActionSource interface {
	Pend(context.Context, string, string) (string, []string, error)
}

type Service struct {
	battleRepo   BattleRepo
	spiritRepo   SpiritRepo
	actionSource ActionSource

	api.UnimplementedBattleServiceServer
}

var _ api.BattleServiceServer = &Service{}

func New(battleRepo BattleRepo, spiritRepo SpiritRepo, actionSource ActionSource) *Service {
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
	battle, err := s.battleRepo.Create(ctx, &api.Battle{}, func(*api.Battle) error { return nil })
	if err != nil {
		return nil, err
	}
	return &api.CreateBattleResponse{Battle: battle}, nil
}

func (s *Service) WatchBattle(
	req *api.WatchBattleRequest,
	watch api.BattleService_WatchBattleServer,
) error {
	c := make(chan *api.Battle) // Closed by Watch()

	if err := s.battleRepo.Watch(watch.Context(), c); err != nil {
		return err
	}

	for battle := range c {
		if battle.GetMeta().GetId() == req.GetId() {
			if err := watch.Send(&api.WatchBattleResponse{Battle: battle}); err != nil {
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
	battles, err := s.battleRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	return &api.ListBattlesResponse{Battles: battles}, nil
}

func (s *Service) AddBattleTeam(
	ctx context.Context,
	req *api.AddBattleTeamRequest,
) (*api.AddBattleTeamResponse, error) {
	battle, err := s.battleRepo.AddBattleTeam(ctx, req.GetBattleId(), req.GetTeamName())
	if err != nil {
		return nil, err
	}
	return &api.AddBattleTeamResponse{Battle: battle}, nil
}

func (s *Service) AddBattleTeamSpirit(
	ctx context.Context,
	req *api.AddBattleTeamSpiritRequest,
) (*api.AddBattleTeamSpiritResponse, error) {
	spirit, err := s.spiritRepo.Get(ctx, req.GetSpiritId())
	if err != nil {
		return nil, err
	}

	battle, err := s.battleRepo.AddBattleTeamSpirit(
		ctx, req.GetBattleId(), req.GetTeamName(), &api.BattleTeamSpirit{
			Spirit:       spirit,
			Intelligence: req.GetIntelligence(),
		})
	if err != nil {
		return nil, err
	}

	return &api.AddBattleTeamSpiritResponse{Battle: battle}, nil
}

func (s *Service) StartBattle(
	ctx context.Context,
	req *api.StartBattleRequest,
) (*api.StartBattleResponse, error) {
	battle, err := s.battleRepo.UpdateBattleState(
		ctx,
		req.GetId(),
		[]api.BattleState{api.BattleState_BATTLE_STATE_PENDING},
		api.BattleState_BATTLE_STATE_STARTED,
	)
	if err != nil {
		return nil, err
	}
	return &api.StartBattleResponse{Battle: battle}, nil
}

func (s *Service) CancelBattle(
	ctx context.Context,
	req *api.CancelBattleRequest,
) (*api.CancelBattleResponse, error) {
	battle, err := s.battleRepo.UpdateBattleState(
		ctx,
		req.GetId(),
		[]api.BattleState{api.BattleState_BATTLE_STATE_STARTED, api.BattleState_BATTLE_STATE_WAITING},
		api.BattleState_BATTLE_STATE_CANCELLED,
	)
	if err != nil {
		return nil, err
	}
	return &api.CancelBattleResponse{Battle: battle}, nil
}

func (s *Service) validateBattle(ctx context.Context, battle *api.Battle) error {
	return nil
}

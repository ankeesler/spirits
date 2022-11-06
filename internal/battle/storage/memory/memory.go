package memory

import (
	"context"
	"math/rand"
	"sync"

	"github.com/ankeesler/spirits/pkg/api"
	genericmemory "github.com/ankeesler/spirits/internal/storage/memory"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Storage struct {
	*genericmemory.Storage[*api.Battle]

	lock sync.Mutex
}

func New(r *rand.Rand) *Storage {
	return &Storage{
		Storage: genericmemory.New[*api.Battle](r),
	}
}

func (s *Storage) AddBattleTeam(
	ctx context.Context,
	battleID string,
	teamName string,
) (*api.Battle, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	battle, err := s.Get(ctx, battleID)
	if err != nil {
		return nil, err
	}

	if battle.GetState() != api.BattleState_BATTLE_STATE_PENDING {
		return nil, status.Error(codes.FailedPrecondition, "battle must be pending")
	}

	for _, team := range battle.Teams {
		if team.GetName() == teamName {
			return nil, status.Error(codes.AlreadyExists, "team already exists")
		}
	}

	battle.Teams = append(battle.GetTeams(), &api.BattleTeam{Name: teamName})

	return s.Update(ctx, battle, func(*api.Battle) error { return nil })
}

func (s *Storage) AddBattleTeamSpirit(
	ctx context.Context,
	battleID string,
	teamName string,
	spirit *api.BattleTeamSpirit,
) (*api.Battle, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	battle, err := s.Get(ctx, battleID)
	if err != nil {
		return nil, err
	}

	if battle.GetState() != api.BattleState_BATTLE_STATE_PENDING {
		return nil, status.Error(codes.FailedPrecondition, "battle must be pending")
	}

	var team *api.BattleTeam
	for i := range battle.Teams {
		if battle.Teams[i].GetName() == teamName {
			team = battle.Teams[i]
			break
		}
	}
	if team == nil {
		return nil, status.Errorf(codes.NotFound, "team not found")
	}

	team.Spirits = append(team.Spirits, spirit)

	return s.Update(ctx, battle, func(*api.Battle) error { return nil })
}

func (s *Storage) UpdateBattleState(
	ctx context.Context,
	id string,
	from []api.BattleState,
	to api.BattleState,
) (*api.Battle, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	battle, err := s.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if !slices.Contains(from, battle.GetState()) {
		return nil, status.Errorf(codes.FailedPrecondition, "battle must be one of %s", from)
	}

	battle.State = to

	return s.Update(ctx, battle, func(_ *api.Battle) error { return nil })
}

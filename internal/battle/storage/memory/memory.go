package memory

import (
	"context"
	"math/rand"
	"sync"

	battlepkg "github.com/ankeesler/spirits/internal/battle"
	spiritpkg "github.com/ankeesler/spirits/internal/spirit"
	genericmemory "github.com/ankeesler/spirits/internal/storage/memory"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Storage struct {
	*genericmemory.Storage[*battlepkg.Battle]

	lock sync.Mutex
}

func New(r *rand.Rand) *Storage {
	return &Storage{
		Storage: genericmemory.New[*battlepkg.Battle](r),
	}
}

func (s *Storage) AddBattleTeam(
	ctx context.Context,
	battleID string,
	teamName string,
) (*battlepkg.Battle, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	battle, err := s.Get(ctx, battleID)
	if err != nil {
		return nil, err
	}

	if battle.State() != battlepkg.StatePending {
		return nil, status.Error(codes.FailedPrecondition, "battle must be pending")
	}

	for _, existingTeamName := range battle.TeamNames() {
		if existingTeamName == teamName {
			return nil, status.Error(codes.AlreadyExists, "team already exists")
		}
	}

	battle.AddTeam(teamName)

	return s.Update(ctx, battle)
}

func (s *Storage) AddBattleTeamSpirit(
	ctx context.Context,
	battleID string,
	teamName string,
	spirit *spiritpkg.Spirit,
) (*battlepkg.Battle, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	battle, err := s.Get(ctx, battleID)
	if err != nil {
		return nil, err
	}

	if battle.State() != battlepkg.StatePending {
		return nil, status.Error(codes.FailedPrecondition, "battle must be pending")
	}

	teamExists := false
	for _, existingTeamName := range battle.TeamNames() {
		if existingTeamName == teamName {
			teamExists = true
			break
		}
	}
	if !teamExists {
		return nil, status.Errorf(codes.NotFound, "team not found")
	}

	battle.AddTeamSpirit(teamName, spirit)

	return s.Update(ctx, battle)
}

func (s *Storage) UpdateBattleState(
	ctx context.Context,
	id string,
	from []battlepkg.State,
	to battlepkg.State,
) (*battlepkg.Battle, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	battle, err := s.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if !slices.Contains(from, battle.State()) {
		return nil, status.Errorf(codes.FailedPrecondition, "battle must be one of %v", from)
	}

	battle.SetState(to)

	return s.Update(ctx, battle)
}

package memory

import (
	"context"
	"log"
	"math/rand"
	"sync"

	battlepkg "github.com/ankeesler/spirits/internal/battle"
	spiritpkg "github.com/ankeesler/spirits/internal/spirit"
	genericmemory "github.com/ankeesler/spirits/internal/storage/memory"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SpiritRepo interface {
	Get(context.Context, string) (*spiritpkg.Spirit, error)
}

type Storage struct {
	*genericmemory.Storage[*battlepkg.Battle]

	spiritRepo SpiritRepo

	lock sync.Mutex
}

func New(r *rand.Rand, spiritRepo SpiritRepo) *Storage {
	return &Storage{
		Storage: genericmemory.New[*battlepkg.Battle](r),

		spiritRepo: spiritRepo,
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
	spiritID string,
	intelligence battlepkg.SpiritIntelligence,
	seed int64,
	actionSource battlepkg.ActionSource,
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

	spirit, err := s.spiritRepo.Get(ctx, spiritID)
	if err != nil {
		return nil, err
	}
	battle.AddTeamSpirit(teamName, spirit, intelligence, seed, actionSource)

	return s.Update(ctx, battle)
}

func (s *Storage) UpdateBattleState(
	ctx context.Context,
	id string,
	from []battlepkg.State,
	to battlepkg.State,
) (*battlepkg.Battle, error) {
	log.Printf("waiting to update battle state from %s to %s", from, to)

	s.lock.Lock()
	defer s.lock.Unlock()

	log.Printf("updating battle state from %s to %s", from, to)

	battle, err := s.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if !slices.Contains(from, battle.State()) {
		return nil, status.Errorf(
			codes.FailedPrecondition, "battle must be one of %v but was %v", from, battle.State())
	}

	battle.SetState(to)

	return s.Update(ctx, battle)
}

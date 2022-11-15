package memory

import (
	"context"
	"math/rand"
	"sort"
	"sync"

	actionpkg "github.com/ankeesler/spirits/internal/action"
	spiritpkg "github.com/ankeesler/spirits/internal/spirit"
	genericmemory "github.com/ankeesler/spirits/internal/storage/memory"
)

type ActionRepo interface {
	Get(context.Context, string) (*actionpkg.Action, error)
}

type Storage struct {
	*genericmemory.Storage[*spiritpkg.Spirit]

	actionRepo ActionRepo

	lock sync.Mutex
}

func New(r *rand.Rand, actionRepo ActionRepo) *Storage {
	return &Storage{
		Storage:    genericmemory.New[*spiritpkg.Spirit](r),
		actionRepo: actionRepo,
	}
}

func (s *Storage) Create(
	ctx context.Context,
	spirit *spiritpkg.Spirit,
) (*spiritpkg.Spirit, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	var err error

	for _, actionName := range spirit.ActionNames() {
		action := spirit.Action(actionName)
		if id := action.ID(); len(id) > 0 {
			action, err = s.actionRepo.Get(ctx, id)
			if err != nil {
				return nil, err
			}
			spirit.SetAction(actionName, action)
		}
	}

	return s.Storage.Create(ctx, spirit)
}

func (s *Storage) ListSpirits(
	ctx context.Context,
	name *string,
) ([]*spiritpkg.Spirit, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	spirits, err := s.List(ctx)
	if err != nil {
		return nil, err
	}

	if name != nil {
		var filteredSpirits []*spiritpkg.Spirit
		for _, spirit := range spirits {
			if spirit.Name() == *name {
				filteredSpirits = append(filteredSpirits, spirit)
			}
		}
		spirits = filteredSpirits
	}

	sort.Slice(spirits, func(i, j int) bool {
		return spirits[i].Name() < spirits[j].Name()
	})

	return spirits, nil
}

func (s *Storage) Update(
	ctx context.Context,
	spirit *spiritpkg.Spirit,
) (*spiritpkg.Spirit, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	var err error

	for _, actionName := range spirit.ActionNames() {
		action := spirit.Action(actionName)
		if id := action.ID(); len(id) > 0 {
			action, err = s.actionRepo.Get(ctx, id)
			if err != nil {
				return nil, err
			}
			spirit.SetAction(actionName, action)
		}
	}

	return s.Storage.Update(ctx, spirit)
}

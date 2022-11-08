package memory

import (
	"context"
	"math/rand"
	"sync"

	spiritpkg "github.com/ankeesler/spirits/internal/spirit"
	genericmemory "github.com/ankeesler/spirits/internal/storage/memory"
)

type Storage struct {
	*genericmemory.Storage[*spiritpkg.Spirit]

	lock sync.Mutex
}

func New(r *rand.Rand) *Storage {
	return &Storage{
		Storage: genericmemory.New[*spiritpkg.Spirit](r),
	}
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

	return spirits, nil
}

package memory

import (
	"context"
	"math/rand"
	"sync"

	"github.com/ankeesler/spirits/pkg/api"
	genericmemory "github.com/ankeesler/spirits/internal/storage/memory"
)

type Storage struct {
	*genericmemory.Storage[*api.Spirit]

	lock sync.Mutex
}

func New(r *rand.Rand) *Storage {
	return &Storage{
		Storage: genericmemory.New[*api.Spirit](r),
	}
}

func (s *Storage) ListSpirits(
	ctx context.Context,
	name *string,
) ([]*api.Spirit, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	spirits, err := s.List(ctx)
	if err != nil {
		return nil, err
	}

	if name != nil {
		var filteredSpirits []*api.Spirit
		for _, spirit := range spirits {
			if spirit.Name == *name {
				filteredSpirits = append(filteredSpirits, spirit)
			}
		}
		spirits = filteredSpirits
	}

	return spirits, nil
}

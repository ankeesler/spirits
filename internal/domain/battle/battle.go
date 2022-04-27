package battle

import (
	"context"
	"errors"
	"sync"

	"github.com/ankeesler/spirits/internal/domain/spirit"
	"github.com/ankeesler/spirits/internal/domain/team"
	"github.com/ankeesler/spirits/internal/store"
)

type State string

const (
	StatePending  = "pending"
	StateStarting = "starting"
	StateRunning  = "running"
	StateFinished = "finished"
	StateError    = "error"
)

type Callback func(*Battle, error)

type Battle struct {
	Name          string
	InBattleTeams []*team.Team
	State         State
	Message       string
	Spirits       *store.Store[spirit.Spirit]

	l      sync.Mutex
	cancel context.CancelFunc
}

func New(name string, inBattleTeams []*team.Team) *Battle {
	return &Battle{
		Name:          name,
		InBattleTeams: inBattleTeams,
		State:         StatePending,
		Spirits:       store.New(func(spirit *spirit.Spirit) string { return spirit.Name }),
	}
}

func (b *Battle) Start(callback Callback) error {
	b.l.Lock()
	defer b.l.Unlock()

	if b.cancel != nil {
		return errors.New("battle already started")
	}

	var ctx context.Context
	ctx, b.cancel = context.WithCancel(context.Background())
	go b.run(ctx, callback)

	return nil
}

func (b *Battle) Stop() error {
	b.l.Lock()
	defer b.l.Unlock()

	if b.cancel == nil {
		return errors.New("battle already canceled")
	}

	b.cancel()
	b.cancel = nil

	return nil
}

func (b *Battle) run(ctx context.Context, callback Callback) {
}

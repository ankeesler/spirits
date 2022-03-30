package runner

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/ankeesler/spirits/api/internal/battle"
	"github.com/ankeesler/spirits/api/internal/spirit"
	"github.com/ankeesler/spirits/api/internal/ui"
)

// Runner provides a way of running an async battle.
type Runner struct {
	mu sync.Mutex

	s      atomic.Value // []*spirit.Spirit
	o      syncBuffer
	cancel func()
}

// Start starts a battle, if one is not already running. Start returns true iff a battle was
// started.
func (b *Runner) Start(ctx context.Context, spirits []*spirit.Spirit, onDone func()) bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.runningLocked() {
		return false
	}

	var battleCtx context.Context
	battleCtx, b.cancel = context.WithCancel(ctx)
	go func() {
		defer func() {
			b.mu.Lock()
			defer b.mu.Unlock()

			b.cancel()
			b.cancel = nil
		}()

		b.o.reset()
		u := ui.New(&b.o)
		battle.Run(battleCtx, spirits, func(spirits []*spirit.Spirit, err error) {
			u.OnSpirits(spirits, err)
			b.s.Store(spirits)
		})
		onDone()
	}()

	return true
}

// Stop stops a battle that is currently running. Stop returns true iff the battle was stopped.
func (b *Runner) Stop() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	if !b.runningLocked() {
		return false
	}

	b.cancel()

	return true
}

func (b *Runner) runningLocked() bool {
	return b.cancel != nil
}

// Output returns the battle output up to this point.
func (b *Runner) Output() string {
	return b.o.read()
}

// Spirits returns the spirits.Spirit's that are currently involved in the battle.
func (b *Runner) Spirits() []*spirit.Spirit {
	return b.s.Load().([]*spirit.Spirit)
}

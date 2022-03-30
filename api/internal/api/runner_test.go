package api

import (
	"context"
	"testing"
	"time"

	"github.com/ankeesler/spirits/api/internal/spirit"
	"github.com/stretchr/testify/require"
)

type actionFunc func(ctx context.Context, from, to *spirit.Spirit) error

func (f actionFunc) Run(ctx context.Context, from, to *spirit.Spirit) error {
	return f(ctx, from, to)
}

func TestRunner(t *testing.T) {
	b := battleRunner{}

	// Not running
	require.False(t, b.running())
	require.Panics(t, func() { b.output() })
	require.Panics(t, func() { b.spirits() })
	require.Panics(t, func() { b.stop() })

	// Is running
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	b.start(ctx, []*spirit.Spirit{
		{
			Name:   "a",
			Health: 3,
			Action: actionFunc(func(ctx context.Context, from, to *spirit.Spirit) error {
				to.Health = 2
				return nil
			}),
		},
		{
			Name:   "b",
			Health: 3,
			Action: actionFunc(func(ctx context.Context, from, to *spirit.Spirit) error {
				<-ctx.Done()
				return ctx.Err()
			}),
		},
	})
	require.Eventually(t, b.running, time.Second*2, time.Millisecond*100)
	require.Eventually(t, func() bool {
		if want, got := "> summary\n  a: 3\n  b: 3\n> summary\n  a: 3\n  b: 2\n", b.output(); want != got {
			t.Logf("want output %q, got output %q", want, got)
			return false
		}
		return true
	}, time.Second*2, time.Millisecond*100)
	spirits := b.spirits()
	require.Len(t, spirits, 2)
	require.Equal(t, spirits[0].Name, "a")
	require.Equal(t, spirits[0].Health, 3)
	require.Equal(t, spirits[1].Name, "b")
	require.Equal(t, spirits[1].Health, 2)
	require.Equal(t, "", b.output())

	// Is cancelled (externally)
	cancel()
	require.Eventually(t, func() bool { return !b.running() }, time.Second*2, time.Millisecond*100)
	require.Panics(t, func() { b.spirits() })
	require.Panics(t, func() { b.stop() })
	// TODO: need to be able to get output?

	// Run it again
	b.start(context.Background(), []*spirit.Spirit{
		{
			Name:   "a",
			Health: 3,
			Action: actionFunc(func(ctx context.Context, from, to *spirit.Spirit) error {
				to.Health = 2
				return nil
			}),
		},
		{
			Name:   "b",
			Health: 3,
			Action: actionFunc(func(ctx context.Context, from, to *spirit.Spirit) error {
				to.Health = 0
				return nil
			}),
		},
	})
	require.Eventually(t, func() bool {
		if want, got := "> summary\n  a: 3\n  b: 3\n> summary\n  a: 3\n  b:2\n> summary\n  a: 0\n  b: 2\n", b.output(); want != got {
			t.Logf("want output %q, got output %q", want, got)
			return false
		}
		return true
	}, time.Second*2, time.Millisecond*100)

	// Is cancelled (internally)
	require.Eventually(t, func() bool { return !b.running() }, time.Second*2, time.Millisecond*100)
}

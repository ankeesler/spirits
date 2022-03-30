package runner

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ankeesler/spirits/api/internal/spirit"
	"github.com/stretchr/testify/require"
)

type actionFunc func(ctx context.Context, from, to *spirit.Spirit) error

func (f actionFunc) Run(ctx context.Context, from, to *spirit.Spirit) error {
	return f(ctx, from, to)
}

type doneFunc struct {
	c int32
}

func (df *doneFunc) f() {
	atomic.AddInt32(&df.c, 1)
}

func (df *doneFunc) called() int {
	return int(atomic.LoadInt32(&df.c))
}

func TestRunner(t *testing.T) {
	b := Runner{}

	// Run asynchronously
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	df := doneFunc{}
	require.True(t, b.Start(ctx, []*spirit.Spirit{
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
	}, df.f))
	require.False(t, b.Start(ctx, nil, func() {}))
	require.Eventually(t, func() bool {
		if want, got := "> summary\n  a: 3\n  b: 3\n> summary\n  a: 3\n  b: 2\n", b.Output(); want != got {
			t.Logf("want output %q, got output %q", want, got)
			return false
		}
		return true
	}, time.Second*2, time.Millisecond*100)
	spirits := b.Spirits()
	require.Len(t, spirits, 2)
	require.Equal(t, spirits[0].Name, "a")
	require.Equal(t, spirits[0].Health, 3)
	require.Equal(t, spirits[1].Name, "b")
	require.Equal(t, spirits[1].Health, 2)
	require.Equal(t, "", b.Output())
	require.Zero(t, df.called())

	// Is cancelled (externally)
	cancel()
	require.Eventually(t, func() bool { return df.called() == 1 }, time.Second*2, time.Millisecond*100)

	// Run synchronously
	df = doneFunc{}
	require.Eventually(t, func() bool {
		return b.Start(context.Background(), []*spirit.Spirit{
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
		}, df.f)
	}, time.Second*2, time.Millisecond*100)
	require.Eventually(t, func() bool {
		if want, got := "> summary\n  a: 3\n  b: 3\n> summary\n  a: 3\n  b: 2\n> summary\n  a: 0\n  b: 2\n", b.Output(); want != got {
			t.Logf("want output %q, got output %q", want, got)
			return false
		}
		return true
	}, time.Second*2, time.Millisecond*100)
	spirits = b.Spirits()
	require.Len(t, spirits, 2)
	require.Equal(t, spirits[0].Name, "a")
	require.Equal(t, spirits[0].Health, 0)
	require.Equal(t, spirits[1].Name, "b")
	require.Equal(t, spirits[1].Health, 2)
	require.Equal(t, "", b.Output())
	require.Equal(t, 1, df.called())

	// Can run again
	require.True(t, b.Start(context.Background(), []*spirit.Spirit{
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
	}, func() {}))
}

package stats_test

import (
	"testing"

	"github.com/ankeesler/spirits/internal/tbass"
	"github.com/ankeesler/spirits/internal/tbass/actor/stats"
	"github.com/stretchr/testify/require"
)

func TestStats(t *testing.T) {
	s := stats.New()
	require.Panics(t, func() { s.Stat("health") })

	s = s.With("health", 5, 0, 5)

	health := s.Stat("health")
	require.Equal(t, tbass.StatValue(5), health.Get())
	health.Set(3)
	require.Equal(t, tbass.StatValue(3), health.Get())
	health.Set(-10)
	require.Equal(t, tbass.StatValue(0), health.Get())
	health.Set(10)
	require.Equal(t, tbass.StatValue(5), health.Get())

	s = s.With("power", 10, 1, 20).With("armour", 10, 1, 14)

	power := s.Stat("power")
	require.Equal(t, tbass.StatValue(10), power.Get())
	power.Set(20)
	require.Equal(t, tbass.StatValue(20), power.Get())
	power.Set(0)
	require.Equal(t, tbass.StatValue(1), power.Get())

	armour := s.Stat("armour")
	armour.Set(100)
	require.Equal(t, tbass.StatValue(14), armour.Get())
}

// Package spirit provides the core spirit domain concept.
//
// See Spirit doc.
package spirit

import (
	"github.com/ankeesler/spirits/internal/tbass"
	"github.com/ankeesler/spirits/internal/tbass/actor"
	"github.com/ankeesler/spirits/internal/tbass/actor/stats"
)

const (
	StatHealth  string = "health"
	StatPower          = "power"
	StatArmour         = "defense"
	StatAgility        = "agility"
)

// Spirit is a participant in a battle.Battle.
type Spirit struct {
	tbass.Actor
}

func New(name string, health, power, armour, agility tbass.StatValue, action actor.Action) *Spirit {
	stats := stats.New().
		With(StatHealth, health, 0, health).
		With(StatPower, power, 1, power*4).
		With(StatArmour, armour, 1, armour*4).
		With(StatAgility, agility, 1, agility*4)
	return &Spirit{
		Actor: actor.New(name, stats, action),
	}
}

// Flatten returns a 1-D slice of Spirit's from a 2-D array of Spirit's.
func Flatten(sss [][]*Spirit) []*Spirit {
	var ss []*Spirit
	for i := range sss {
		ss = append(ss, sss[i]...)
	}
	return ss
}

// Min finds the Spirit with the minium stat returned by statFunc.
//
// The first Spirit encountered with the minimum stat will be returned.
//
// Spirit's with current health of 0 are ignored.
func Min(ss []*Spirit, statFunc func(*Spirit) int) *Spirit {
	var s *Spirit
	for i := range ss {
		if ss[i].Stat(StatHealth).Get() > 0 && (s == nil || statFunc(ss[i]) < statFunc(s)) {
			s = ss[i]
		}
	}
	return s
}

// ForEach calls f for every Spirit in ss, in the order in which they are provided.
func ForEach(ss []*Spirit, f func(*Spirit)) {
	for _, s := range ss {
		f(s)
	}
}

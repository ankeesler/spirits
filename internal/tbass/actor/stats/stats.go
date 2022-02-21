// Package stats provides a basic actor.Stats and tbass.Stat implementation.
package stats

import (
	"fmt"

	"github.com/ankeesler/spirits/internal/tbass"
	"github.com/ankeesler/spirits/internal/tbass/actor"
)

// stat is a single tbass.Stat.
type stat struct {
	name          string
	val, min, max tbass.StatValue
}

func (s *stat) Name() string { return s.name }

func (s *stat) Get() tbass.StatValue { return s.val }

func (s *stat) Set(val tbass.StatValue) {
	if val < s.min {
		s.val = s.min
	} else if val > s.max {
		s.val = s.max
	} else {
		s.val = val
	}
}

func (s *stat) String() string {
	return fmt.Sprintf("%d <= %d <= %d", s.min, s.val, s.max)
}

// Stats stores tbass.Stat's by name.
type Stats map[string]*stat

// New creates a actor.Stats with no tbass.Stat's.
func New() Stats { return Stats{} }

// With adds a tbass.Stat to this actor.Stats.
//
// The new tbass.Stat has an initial value (val), a minimum value (min), and maximum value (max).
func (s Stats) With(name string, val, min, max tbass.StatValue) Stats {
	s[name] = &stat{name: name, val: val, min: min, max: max}
	return s
}

// Stat implements actor.Stats.Stat().
//
// If there is no know tbass.Stat with the provided name, this function will panic.
func (s Stats) Stat(name string) tbass.Stat {
	stat, ok := s[name]
	if !ok {
		panic(fmt.Sprintf("unknown stat with name %q", name))
	}
	return stat
}

func (s Stats) Clone() actor.Stats {
	sCopy := Stats{}
	for k, v := range s {
		vCopy := *v
		sCopy[k] = &vCopy
	}
	return sCopy
}

package strategy

import (
	"container/heap"

	"github.com/ankeesler/spirits/internal/tbass"
	"github.com/ankeesler/spirits/internal/tbass/activity"
)

type healthAndSpeed struct {
	healthStat, speedStat string

	tt tbass.Teams

	speedHeap speedHeap
}

// HealthAndSpeed provides a tbass.Strategy that:
//   - decides that there is a winner when all teams but one have lost all of their health, based
//     on a generic speed tbass.Stat (with name healthStat)
//   - decides which tbass.Actor acts next depending on a generic speed tbass.Stat (with name
//     speedStat)
func HealthAndSpeed(healthStat, speedStat string, tt tbass.Teams) activity.Strategy {
	return &healthAndSpeed{
		healthStat: healthStat,
		speedStat:  speedStat,
		tt:         tt,
		speedHeap:  newSpeedHeap(speedStat, tt),
	}
}

func (has *healthAndSpeed) Next() (tbass.Actor, tbass.Team) {
	return has.speedHeap.next()
}

func (has *healthAndSpeed) Winner() tbass.Teams {
	if tt := has.teamsHealthy(); len(tt) < 2 {
		return tt
	}
	return nil
}

func (has *healthAndSpeed) teamsHealthy() []tbass.Team {
	teamsHealthy := []tbass.Team{}

	for _, t := range has.tt {
		for _, a := range t.Actors() {
			if a.Stat(has.healthStat).Get() > 0 {
				teamsHealthy = append(teamsHealthy, t)
				break
			}
		}
	}

	return teamsHealthy
}

type speedHeap struct {
	speedStat string
	states    []*speedState
}

func (h speedHeap) Len() int { return len(h.states) }

func (h speedHeap) Less(i, j int) bool {
	if h.states[i].ticks < h.states[j].ticks {
		return true
	}

	// If there is a tie, the tbass.Actor with the highest Agility() goes first.
	//
	// This is so that "s(1.1.1.1)" will only go half as many times as "s(1.1.1.2)" (consider their
	// speeds are 1.0 and 0.5, respectively).
	if h.states[i].ticks == h.states[j].ticks {
		return h.states[i].a.Stat(h.speedStat).Get() > h.states[j].a.Stat(h.speedStat).Get()
	}

	return false
}

func (h speedHeap) Swap(i, j int) { h.states[i], h.states[j] = h.states[j], h.states[i] }

func (h *speedHeap) Push(x interface{}) {
	h.states = append(h.states, x.(*speedState))
}

func (h *speedHeap) Pop() interface{} {
	old := h.states
	n := len(old)
	x := old[n-1]
	h.states = old[0 : n-1]
	return x
}

func newSpeedHeap(speedStat string, tt tbass.Teams) speedHeap {
	var h speedHeap
	h.speedStat = speedStat

	// Add all tbass.Actor's to the initial (unsorted) heap.
	for _, t := range tt {
		for _, a := range t.Actors() {
			speed := a.Stat(speedStat).Get()
			h.states = append(h.states, newSpeedState(a, t, float64(1)/float64(speed)))
		}
	}

	// Heapify all tbass.Actor's.
	heap.Init(&h)

	return h
}

func (h *speedHeap) next() (tbass.Actor, tbass.Team) {
	// Get the next tbass.Actor to go.
	next := heap.Pop(h).(*speedState)

	// Decrement all ticks by the next tbass.Actor's tick value. Since we are decreasing all ticks
	// by the same scalar value, the heap invariant will be maintained.
	for _, state := range h.states {
		state.ticks -= next.ticks
	}

	// Reinitialize the next tbass.Actor so that it is enqueued again with its baseline ticks.
	speed := next.a.Stat(h.speedStat).Get()
	next = newSpeedState(next.a, next.t, float64(1)/float64(speed))

	// Add the next tbass.Actor back again to the heap.
	heap.Push(h, next)

	return next.a, next.t
}

type speedState struct {
	a tbass.Actor
	t tbass.Team

	// ticks is a countdown to this tbass.Actor's turn. When this value gets to 0 or below, it is
	// time for this tbass.Actor to run their spirit.Action.
	ticks float64
}

func newSpeedState(a tbass.Actor, t tbass.Team, ticks float64) *speedState {
	return &speedState{a: a, t: t, ticks: ticks}
}

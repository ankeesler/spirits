package battlerunner

import (
	"container/heap"

	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
)

type strategy struct {
	spirits []*spiritsinternal.Spirit

	speedHeap speedHeap
}

func newStrategy(spirits []*spiritsinternal.Spirit) *strategy {
	return &strategy{
		spirits:   spirits,
		speedHeap: newSpeedHeap(spirits),
	}
}

func (s *strategy) hasNext() bool {
	spiritsAlive := 0
	for _, spirit := range s.spirits {
		if spirit.Spec.Stats.Health > 0 {
			spiritsAlive++
		}
	}
	return spiritsAlive > 1
}

func (s *strategy) next() (*spiritsinternal.Spirit, *spiritsinternal.Spirit) {
	return s.speedHeap.next()
}

type speedHeap []*speedState

func (h speedHeap) Len() int { return len(h) }

func (h speedHeap) Less(i, j int) bool {
	if h[i].ticks < h[j].ticks {
		return true
	}

	// If there is a tie, the tbass.Actor with the highest Agility() goes first.
	//
	// This is so that "s(1.1.1.1)" will only go half as many times as "s(1.1.1.2)" (consider their
	// speeds are 1.0 and 0.5, respectively).
	if h[i].ticks == h[j].ticks {
		return h[i].s.Spec.Stats.Agility > h[j].s.Spec.Stats.Agility
	}

	return false
}

func (h speedHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *speedHeap) Push(x interface{}) {
	*h = append(*h, x.(*speedState))
}

func (h *speedHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func newSpeedHeap(spirits []*spiritsinternal.Spirit) speedHeap {
	var h speedHeap

	// Add all spirits to the initial (unsorted) heap.
	for _, spirit := range spirits {
		h = append(h, newSpeedState(spirit))
	}

	// Heapify all spirits.
	heap.Init(&h)

	return h
}

func (h *speedHeap) next() (*spiritsinternal.Spirit, *spiritsinternal.Spirit) {
	// Get the next spirit to go.
	next := heap.Pop(h).(*speedState)

	// Decrement all ticks by the next spirit's tick value. Since we are decreasing all ticks
	// by the same scalar value, the heap invariant will be maintained.
	for _, state := range *h {
		state.ticks -= next.ticks
	}

	// Reinitialize the next tbass.Actor so that it is enqueued again with its baseline ticks.
	next = newSpeedState(next.s)

	// Add the next tbass.Actor back again to the heap.
	heap.Push(h, next)

	return next.s, h.getOther(next.s)
}

func (h speedHeap) getOther(s *spiritsinternal.Spirit) *spiritsinternal.Spirit {
	// This janky logic won't always be here :)
	if h[0].s == s {
		return h[1].s
	}

	return h[0].s
}

type speedState struct {
	s *spiritsinternal.Spirit

	// ticks is a countdown to this spirit's turn. When this value gets to 0 or below, it is
	// time for this spirit to run.
	ticks float64
}

func newSpeedState(s *spiritsinternal.Spirit) *speedState {
	return &speedState{s: s, ticks: float64(1) / float64(s.Spec.Stats.Agility)}
}

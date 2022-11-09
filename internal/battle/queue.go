package battle

import (
	"container/heap"

	spiritpkg "github.com/ankeesler/spirits/internal/spirit"
)

type queue []*queueEntry

type queueEntry struct {
	ticks  float64
	spirit *spiritpkg.Spirit
}

func newQueueEntry(spirit *spiritpkg.Spirit) *queueEntry {
	return &queueEntry{ticks: float64(1) / float64(spirit.Health()), spirit: spirit}
}

func (q *queue) AddSpirit(spirit *spiritpkg.Spirit) {
	*q = append(*q, newQueueEntry(spirit))
}

func (q *queue) Init() {
	heap.Init(q)
}

func (q queue) Len() int { return len(q) }

func (q queue) Less(i, j int) bool {
	if q[i].ticks < q[j].ticks {
		return true
	}

	// If there is a tie, the highest Agility() goes first.
	//
	// This is so that "s(1.1.1.1)" will only go half as many times as "s(1.1.1.2)" (consider their
	// speeds are 1.0 and 0.5, respectively).
	if q[i].ticks == q[j].ticks {
		return q[i].spirit.Agility() > q[j].spirit.Agility()
	}

	return false
}

func (q queue) Swap(i, j int) { q[i], q[j] = q[j], q[i] }

func (q *queue) Push(x any) {
	*q = append(*q, x.(*queueEntry))
}

func (q *queue) Pop() any {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}

func (q *queue) Next() *spiritpkg.Spirit {
	// Get the next spirit to go.
	nextEntry := heap.Pop(q).(*queueEntry)

	// Decrement all ticks by the next spirit's tick value. Since we are decreasing all ticks
	// by the same scalar value, the heap invariant will be maintained.
	for _, entry := range *q {
		entry.ticks -= nextEntry.ticks
	}

	// Reinitialize the next tbass.Actor so that it is enqueued again with its baseline ticks.
	newEntry := newQueueEntry(nextEntry.spirit)

	// Add the next tbass.Actor back again to the heap.
	heap.Push(q, newEntry)

	return nextEntry.spirit
}

func (q queue) NextIDs() []string {
	var ids []string
	for _, e := range q {
		ids = append(ids, e.spirit.ID())
	}
	return ids
}

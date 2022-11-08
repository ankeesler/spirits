package battle

import spiritpkg "github.com/ankeesler/spirits/internal/spirit"

type queue []*queueEntry

type queueEntry struct {
	ticks  int64
	spirit *spiritpkg.Spirit
}

func (q *queue) AddSpirit(spirit *spiritpkg.Spirit) {
	*q = append(*q, &queueEntry{ticks: spirit.Health(), spirit: spirit})
}

func (q *queue) Next() *spiritpkg.Spirit {
	return nil
}

func (q *queue) NextIDs() []string {
	return nil
}

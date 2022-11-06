package battle

import (
	"github.com/ankeesler/spirits0/internal/spirit"
)

type Queue interface {
	HasNext() bool
	Next() (*spirit.Spirit, []*spirit.Spirit, [][]*spirit.Spirit)

	Peek() *spirit.Spirit
}

type Battle struct {
	id    string
	teams []*Team
	queue Queue
}

func New(id string, teams []*Team, queue Queue) *Battle {
	return &Battle{
		id:    id,
		teams: teams,
		queue: queue,
	}
}

func (b *Battle) ID() string     { return b.id }
func (b *Battle) Teams() []*Team { return b.teams }
func (b *Battle) Queue() Queue   { return b.queue }

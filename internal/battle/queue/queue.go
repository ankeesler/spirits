package queue

import "github.com/ankeesler/spirits0/internal/spirit"

type Queue struct {
	teams [][]*spirit.Spirit
}

func New(teams [][]*spirit.Spirit) *Queue {
	return &Queue{
		teams: teams,
	}
}

func (q *Queue) HasNext() bool {
	healthyTeams := 0
	for _, team := range q.teams {
		healthySpirits := 0
		for _, spirit := range team {
			if spirit.Stats().Health() > 0 {
				healthySpirits++
			}
		}
		if healthySpirits > 0 {
			healthyTeams++
		}
	}
	return healthyTeams > 1
}

func (q *Queue) Next() (*spirit.Spirit, []*spirit.Spirit, [][]*spirit.Spirit) {
	return nil, nil, nil
}

func (q *Queue) Peek() *spirit.Spirit {
	return nil
}

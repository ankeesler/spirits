package queue

import "github.com/ankeesler/spirits0/internal/spirit"

type Queue struct {
}

func New(spirits [][]*spirit.Spirit) *Queue {
	return &Queue{}
}

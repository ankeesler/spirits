package tbass

type Team interface {
	Name() string
	Actors() []Actor
}

type Teams []Team

func (tt Teams) Actors() []Actor {
	aa := []Actor{}
	for _, t := range tt {
		aa = append(aa, t.Actors()...)
	}
	return aa
}

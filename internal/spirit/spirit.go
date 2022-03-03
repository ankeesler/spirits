package spirit

type Action interface {
	Run(to, from *Spirit)
}

type Spirit struct {
	Name    string
	Health  int
	Power   int
	Agility int
	Armor   int
	Action  Action
}

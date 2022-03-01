package spirit

type Action interface {
	Name() string
	Run(to, from *Spirit)
}

type Spirit struct {
	Name    string
	Health  int
	Power   int
	Agility int
	Armour  int
	Action  Action
}

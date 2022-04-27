package spirit

type Spirit struct {
	Name  string
	Stats Stats
}

type Stats struct {
	Health int
}

func New(name string, health int) *Spirit {
	return &Spirit{
		Name: name,
		Stats: Stats{
			Health: health,
		},
	}
}

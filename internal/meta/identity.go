package meta

type Identity struct {
	principle string
}

func (i *Identity) Clone() *Identity {
	return &Identity{principle: i.principle}
}

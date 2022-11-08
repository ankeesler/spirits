package meta

import "github.com/ankeesler/spirits/pkg/api"

type Identity struct {
	principle string
}

func identityFromAPI(apiIdentity *api.Identity) *Identity {
	return &Identity{principle: apiIdentity.GetPrinciple()}
}

func (i *Identity) Clone() *Identity {
	return &Identity{principle: i.principle}
}

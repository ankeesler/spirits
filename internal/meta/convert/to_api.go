package convert

import (
	metapkg "github.com/ankeesler/spirits/internal/meta"
	"github.com/ankeesler/spirits/pkg/api"
)

func ToAPI(internalMeta *metapkg.Meta) *api.Meta {
	return &api.Meta{
		Id: internalMeta.ID(),
	}
}

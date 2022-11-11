package convert

import (
	metapkg "github.com/ankeesler/spirits/internal/meta"
	"github.com/ankeesler/spirits/pkg/api"
)

func FromAPI(apiMeta *api.Meta) *metapkg.Meta {
	internalMeta := metapkg.New()
	internalMeta.SetID(apiMeta.GetId())
	return internalMeta
}

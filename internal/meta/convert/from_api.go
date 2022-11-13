package convert

import (
	metapkg "github.com/ankeesler/spirits/internal/meta"
	"github.com/ankeesler/spirits/pkg/api/spirits/v1"
)

func FromAPI(apiMeta *spiritsv1.Meta) *metapkg.Meta {
	internalMeta := metapkg.New()
	internalMeta.SetID(apiMeta.GetId())
	return internalMeta
}

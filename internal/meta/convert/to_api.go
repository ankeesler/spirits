package convert

import (
	metapkg "github.com/ankeesler/spirits/internal/meta"
	"github.com/ankeesler/spirits/pkg/api/spirits/v1"
)

func ToAPI(internalMeta *metapkg.Meta) *spiritsv1.Meta {
	return &spiritsv1.Meta{
		Id: internalMeta.ID(),
	}
}

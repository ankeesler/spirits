package convert

import (
	metapkg "github.com/ankeesler/spirits/internal/meta"
	spiritsv1 "github.com/ankeesler/spirits/pkg/api/spirits/v1"
)

func ToAPI(internalMeta *metapkg.Meta) *spiritsv1.Meta {
	return &spiritsv1.Meta{
		Id: stringPtr(internalMeta.ID()),
	}
}

func stringPtr(s string) *string { return &s }

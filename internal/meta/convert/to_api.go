package convert

import (
	metapkg "github.com/ankeesler/spirits/internal/meta"
	spiritsv1 "github.com/ankeesler/spirits/pkg/api/spirits/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToAPI(internalMeta *metapkg.Meta) *spiritsv1.Meta {
	return &spiritsv1.Meta{
		Id:          stringPtr(internalMeta.ID()),
		CreatedTime: timestamppb.New(internalMeta.CreatedTime()),
		UpdatedTime: timestamppb.New(internalMeta.UpdatedTime()),
	}
}

func stringPtr(s string) *string { return &s }

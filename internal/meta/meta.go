package meta

import (
	"time"

	"github.com/ankeesler/spirits/pkg/api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Meta struct {
	id          string
	createdTime time.Time
	createdBy   *Identity
	updatedTime time.Time
	updatedBy   *Identity
}

func (m *Meta) ID() string      { return m.id }
func (m *Meta) SetID(id string) { m.id = id }

func (m *Meta) CreatedTime() time.Time               { return m.createdTime }
func (m *Meta) SetCreatedTime(createdTime time.Time) { m.createdTime = createdTime }

func (m *Meta) UpdatedTime() time.Time               { return m.updatedTime }
func (m *Meta) SetUpdatedTime(updatedTime time.Time) { m.updatedTime = updatedTime }

func (m *Meta) Clone() *Meta {
	return &Meta{
		id:          m.id,
		createdTime: m.createdTime,
		createdBy:   m.createdBy.Clone(),
		updatedTime: m.updatedTime,
		updatedBy:   m.updatedBy.Clone(),
	}
}

func (m *Meta) ToAPI() *api.Meta {
	return &api.Meta{
		Id:          m.id,
		CreatedTime: timestamppb.New(m.CreatedTime()),
		UpdatedTime: timestamppb.New(m.CreatedTime()),
	}
}

func FromAPI(apiMeta *api.Meta) *Meta {
	return &Meta{
		id:          apiMeta.GetId(),
		createdTime: apiMeta.GetCreatedTime().AsTime(),
		createdBy:   identityFromAPI(apiMeta.GetCreatedBy()),
		updatedTime: apiMeta.GetUpdatedTime().AsTime(),
		updatedBy:   identityFromAPI(apiMeta.GetUpdatedBy()),
	}
}

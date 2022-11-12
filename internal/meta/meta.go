package meta

import (
	"time"
)

type Meta struct {
	id          string
	createdTime time.Time
	createdBy   *Identity
	updatedTime time.Time
	updatedBy   *Identity
}

func New() *Meta {
	return &Meta{}
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
		updatedTime: m.updatedTime,
	}
}

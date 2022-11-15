package spirit

import spiritsv1 "github.com/ankeesler/spirits/pkg/api/spirits/v1"

type stats struct {
	health               int64
	physicalPower        int64
	physicalConstitution int64
	mentalPower          int64
	mentalConstitution   int64
	agility              int64
}

func (s *stats) Health() int64          { return s.health }
func (s *stats) SetHealth(health int64) { s.health = health }

func (s *stats) PhysicalPower() int64                 { return s.physicalPower }
func (s *stats) SetPhysicalPower(physicalPower int64) { s.physicalPower = physicalPower }

func (s *stats) PhysicalConstitution() int64 { return s.physicalConstitution }
func (s *stats) SetPhysicalConstitution(physicalConstitution int64) {
	s.physicalConstitution = physicalConstitution
}

func (s *stats) MentalPower() int64               { return s.mentalPower }
func (s *stats) SetMentalPower(mentalPower int64) { s.mentalPower = mentalPower }

func (s *stats) MentalConstitution() int64 { return s.mentalConstitution }
func (s *stats) SetMentalConstitution(mentalConstitution int64) {
	s.mentalConstitution = mentalConstitution
}

func (s *stats) Agility() int64           { return s.agility }
func (s *stats) SetAgility(agility int64) { s.agility = agility }

func (s *stats) ToAPI() *spiritsv1.SpiritStats {
	return &spiritsv1.SpiritStats{
		Health: int64Ptr(s.Health()),

		PhysicalPower:        int64Ptr(s.PhysicalPower()),
		PhysicalConstitution: int64Ptr(s.PhysicalConstitution()),

		MentalPower:        int64Ptr(s.MentalPower()),
		MentalConstitution: int64Ptr(s.MentalConstitution()),

		Agility: int64Ptr(s.Agility()),
	}
}

func (s *stats) Clone() *stats {
	return &stats{
		health:               s.health,
		physicalPower:        s.physicalPower,
		physicalConstitution: s.physicalConstitution,
		mentalPower:          s.mentalPower,
		mentalConstitution:   s.mentalConstitution,
		agility:              s.agility,
	}
}

func int64Ptr(i int64) *int64 { return &i }

package spirit

type Stats struct {
	health               int64
	physicalPower        int64
	physicalConstitution int64
	mentalPower          int64
	mentalConstitution   int64
	agility              int64
}

func NewStats(
	health int64,
	physicalPower int64,
	physicalConstitution int64,
	mentalPower int64,
	mentalConstitution int64,
	agility int64,
) *Stats {
	return &Stats{
		health:               health,
		physicalPower:        physicalPower,
		physicalConstitution: physicalConstitution,
		mentalPower:          mentalPower,
		mentalConstitution:   mentalConstitution,
		agility:              agility,
	}
}

func (s *Stats) Health() int64          { return s.health }
func (s *Stats) SetHealth(health int64) { s.health = health }

func (s *Stats) PhysicalPower() int64                 { return s.physicalPower }
func (s *Stats) SetPhysicalPower(physicalPower int64) { s.physicalPower = physicalPower }

func (s *Stats) PhysicalConstitution() int64 { return s.physicalConstitution }
func (s *Stats) SetPhysicalConstitution(physicalConstitution int64) {
	s.physicalConstitution = physicalConstitution
}

func (s *Stats) MentalPower() int64               { return s.mentalPower }
func (s *Stats) SetMentalPower(mentalPower int64) { s.mentalPower = mentalPower }

func (s *Stats) MentalConstitution() int64 { return s.mentalConstitution }
func (s *Stats) SetMentalConstitution(mentalConstitution int64) {
	s.mentalConstitution = mentalConstitution
}

func (s *Stats) Agility() int64           { return s.agility }
func (s *Stats) SetAgility(agility int64) { s.agility = agility }

func (s *Stats) Clone() *Stats {
	return &Stats{
		health:               s.health,
		physicalPower:        s.physicalPower,
		physicalConstitution: s.physicalConstitution,
		mentalPower:          s.mentalPower,
		mentalConstitution:   s.mentalConstitution,
		agility:              s.agility,
	}
}

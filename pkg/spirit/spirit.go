// Package spirit provides the core spirit domain concept.
//
// See Spirit doc.
package spirit

// Spirit is a participant in a battle.Battle.
//
// A Spirit has health. When a Spirit's health is decreased to 0, it cannot participate in a
// battle.Battle anymore.
//
// A Spirit has power, which determines the strength of its attacks.
//
// A Spirit has armour, which determines the strength of its defense.
//
// A Spirit has agility, which determines how often it can perform an Action.
//
// A Spirit has an Action, which is actually how it contributes to the battle.Battle.
type Spirit struct {
	name string

	action Action

	baseHealth, health   int
	basePower, power     int
	baseArmour, armour   int
	baseAgility, agility int
}

func New(name string, health, power, armour, agility int, action Action) *Spirit {
	s := &Spirit{
		name:   name,
		action: action,
	}

	s.baseHealth = s.health
	s.health = health

	s.basePower = s.power
	s.power = power

	s.baseArmour = s.armour
	s.armour = armour

	s.baseAgility = s.agility
	s.agility = agility

	return s
}

func (s *Spirit) Name() string { return s.name }

func (s *Spirit) Action() Action { return s.action }

// Health returns this Spirit's current health.
func (s *Spirit) Health() int              { return s.health }
func (s *Spirit) IncreaseHealth(count int) { changeStat(true, &s.health, count, 0, s.baseHealth) }
func (s *Spirit) DecreaseHealth(count int) { changeStat(false, &s.health, count, 0, s.baseHealth) }

// Power returns this Spirit's current power stat.
func (s *Spirit) Power() int              { return s.power }
func (s *Spirit) IncreasePower(count int) { changeStat(true, &s.power, count, 1, s.basePower) }
func (s *Spirit) DecreasePower(count int) { changeStat(false, &s.power, count, 1, s.basePower) }

// Armour returns this Spirit's current armour stat.
func (s *Spirit) Armour() int              { return s.armour }
func (s *Spirit) IncreaseArmour(count int) { changeStat(true, &s.armour, count, 0, s.baseArmour) }
func (s *Spirit) DecreaseArmour(count int) { changeStat(false, &s.armour, count, 0, s.baseArmour) }

// Agility returns this Spirit's current agility stat.
func (s *Spirit) Agility() int              { return s.agility }
func (s *Spirit) IncreaseAgility(count int) { changeStat(true, &s.agility, count, 1, s.baseAgility) }
func (s *Spirit) DecreaseAgility(count int) { changeStat(false, &s.agility, count, 1, s.baseAgility) }

func changeStat(increase bool, stat *int, count int, min, max int) {
	if count < 0 {
		return
	}

	if !increase {
		count *= -1
	}

	*stat += count

	if *stat < min {
		*stat = min
	}

	if *stat > max {
		*stat = max
	}
}

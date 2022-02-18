package action

import (
	"strings"

	"github.com/ankeesler/spirits/internal/spirit"
)

// Target is a bitmask that describes on what/who a spirit.Action is performed.
type Target uint32

const (
	// TargetAll refers to all spirit.Spirit's in a battle.Battle.
	TargetAll Target = 0xFF
	// TargetThem refers to all spirit.Spirit's not on one's team.Team in a battle.Battle.
	TargetThem = 0x07
	// TargetRandomThem refers to a random selection of 1 or more spirit.Spirit's not on one's
	// team.Team in a battle.Battle.
	TargetRandomThem = 0x03
	// TargetOneOfThem refers to a random selection of 1 spirit.Spirit's on an opposing team.Team in a
	// battle.Battle.
	TargetOneOfThem = 0x01
	// TargetUs refers to the spirit.Spirit's on one's team.Team in a battle.Battle.
	TargetUs = 0x18
	// TargetMe refers to the spirit.Spirit performing the spirit.Action.
	TargetMe = 0x08

	// TargetPower refers to the spirit.Spirit's power stat.
	TargetPower = (1 << (8 + iota))
	// TargetPower refers to the spirit.Spirit's armour stat.
	TargetArmour
	// TargetPower refers to the spirit.Spirit's agility stat.
	TargetAgility
)

func (t Target) String() string {
	b := strings.Builder{}
	switch t & TargetAll {
	case TargetAll:
		b.WriteString("all")
	case TargetThem:
		b.WriteString("them")
	case TargetRandomThem:
		b.WriteString("random-them")
	case TargetOneOfThem:
		b.WriteString("one-of-them")
	case TargetUs:
		b.WriteString("us")
	case TargetMe:
		b.WriteString("me")
	default:
		b.WriteString("???")
	}

	if t&TargetPower != 0 {
		b.WriteString("+power")
	}
	if t&TargetArmour != 0 {
		b.WriteString("+armour")
	}
	if t&TargetAgility != 0 {
		b.WriteString("+agility")
	}

	return b.String()
}

func (t Target) find(ac *spirit.ActionContext) []*spirit.Spirit {
	return []*spirit.Spirit{}
}

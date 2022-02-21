// Package manifest provides a mapping from an API manifest into the runtime data model.
package manifest

import (
	"fmt"

	"github.com/ankeesler/spirits/internal/spirit"
	api "github.com/ankeesler/spirits/pkg/v0"
)

// Load instantiates a list of team.Team's from an api.Manifest.
func Load(m *api.Manifest) ([]*team.Team, error) {
	teams := make([]*team.Team, len(m.Data.Teams))

	for i, t := range m.Data.Teams {
		var err error
		teams[i], err = loadTeam(t)
		if err != nil {
			return nil, fmt.Errorf("cannot instantiate team %q: %w", t.Name, err)
		}
	}

	return teams, nil
}

func loadTeam(t *api.Team) (*team.Team, error) {
	spirits := make([]*spirit.Spirit, len(t.Spirits))

	for i, s := range t.Spirits {
		var err error
		spirits[i], err = loadSpirit(s)
		if err != nil {
			return nil, fmt.Errorf("cannot instantiate spirit %q: %w", s.Name, err)
		}
	}

	return team.New(t.Name, spirits...), nil
}

func loadSpirit(s *api.Spirit) (*spirit.Spirit, error) {
	a, err := loadAction(s.Action)
	if err != nil {
		return nil, fmt.Errorf("cannot instantiate action: %w", err)
	}

	return spirit.New(s.Name, s.Health, s.Power, s.Armour, s.Agility, a), nil
}

func loadAction(a *api.Action) (spirit.Action, error) {
	switch a.Type {
	case "attack":
		return action.Attack(action.Target(a.Target)), nil
	case "heal":
		return action.Heal(action.Target(a.Target)), nil
	case "buf":
		return action.Buf(action.Target(a.Target)), nil
	case "debuf":
		return action.Debuf(action.Target(a.Target)), nil
	default:
		return nil, fmt.Errorf("unrecognized action type: %q", a.Type)
	}
}

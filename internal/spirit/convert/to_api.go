package convert

import (
	convertaction "github.com/ankeesler/spirits/internal/action/convert"
	convertmeta "github.com/ankeesler/spirits/internal/meta/convert"
	spiritpkg "github.com/ankeesler/spirits/internal/spirit"
	"github.com/ankeesler/spirits/pkg/api"
)

func ToAPI(internalSpirit *spiritpkg.Spirit) *api.Spirit {
	apiSpirit := &api.Spirit{
		Meta: convertmeta.ToAPI(internalSpirit.Meta),
		Name: internalSpirit.Name(),
		Stats: &api.SpiritStats{
			Health:               internalSpirit.Health(),
			PhysicalPower:        internalSpirit.PhysicalPower(),
			PhysicalConstitution: internalSpirit.PhysicalConstitution(),
			MentalPower:          internalSpirit.MentalPower(),
			MentalConstitution:   internalSpirit.MentalConstitution(),
			Agility:              internalSpirit.Agility(),
		},
	}

	for _, actionName := range internalSpirit.ActionNames() {
		internalAction := internalSpirit.Action(actionName)
		apiSpiritAction := &api.SpiritAction{
			Name: actionName,
		}

		if id := internalAction.ID(); len(id) > 0 {
			apiSpiritAction.Definition = &api.SpiritAction_ActionId{ActionId: id}
		} else {
			apiSpiritAction.Definition = &api.SpiritAction_Inline{
				Inline: convertaction.ToAPI(internalAction),
			}
		}

		apiSpirit.Actions = append(apiSpirit.Actions, apiSpiritAction)
	}

	return apiSpirit
}

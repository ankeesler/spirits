package convert

import (
	convertaction "github.com/ankeesler/spirits/internal/action/convert"
	convertmeta "github.com/ankeesler/spirits/internal/meta/convert"
	spiritpkg "github.com/ankeesler/spirits/internal/spirit"
	"github.com/ankeesler/spirits/pkg/api/spirits/v1"
)

func ToAPI(internalSpirit *spiritpkg.Spirit) *spiritsv1.Spirit {
	apiSpirit := &spiritsv1.Spirit{
		Meta: convertmeta.ToAPI(internalSpirit.Meta),
		Name: internalSpirit.Name(),
		Stats: &spiritsv1.SpiritStats{
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
		apiSpiritAction := &spiritsv1.SpiritAction{
			Name: actionName,
		}

		if id := internalAction.ID(); len(id) > 0 {
			apiSpiritAction.Definition = &spiritsv1.SpiritAction_ActionId{ActionId: id}
		} else {
			apiSpiritAction.Definition = &spiritsv1.SpiritAction_Inline{
				Inline: convertaction.ToAPI(internalAction),
			}
		}

		apiSpirit.Actions = append(apiSpirit.Actions, apiSpiritAction)
	}

	return apiSpirit
}

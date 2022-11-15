package convert

import (
	convertaction "github.com/ankeesler/spirits/internal/action/convert"
	convertmeta "github.com/ankeesler/spirits/internal/meta/convert"
	spiritpkg "github.com/ankeesler/spirits/internal/spirit"
	spiritsv1 "github.com/ankeesler/spirits/pkg/api/spirits/v1"
)

func ToAPI(internalSpirit *spiritpkg.Spirit) *spiritsv1.Spirit {
	apiSpirit := &spiritsv1.Spirit{
		Meta: convertmeta.ToAPI(internalSpirit.Meta),
		Name: stringPtr(internalSpirit.Name()),
		Stats: &spiritsv1.SpiritStats{
			Health:               int64Ptr(internalSpirit.Health()),
			PhysicalPower:        int64Ptr(internalSpirit.PhysicalPower()),
			PhysicalConstitution: int64Ptr(internalSpirit.PhysicalConstitution()),
			MentalPower:          int64Ptr(internalSpirit.MentalPower()),
			MentalConstitution:   int64Ptr(internalSpirit.MentalConstitution()),
			Agility:              int64Ptr(internalSpirit.Agility()),
		},
	}

	for _, actionName := range internalSpirit.ActionNames() {
		internalAction := internalSpirit.Action(actionName)
		apiSpiritAction := &spiritsv1.SpiritAction{
			Name: &actionName,
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

func stringPtr(s string) *string { return &s }
func int64Ptr(i int64) *int64    { return &i }

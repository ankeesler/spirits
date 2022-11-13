package convert

import (
	"fmt"

	actionpkg "github.com/ankeesler/spirits/internal/action"
	convertaction "github.com/ankeesler/spirits/internal/action/convert"
	metapkg "github.com/ankeesler/spirits/internal/meta"
	convertmeta "github.com/ankeesler/spirits/internal/meta/convert"
	spiritpkg "github.com/ankeesler/spirits/internal/spirit"
	"github.com/ankeesler/spirits/pkg/api/spirits/v1"
)

func FromAPI(apiSpirit *spiritsv1.Spirit) (*spiritpkg.Spirit, error) {
	internalMeta := convertmeta.FromAPI(apiSpirit.GetMeta())
	internalSpirit := spiritpkg.New(internalMeta)

	internalSpirit.SetName(apiSpirit.GetName())

	internalSpirit.SetHealth(apiSpirit.Stats.GetHealth())
	internalSpirit.SetPhysicalPower(apiSpirit.Stats.GetPhysicalPower())
	internalSpirit.SetPhysicalConstitution(apiSpirit.Stats.GetPhysicalConstitution())
	internalSpirit.SetMentalPower(apiSpirit.Stats.GetMentalPower())
	internalSpirit.SetMentalConstitution(apiSpirit.Stats.GetMentalConstitution())
	internalSpirit.SetAgility(apiSpirit.Stats.GetAgility())

	for _, apiSpiritAction := range apiSpirit.GetActions() {
		actionName := apiSpiritAction.GetName()
		if internalSpirit.Action(actionName) != nil {
			return nil, fmt.Errorf("duplicate action name: %s", actionName)
		}

		var internalAction *actionpkg.Action
		switch definition := apiSpiritAction.GetDefinition().(type) {
		case *spiritsv1.SpiritAction_ActionId:
			internalAction = actionpkg.New(metapkg.New())
			internalAction.SetID(definition.ActionId)

		case *spiritsv1.SpiritAction_Inline:
			var err error
			internalAction, err = convertaction.FromAPI(definition.Inline)
			if err != nil {
				return nil, fmt.Errorf("invalid action inline for %s: %w",
					apiSpiritAction.GetName(), err)
			}
		}

		internalSpirit.SetAction(actionName, internalAction)
	}

	return internalSpirit, nil
}

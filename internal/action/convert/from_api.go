package convert

import (
	"errors"
	"fmt"

	actionpkg "github.com/ankeesler/spirits/internal/action"
	convertmeta "github.com/ankeesler/spirits/internal/meta/convert"
	"github.com/ankeesler/spirits/pkg/api"
)

func FromAPI(apiAction *api.Action) (*actionpkg.Action, error) {
	internalMeta := convertmeta.FromAPI(apiAction.GetMeta())
	internalAction := actionpkg.New(internalMeta)

	internalAction.SetDescription(apiAction.GetDescription())

	switch definition := apiAction.Definition.(type) {
	case *api.Action_Script:
		if err := internalAction.SetScript(definition.Script); err != nil {
			return nil, fmt.Errorf("set action script: %w", err)
		}
	default:
		return nil, errors.New("definition not set")
	}

	return internalAction, nil
}

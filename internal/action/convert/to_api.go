package convert

import (
	actionpkg "github.com/ankeesler/spirits/internal/action"
	convertmeta "github.com/ankeesler/spirits/internal/meta/convert"
	"github.com/ankeesler/spirits/pkg/api"
)

func ToAPI(internalAction *actionpkg.Action) *api.Action {
	apiAction := &api.Action{
		Meta:        convertmeta.ToAPI(internalAction.Meta),
		Description: internalAction.Description(),
	}
	if script := internalAction.Script(); script != nil {
		apiAction.Definition = &api.Action_Script{Script: *script}
	}
	return apiAction
}

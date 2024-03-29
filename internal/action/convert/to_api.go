package convert

import (
	actionpkg "github.com/ankeesler/spirits/internal/action"
	convertmeta "github.com/ankeesler/spirits/internal/meta/convert"
	spiritsv1 "github.com/ankeesler/spirits/pkg/api/spirits/v1"
)

func ToAPI(internalAction *actionpkg.Action) *spiritsv1.Action {
	apiAction := &spiritsv1.Action{
		Meta:        convertmeta.ToAPI(internalAction.Meta),
		Description: stringPtr(internalAction.Description()),
	}
	if script := internalAction.Script(); script != nil {
		apiAction.Definition = &spiritsv1.Action_Script{Script: *script}
	}
	return apiAction
}

func stringPtr(s string) *string { return &s }

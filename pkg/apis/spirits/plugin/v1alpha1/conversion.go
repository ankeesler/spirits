package v1alpha1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"

	plugin "github.com/ankeesler/spirits/internal/apis/spirits/plugin"
	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
)

func Convert_v1alpha1_ActionRunSpec_To_plugin_ActionRunSpec(in *ActionRunSpec, out *plugin.ActionRunSpec, s conversion.Scope) error {
	if err := spiritsv1alpha1.Convert_v1alpha1_SpiritSpec_To_spirits_SpiritSpec(&in.From, &out.From, s); err != nil {
		return err
	}
	if err := spiritsv1alpha1.Convert_v1alpha1_SpiritSpec_To_spirits_SpiritSpec(&in.To, &out.To, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_ActionRunStatus_To_plugin_ActionRunStatus(in *ActionRunStatus, out *plugin.ActionRunStatus, s conversion.Scope) error {
	if err := spiritsv1alpha1.Convert_v1alpha1_SpiritSpec_To_spirits_SpiritSpec(&in.From, &out.From, s); err != nil {
		return err
	}
	if err := spiritsv1alpha1.Convert_v1alpha1_SpiritSpec_To_spirits_SpiritSpec(&in.To, &out.To, s); err != nil {
		return err
	}
	return nil
}

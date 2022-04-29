package v1alpha1

import (
	spirits "github.com/ankeesler/spirits/internal/apis/spirits"
	conversion "k8s.io/apimachinery/pkg/conversion"
)

func Convert_spirits_SpiritSpec_To_v1alpha1_SpiritSpec(in *spirits.SpiritSpec, out *SpiritSpec, s conversion.Scope) error {
	return autoConvert_spirits_SpiritSpec_To_v1alpha1_SpiritSpec(in, out, s)
}

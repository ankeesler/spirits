package v1alpha1

import (
	"fmt"
	"sort"

	spirits "github.com/ankeesler/spirits/internal/apis/spirits"
	conversion "k8s.io/apimachinery/pkg/conversion"
)

func Convert_spirits_SpiritSpec_To_v1alpha1_SpiritSpec(in *spirits.SpiritSpec, out *SpiritSpec, s conversion.Scope) error {
	return autoConvert_spirits_SpiritSpec_To_v1alpha1_SpiritSpec(in, out, s)
}

func Convert_Slice_v1alpha1_NamedSpiritAction_To_Map_string_To_Pointer_spirits_SpiritAction(in *[]NamedSpiritAction, out *map[string]*spirits.SpiritAction, s conversion.Scope) error {
	for _, namedSpiritAction := range *in {
		if _, ok := (*out)[namedSpiritAction.Name]; ok {
			return fmt.Errorf("error converting *[]NamedContext into *map[string]*api.Context: duplicate name \"%v\" in list: %v", namedSpiritAction.Name, *in)
		}
		var internalSpiritAction spirits.SpiritAction
		if err := autoConvert_v1alpha1_SpiritAction_To_spirits_SpiritAction(&namedSpiritAction.Action, &internalSpiritAction, s); err != nil {
			return err
		}
		if (*out) == nil {
			*out = map[string]*spirits.SpiritAction{}
		}
		(*out)[namedSpiritAction.Name] = &internalSpiritAction
	}
	return nil
}

func Convert_Map_string_To_Pointer_spirits_SpiritAction_To_Slice_v1alpha1_NamedSpiritAction(in *map[string]*spirits.SpiritAction, out *[]NamedSpiritAction, s conversion.Scope) error {
	allKeys := make([]string, 0, len(*in))
	for key := range *in {
		allKeys = append(allKeys, key)
	}
	sort.Strings(allKeys)

	for _, key := range allKeys {
		internalSpiritAction := (*in)[key]
		var externalSpiritAction SpiritAction
		if err := autoConvert_spirits_SpiritAction_To_v1alpha1_SpiritAction(internalSpiritAction, &externalSpiritAction, s); err != nil {
			return err
		}
		namedSpiritAction := NamedSpiritAction{Name: key, Action: externalSpiritAction}
		*out = append(*out, namedSpiritAction)
	}
	return nil
}

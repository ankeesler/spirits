//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	spirits "github.com/ankeesler/spirits/internal/apis/spirits"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Battle)(nil), (*spirits.Battle)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Battle_To_spirits_Battle(a.(*Battle), b.(*spirits.Battle), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*spirits.Battle)(nil), (*Battle)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_spirits_Battle_To_v1alpha1_Battle(a.(*spirits.Battle), b.(*Battle), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*BattleList)(nil), (*spirits.BattleList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_BattleList_To_spirits_BattleList(a.(*BattleList), b.(*spirits.BattleList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*spirits.BattleList)(nil), (*BattleList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_spirits_BattleList_To_v1alpha1_BattleList(a.(*spirits.BattleList), b.(*BattleList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*BattleSpec)(nil), (*spirits.BattleSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_BattleSpec_To_spirits_BattleSpec(a.(*BattleSpec), b.(*spirits.BattleSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*spirits.BattleSpec)(nil), (*BattleSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_spirits_BattleSpec_To_v1alpha1_BattleSpec(a.(*spirits.BattleSpec), b.(*BattleSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*BattleStatus)(nil), (*spirits.BattleStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_BattleStatus_To_spirits_BattleStatus(a.(*BattleStatus), b.(*spirits.BattleStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*spirits.BattleStatus)(nil), (*BattleStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_spirits_BattleStatus_To_v1alpha1_BattleStatus(a.(*spirits.BattleStatus), b.(*BattleStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Spirit)(nil), (*spirits.Spirit)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Spirit_To_spirits_Spirit(a.(*Spirit), b.(*spirits.Spirit), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*spirits.Spirit)(nil), (*Spirit)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_spirits_Spirit_To_v1alpha1_Spirit(a.(*spirits.Spirit), b.(*Spirit), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SpiritList)(nil), (*spirits.SpiritList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_SpiritList_To_spirits_SpiritList(a.(*SpiritList), b.(*spirits.SpiritList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*spirits.SpiritList)(nil), (*SpiritList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_spirits_SpiritList_To_v1alpha1_SpiritList(a.(*spirits.SpiritList), b.(*SpiritList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SpiritSpec)(nil), (*spirits.SpiritSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_SpiritSpec_To_spirits_SpiritSpec(a.(*SpiritSpec), b.(*spirits.SpiritSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SpiritStats)(nil), (*spirits.SpiritStats)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_SpiritStats_To_spirits_SpiritStats(a.(*SpiritStats), b.(*spirits.SpiritStats), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*spirits.SpiritStats)(nil), (*SpiritStats)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_spirits_SpiritStats_To_v1alpha1_SpiritStats(a.(*spirits.SpiritStats), b.(*SpiritStats), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SpiritStatus)(nil), (*spirits.SpiritStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_SpiritStatus_To_spirits_SpiritStatus(a.(*SpiritStatus), b.(*spirits.SpiritStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*spirits.SpiritStatus)(nil), (*SpiritStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_spirits_SpiritStatus_To_v1alpha1_SpiritStatus(a.(*spirits.SpiritStatus), b.(*SpiritStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*spirits.SpiritSpec)(nil), (*SpiritSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_spirits_SpiritSpec_To_v1alpha1_SpiritSpec(a.(*spirits.SpiritSpec), b.(*SpiritSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_Battle_To_spirits_Battle(in *Battle, out *spirits.Battle, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_BattleSpec_To_spirits_BattleSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_BattleStatus_To_spirits_BattleStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Battle_To_spirits_Battle is an autogenerated conversion function.
func Convert_v1alpha1_Battle_To_spirits_Battle(in *Battle, out *spirits.Battle, s conversion.Scope) error {
	return autoConvert_v1alpha1_Battle_To_spirits_Battle(in, out, s)
}

func autoConvert_spirits_Battle_To_v1alpha1_Battle(in *spirits.Battle, out *Battle, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_spirits_BattleSpec_To_v1alpha1_BattleSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_spirits_BattleStatus_To_v1alpha1_BattleStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_spirits_Battle_To_v1alpha1_Battle is an autogenerated conversion function.
func Convert_spirits_Battle_To_v1alpha1_Battle(in *spirits.Battle, out *Battle, s conversion.Scope) error {
	return autoConvert_spirits_Battle_To_v1alpha1_Battle(in, out, s)
}

func autoConvert_v1alpha1_BattleList_To_spirits_BattleList(in *BattleList, out *spirits.BattleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]spirits.Battle)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_BattleList_To_spirits_BattleList is an autogenerated conversion function.
func Convert_v1alpha1_BattleList_To_spirits_BattleList(in *BattleList, out *spirits.BattleList, s conversion.Scope) error {
	return autoConvert_v1alpha1_BattleList_To_spirits_BattleList(in, out, s)
}

func autoConvert_spirits_BattleList_To_v1alpha1_BattleList(in *spirits.BattleList, out *BattleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Battle)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_spirits_BattleList_To_v1alpha1_BattleList is an autogenerated conversion function.
func Convert_spirits_BattleList_To_v1alpha1_BattleList(in *spirits.BattleList, out *BattleList, s conversion.Scope) error {
	return autoConvert_spirits_BattleList_To_v1alpha1_BattleList(in, out, s)
}

func autoConvert_v1alpha1_BattleSpec_To_spirits_BattleSpec(in *BattleSpec, out *spirits.BattleSpec, s conversion.Scope) error {
	out.Spirits = *(*[]v1.LocalObjectReference)(unsafe.Pointer(&in.Spirits))
	return nil
}

// Convert_v1alpha1_BattleSpec_To_spirits_BattleSpec is an autogenerated conversion function.
func Convert_v1alpha1_BattleSpec_To_spirits_BattleSpec(in *BattleSpec, out *spirits.BattleSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_BattleSpec_To_spirits_BattleSpec(in, out, s)
}

func autoConvert_spirits_BattleSpec_To_v1alpha1_BattleSpec(in *spirits.BattleSpec, out *BattleSpec, s conversion.Scope) error {
	out.Spirits = *(*[]v1.LocalObjectReference)(unsafe.Pointer(&in.Spirits))
	return nil
}

// Convert_spirits_BattleSpec_To_v1alpha1_BattleSpec is an autogenerated conversion function.
func Convert_spirits_BattleSpec_To_v1alpha1_BattleSpec(in *spirits.BattleSpec, out *BattleSpec, s conversion.Scope) error {
	return autoConvert_spirits_BattleSpec_To_v1alpha1_BattleSpec(in, out, s)
}

func autoConvert_v1alpha1_BattleStatus_To_spirits_BattleStatus(in *BattleStatus, out *spirits.BattleStatus, s conversion.Scope) error {
	out.Conditions = *(*[]metav1.Condition)(unsafe.Pointer(&in.Conditions))
	out.Phase = spirits.BattlePhase(in.Phase)
	out.Message = in.Message
	out.InBattleSpirits = *(*[]v1.LocalObjectReference)(unsafe.Pointer(&in.InBattleSpirits))
	out.ActingSpirit = in.ActingSpirit
	return nil
}

// Convert_v1alpha1_BattleStatus_To_spirits_BattleStatus is an autogenerated conversion function.
func Convert_v1alpha1_BattleStatus_To_spirits_BattleStatus(in *BattleStatus, out *spirits.BattleStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_BattleStatus_To_spirits_BattleStatus(in, out, s)
}

func autoConvert_spirits_BattleStatus_To_v1alpha1_BattleStatus(in *spirits.BattleStatus, out *BattleStatus, s conversion.Scope) error {
	out.Conditions = *(*[]metav1.Condition)(unsafe.Pointer(&in.Conditions))
	out.Phase = BattlePhase(in.Phase)
	out.Message = in.Message
	out.InBattleSpirits = *(*[]v1.LocalObjectReference)(unsafe.Pointer(&in.InBattleSpirits))
	out.ActingSpirit = in.ActingSpirit
	return nil
}

// Convert_spirits_BattleStatus_To_v1alpha1_BattleStatus is an autogenerated conversion function.
func Convert_spirits_BattleStatus_To_v1alpha1_BattleStatus(in *spirits.BattleStatus, out *BattleStatus, s conversion.Scope) error {
	return autoConvert_spirits_BattleStatus_To_v1alpha1_BattleStatus(in, out, s)
}

func autoConvert_v1alpha1_Spirit_To_spirits_Spirit(in *Spirit, out *spirits.Spirit, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_SpiritSpec_To_spirits_SpiritSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_SpiritStatus_To_spirits_SpiritStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Spirit_To_spirits_Spirit is an autogenerated conversion function.
func Convert_v1alpha1_Spirit_To_spirits_Spirit(in *Spirit, out *spirits.Spirit, s conversion.Scope) error {
	return autoConvert_v1alpha1_Spirit_To_spirits_Spirit(in, out, s)
}

func autoConvert_spirits_Spirit_To_v1alpha1_Spirit(in *spirits.Spirit, out *Spirit, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_spirits_SpiritSpec_To_v1alpha1_SpiritSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_spirits_SpiritStatus_To_v1alpha1_SpiritStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_spirits_Spirit_To_v1alpha1_Spirit is an autogenerated conversion function.
func Convert_spirits_Spirit_To_v1alpha1_Spirit(in *spirits.Spirit, out *Spirit, s conversion.Scope) error {
	return autoConvert_spirits_Spirit_To_v1alpha1_Spirit(in, out, s)
}

func autoConvert_v1alpha1_SpiritList_To_spirits_SpiritList(in *SpiritList, out *spirits.SpiritList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]spirits.Spirit, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_Spirit_To_spirits_Spirit(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha1_SpiritList_To_spirits_SpiritList is an autogenerated conversion function.
func Convert_v1alpha1_SpiritList_To_spirits_SpiritList(in *SpiritList, out *spirits.SpiritList, s conversion.Scope) error {
	return autoConvert_v1alpha1_SpiritList_To_spirits_SpiritList(in, out, s)
}

func autoConvert_spirits_SpiritList_To_v1alpha1_SpiritList(in *spirits.SpiritList, out *SpiritList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Spirit, len(*in))
		for i := range *in {
			if err := Convert_spirits_Spirit_To_v1alpha1_Spirit(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_spirits_SpiritList_To_v1alpha1_SpiritList is an autogenerated conversion function.
func Convert_spirits_SpiritList_To_v1alpha1_SpiritList(in *spirits.SpiritList, out *SpiritList, s conversion.Scope) error {
	return autoConvert_spirits_SpiritList_To_v1alpha1_SpiritList(in, out, s)
}

func autoConvert_v1alpha1_SpiritSpec_To_spirits_SpiritSpec(in *SpiritSpec, out *spirits.SpiritSpec, s conversion.Scope) error {
	if err := Convert_v1alpha1_SpiritStats_To_spirits_SpiritStats(&in.Stats, &out.Stats, s); err != nil {
		return err
	}
	out.Actions = *(*[]string)(unsafe.Pointer(&in.Actions))
	out.Intelligence = spirits.SpiritIntelligence(in.Intelligence)
	out.Attributes = *(*map[string]string)(unsafe.Pointer(&in.Attributes))
	return nil
}

// Convert_v1alpha1_SpiritSpec_To_spirits_SpiritSpec is an autogenerated conversion function.
func Convert_v1alpha1_SpiritSpec_To_spirits_SpiritSpec(in *SpiritSpec, out *spirits.SpiritSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_SpiritSpec_To_spirits_SpiritSpec(in, out, s)
}

func autoConvert_spirits_SpiritSpec_To_v1alpha1_SpiritSpec(in *spirits.SpiritSpec, out *SpiritSpec, s conversion.Scope) error {
	if err := Convert_spirits_SpiritStats_To_v1alpha1_SpiritStats(&in.Stats, &out.Stats, s); err != nil {
		return err
	}
	out.Actions = *(*[]string)(unsafe.Pointer(&in.Actions))
	out.Intelligence = SpiritIntelligence(in.Intelligence)
	out.Attributes = *(*map[string]string)(unsafe.Pointer(&in.Attributes))
	// WARNING: in.Internal requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha1_SpiritStats_To_spirits_SpiritStats(in *SpiritStats, out *spirits.SpiritStats, s conversion.Scope) error {
	out.Health = in.Health
	out.Power = in.Power
	out.Armor = in.Armor
	out.Agility = in.Agility
	return nil
}

// Convert_v1alpha1_SpiritStats_To_spirits_SpiritStats is an autogenerated conversion function.
func Convert_v1alpha1_SpiritStats_To_spirits_SpiritStats(in *SpiritStats, out *spirits.SpiritStats, s conversion.Scope) error {
	return autoConvert_v1alpha1_SpiritStats_To_spirits_SpiritStats(in, out, s)
}

func autoConvert_spirits_SpiritStats_To_v1alpha1_SpiritStats(in *spirits.SpiritStats, out *SpiritStats, s conversion.Scope) error {
	out.Health = in.Health
	out.Power = in.Power
	out.Armor = in.Armor
	out.Agility = in.Agility
	return nil
}

// Convert_spirits_SpiritStats_To_v1alpha1_SpiritStats is an autogenerated conversion function.
func Convert_spirits_SpiritStats_To_v1alpha1_SpiritStats(in *spirits.SpiritStats, out *SpiritStats, s conversion.Scope) error {
	return autoConvert_spirits_SpiritStats_To_v1alpha1_SpiritStats(in, out, s)
}

func autoConvert_v1alpha1_SpiritStatus_To_spirits_SpiritStatus(in *SpiritStatus, out *spirits.SpiritStatus, s conversion.Scope) error {
	out.Conditions = *(*[]metav1.Condition)(unsafe.Pointer(&in.Conditions))
	out.Phase = spirits.SpiritPhase(in.Phase)
	return nil
}

// Convert_v1alpha1_SpiritStatus_To_spirits_SpiritStatus is an autogenerated conversion function.
func Convert_v1alpha1_SpiritStatus_To_spirits_SpiritStatus(in *SpiritStatus, out *spirits.SpiritStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_SpiritStatus_To_spirits_SpiritStatus(in, out, s)
}

func autoConvert_spirits_SpiritStatus_To_v1alpha1_SpiritStatus(in *spirits.SpiritStatus, out *SpiritStatus, s conversion.Scope) error {
	out.Conditions = *(*[]metav1.Condition)(unsafe.Pointer(&in.Conditions))
	out.Phase = SpiritPhase(in.Phase)
	return nil
}

// Convert_spirits_SpiritStatus_To_v1alpha1_SpiritStatus is an autogenerated conversion function.
func Convert_spirits_SpiritStatus_To_v1alpha1_SpiritStatus(in *spirits.SpiritStatus, out *SpiritStatus, s conversion.Scope) error {
	return autoConvert_spirits_SpiritStatus_To_v1alpha1_SpiritStatus(in, out, s)
}

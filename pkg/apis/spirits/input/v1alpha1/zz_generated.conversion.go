//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	input "github.com/ankeesler/spirits/internal/apis/spirits/input"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*ActionCall)(nil), (*input.ActionCall)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ActionCall_To_input_ActionCall(a.(*ActionCall), b.(*input.ActionCall), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*input.ActionCall)(nil), (*ActionCall)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_input_ActionCall_To_v1alpha1_ActionCall(a.(*input.ActionCall), b.(*ActionCall), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ActionCallSpec)(nil), (*input.ActionCallSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ActionCallSpec_To_input_ActionCallSpec(a.(*ActionCallSpec), b.(*input.ActionCallSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*input.ActionCallSpec)(nil), (*ActionCallSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_input_ActionCallSpec_To_v1alpha1_ActionCallSpec(a.(*input.ActionCallSpec), b.(*ActionCallSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ActionCallStatus)(nil), (*input.ActionCallStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ActionCallStatus_To_input_ActionCallStatus(a.(*ActionCallStatus), b.(*input.ActionCallStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*input.ActionCallStatus)(nil), (*ActionCallStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_input_ActionCallStatus_To_v1alpha1_ActionCallStatus(a.(*input.ActionCallStatus), b.(*ActionCallStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_ActionCall_To_input_ActionCall(in *ActionCall, out *input.ActionCall, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ActionCallSpec_To_input_ActionCallSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ActionCallStatus_To_input_ActionCallStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ActionCall_To_input_ActionCall is an autogenerated conversion function.
func Convert_v1alpha1_ActionCall_To_input_ActionCall(in *ActionCall, out *input.ActionCall, s conversion.Scope) error {
	return autoConvert_v1alpha1_ActionCall_To_input_ActionCall(in, out, s)
}

func autoConvert_input_ActionCall_To_v1alpha1_ActionCall(in *input.ActionCall, out *ActionCall, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_input_ActionCallSpec_To_v1alpha1_ActionCallSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_input_ActionCallStatus_To_v1alpha1_ActionCallStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_input_ActionCall_To_v1alpha1_ActionCall is an autogenerated conversion function.
func Convert_input_ActionCall_To_v1alpha1_ActionCall(in *input.ActionCall, out *ActionCall, s conversion.Scope) error {
	return autoConvert_input_ActionCall_To_v1alpha1_ActionCall(in, out, s)
}

func autoConvert_v1alpha1_ActionCallSpec_To_input_ActionCallSpec(in *ActionCallSpec, out *input.ActionCallSpec, s conversion.Scope) error {
	out.Battle = in.Battle
	out.Spirit = in.Spirit
	out.ActionName = in.ActionName
	return nil
}

// Convert_v1alpha1_ActionCallSpec_To_input_ActionCallSpec is an autogenerated conversion function.
func Convert_v1alpha1_ActionCallSpec_To_input_ActionCallSpec(in *ActionCallSpec, out *input.ActionCallSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ActionCallSpec_To_input_ActionCallSpec(in, out, s)
}

func autoConvert_input_ActionCallSpec_To_v1alpha1_ActionCallSpec(in *input.ActionCallSpec, out *ActionCallSpec, s conversion.Scope) error {
	out.Battle = in.Battle
	out.Spirit = in.Spirit
	out.ActionName = in.ActionName
	return nil
}

// Convert_input_ActionCallSpec_To_v1alpha1_ActionCallSpec is an autogenerated conversion function.
func Convert_input_ActionCallSpec_To_v1alpha1_ActionCallSpec(in *input.ActionCallSpec, out *ActionCallSpec, s conversion.Scope) error {
	return autoConvert_input_ActionCallSpec_To_v1alpha1_ActionCallSpec(in, out, s)
}

func autoConvert_v1alpha1_ActionCallStatus_To_input_ActionCallStatus(in *ActionCallStatus, out *input.ActionCallStatus, s conversion.Scope) error {
	out.Result = input.ActionCallResult(in.Result)
	out.Message = in.Message
	return nil
}

// Convert_v1alpha1_ActionCallStatus_To_input_ActionCallStatus is an autogenerated conversion function.
func Convert_v1alpha1_ActionCallStatus_To_input_ActionCallStatus(in *ActionCallStatus, out *input.ActionCallStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ActionCallStatus_To_input_ActionCallStatus(in, out, s)
}

func autoConvert_input_ActionCallStatus_To_v1alpha1_ActionCallStatus(in *input.ActionCallStatus, out *ActionCallStatus, s conversion.Scope) error {
	out.Result = ActionCallResult(in.Result)
	out.Message = in.Message
	return nil
}

// Convert_input_ActionCallStatus_To_v1alpha1_ActionCallStatus is an autogenerated conversion function.
func Convert_input_ActionCallStatus_To_v1alpha1_ActionCallStatus(in *input.ActionCallStatus, out *ActionCallStatus, s conversion.Scope) error {
	return autoConvert_input_ActionCallStatus_To_v1alpha1_ActionCallStatus(in, out, s)
}

//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package spirits

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Battle) DeepCopyInto(out *Battle) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Battle.
func (in *Battle) DeepCopy() *Battle {
	if in == nil {
		return nil
	}
	out := new(Battle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Battle) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BattleList) DeepCopyInto(out *BattleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Battle, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BattleList.
func (in *BattleList) DeepCopy() *BattleList {
	if in == nil {
		return nil
	}
	out := new(BattleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BattleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BattleSpec) DeepCopyInto(out *BattleSpec) {
	*out = *in
	if in.Spirits != nil {
		in, out := &in.Spirits, &out.Spirits
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BattleSpec.
func (in *BattleSpec) DeepCopy() *BattleSpec {
	if in == nil {
		return nil
	}
	out := new(BattleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BattleStatus) DeepCopyInto(out *BattleStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InBattleSpirits != nil {
		in, out := &in.InBattleSpirits, &out.InBattleSpirits
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	out.ActingSpirit = in.ActingSpirit
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BattleStatus.
func (in *BattleStatus) DeepCopy() *BattleStatus {
	if in == nil {
		return nil
	}
	out := new(BattleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Spirit) DeepCopyInto(out *Spirit) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Spirit.
func (in *Spirit) DeepCopy() *Spirit {
	if in == nil {
		return nil
	}
	out := new(Spirit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Spirit) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiritList) DeepCopyInto(out *SpiritList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Spirit, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiritList.
func (in *SpiritList) DeepCopy() *SpiritList {
	if in == nil {
		return nil
	}
	out := new(SpiritList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpiritList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiritSpec) DeepCopyInto(out *SpiritSpec) {
	*out = *in
	out.Stats = in.Stats
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Internal.DeepCopyInto(&out.Internal)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiritSpec.
func (in *SpiritSpec) DeepCopy() *SpiritSpec {
	if in == nil {
		return nil
	}
	out := new(SpiritSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiritSpecInternal) DeepCopyInto(out *SpiritSpecInternal) {
	*out = *in
	if in.Action != nil {
		out.Action = in.Action.DeepCopyAction()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiritSpecInternal.
func (in *SpiritSpecInternal) DeepCopy() *SpiritSpecInternal {
	if in == nil {
		return nil
	}
	out := new(SpiritSpecInternal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiritStats) DeepCopyInto(out *SpiritStats) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiritStats.
func (in *SpiritStats) DeepCopy() *SpiritStats {
	if in == nil {
		return nil
	}
	out := new(SpiritStats)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiritStatus) DeepCopyInto(out *SpiritStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiritStatus.
func (in *SpiritStatus) DeepCopy() *SpiritStatus {
	if in == nil {
		return nil
	}
	out := new(SpiritStatus)
	in.DeepCopyInto(out)
	return out
}

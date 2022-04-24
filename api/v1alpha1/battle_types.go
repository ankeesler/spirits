/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BattleSpec defines the desired state of Battle
type BattleSpec struct {
	// Spirits are the spirits involved in this Battle
	// +kubebuilder:validation:MinItems=2
	// +kubebuilder:validation:MaxItems=2
	Spirits []string `json:"spirits"`
}

// BattleStatus defines the observed state of Battle
type BattleStatus struct {
	// Phase summarizes the overall status of the Battle
	// +kubebuilder:default=Pending
	// +kubebuilder:validation:Enum=Pending;Ready;Error
	Phase Phase `json:"phase,omitempty"`

	// Conditions represents the observations of a Spirit's current state
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Battle is the Schema for the battles API
// +kubebuilder:resource:categories=spiritsworld
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.spec.status.phase`
type Battle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BattleSpec   `json:"spec,omitempty"`
	Status BattleStatus `json:"status,omitempty"`
}

var _ Object = &Battle{}

func (b *Battle) Conditions() *[]metav1.Condition {
	return &b.Status.Conditions
}

//+kubebuilder:object:root=true

// BattleList contains a list of Battle
type BattleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Battle `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Battle{}, &BattleList{})
}

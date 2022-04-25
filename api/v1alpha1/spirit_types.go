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

type SpiritIntelligence string

const (
	// SpiritIntelligenceRoundRobin is the default intelligence for a Spirit.
	// It describes a Spirit that performs actions in a sequential, deterministic order.
	SpiritIntelligenceRoundRobin SpiritIntelligence = "RoundRobin"

	// SpiritIntelligenceRandom describes a Spirit that performs actions in a random order.
	SpiritIntelligenceRandom SpiritIntelligence = "Random"

	// SpiritIntelligenceHuman describes a Spirit whose actions are driven by human interaction.
	SpiritIntelligenceHuman SpiritIntelligence = "Human"
)

// SpiritStats are quantitative properties of the Spirit.
// These are utilized and manipulated throughout the course of a Battle.
type SpiritStats struct {
	// Health is a quantitative representation of the energy of the Spirit.
	// When this drops to 0, the Spirit is no longer able to participate in a Battle.
	// +kubebuilder:validation:Minimum=1
	Health int `json:"health"`

	// Power is a quantitative representation of the attacking ability of the Spirit.
	// +kubebuilder:validation:Minimum=0
	// +optional
	Power int `json:"power"`

	// Armor is a quantitative representation of the defending ability of the Spirit.
	// +optional
	Armor int `json:"armor"`

	// Agility is a quantitative representation of the speed of the Spirit.
	// +optional
	Agility int `json:"agility"`
}

// SpiritSpec defines the desired state of Spirit
type SpiritSpec struct {
	// Stats are the current statistics that describe this Spirit
	Stats SpiritStats `json:"stats"`

	// Actions are the list of actions that this Spirit can perform
	// +kubebuilder:default={attack}
	// +optional
	Actions []string `json:"actions"`

	// Intelligence describes how a Spirit will select actions to perform
	// +kubebuilder:default=RoundRobin
	// +optional
	Intelligence SpiritIntelligence `json:"intelligence"`
}

// SpiritStatus defines the observed state of Spirit
type SpiritStatus struct {
	// Phase summarizes the overall status of the Spirit
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

// Spirit is the Schema for the spirits API
// +kubebuilder:resource:categories=spiritsworld
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Health",type=string,JSONPath=`.spec.stats.health`
type Spirit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpiritSpec   `json:"spec,omitempty"`
	Status SpiritStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SpiritList contains a list of Spirit
type SpiritList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Spirit `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Spirit{}, &SpiritList{})
}

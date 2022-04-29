package spirits

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

type SpiritPhase string

const (
	// SpiritPhasePending is the default phase for newly-created resources.
	SpiritPhasePending SpiritPhase = "Pending"

	// SpiritPhaseReady is the phase for a resource in a healthy state.
	SpiritPhaseReady SpiritPhase = "Ready"

	// SpiritPhaseError is the phase for a in an unhealthy state.
	SpiritPhaseError SpiritPhase = "Error"
)

// SpiritStats are quantitative properties of the Spirit.
// These are utilized and manipulated throughout the course of a Battle.
type SpiritStats struct {
	// Health is a quantitative representation of the energy of the Spirit.
	// When this drops to 0, the Spirit is no longer able to participate in a Battle.
	Health int `json:"health"`

	// Power is a quantitative representation of the attacking ability of the Spirit.
	// +optional
	Power int `json:"power"`

	// Armor is a quantitative representation of the defending ability of the Spirit.
	// +optional
	Armor int `json:"armor"`

	// Agility is a quantitative representation of the speed of the Spirit.
	// +optional
	Agility int `json:"agility"`
}

type SpiritSpecInternal struct {
	Action Action
}

// SpiritSpec defines the desired state of Spirit
type SpiritSpec struct {
	// Stats are the current statistics that describe this Spirit
	Stats SpiritStats `json:"stats"`

	// Actions are the list of actions that this Spirit can perform
	// +optional
	Actions []string `json:"actions"`

	// Intelligence describes how a Spirit will select actions to perform
	// +optional
	Intelligence SpiritIntelligence `json:"intelligence"`

	Internal SpiritSpecInternal
}

// SpiritStatus defines the observed state of Spirit
type SpiritStatus struct {
	// Conditions represents the observations of a Spirit's current state
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`

	// Phase summarizes the overall status of the Spirit
	Phase SpiritPhase `json:"phase,omitempty"`
}

// Spirit is the Schema for the spirits API
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Spirit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpiritSpec   `json:"spec,omitempty"`
	Status SpiritStatus `json:"status,omitempty"`
}

// SpiritList contains a list of Spirit
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SpiritList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Spirit `json:"items"`
}

package spirits

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SpiritIntelligence string

const (
	SpiritIntelligenceRoundRobin SpiritIntelligence = "RoundRobin"
	SpiritIntelligenceRandom     SpiritIntelligence = "Random"
	SpiritIntelligenceHuman      SpiritIntelligence = "Human"
)

type SpiritPhase string

const (
	SpiritPhasePending SpiritPhase = "Pending"
	SpiritPhaseReady   SpiritPhase = "Ready"
	SpiritPhaseError   SpiritPhase = "Error"
)

type SpiritStats struct {
	Health  int
	Power   int
	Armor   int
	Agility int
}

type SpiritSpecInternal struct {
	Action Action
}

type SpiritSpec struct {
	Stats        SpiritStats
	Actions      []string
	Intelligence SpiritIntelligence
	Attributes   map[string]string
	Internal     SpiritSpecInternal
}

type SpiritStatus struct {
	Conditions []metav1.Condition
	Phase      SpiritPhase
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Spirit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpiritSpec   `json:"spec,omitempty"`
	Status SpiritStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SpiritList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Spirit `json:"items"`
}

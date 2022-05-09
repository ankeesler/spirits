package spirits

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SpiritWellKnownAction string

const (
	SpiritWellKnownActionAttack SpiritWellKnownAction = "Attack"
	SpiritWellKnownActionNoop   SpiritWellKnownAction = "Noop"
)

type SpiritActionChoicesIntelligence string

const (
	SpiritActionChoicesIntelligenceRoundRobin SpiritActionChoicesIntelligence = "RoundRobin"
	SpiritActionChoicesIntelligenceRandom     SpiritActionChoicesIntelligence = "Random"
	SpiritActionChoicesIntelligenceHuman      SpiritActionChoicesIntelligence = "Human"
)

type SpiritPhase string

const (
	SpiritPhasePending SpiritPhase = "Pending"
	SpiritPhaseReady   SpiritPhase = "Ready"
	SpiritPhaseError   SpiritPhase = "Error"
)

type SpiritStats struct {
	Health  int64
	Power   int64
	Armor   int64
	Agility int64
}

type SpiritAttributes struct {
	Stats SpiritStats `json:"stats,omitempty"`
}

type SpiritActionChoices struct {
	Intelligence SpiritActionChoicesIntelligence `json:"intelligence"`
	Actions      map[string]*SpiritAction        `json:"action"`
}

type SpiritAction struct {
	WellKnown *SpiritWellKnownAction `json:"wellKnown,omitempty"`
	Choices   *SpiritActionChoices   `json:"choices,omitempty"`
	Script    *Script                `json:"script,omitempty"`
	Registry  *HTTP                  `json:"registry,omitempty"`
}

type SpiritSpecInternal struct {
	Action Action
}

type SpiritSpec struct {
	Attributes SpiritAttributes
	Action     SpiritAction
	Internal   SpiritSpecInternal
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

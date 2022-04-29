package spirits

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BattlePhase string

const (
	BattlePhasePending  BattlePhase = "Pending"
	BattlePhaseRunning  BattlePhase = "Running"
	BattlePhaseFinished BattlePhase = "Finished"
	BattlePhaseError    BattlePhase = "Error"
)

// BattleSpec defines the desired state of Battle
type BattleSpec struct {
	// Spirits are the spirits involved in this Battle
	Spirits []string `json:"spirits"`
}

// BattleStatus defines the observed state of Battle
type BattleStatus struct {
	// Conditions represents the observations of a Battle's current state
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`

	// Phase describes what stage of the battle lifecycle this Battle is in currently
	Phase BattlePhase `json:"phase"`

	Message string `json:"message"`

	// InBattleSpirits holds the names of the Spirit's that are participating in this Battle
	InBattleSpirits []string `json:"inBattleSpirits"`
}

// Battle is the Schema for the battles API
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Battle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BattleSpec   `json:"spec,omitempty"`
	Status BattleStatus `json:"status,omitempty"`
}

// BattleList contains a list of Battle
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type BattleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Battle `json:"items"`
}

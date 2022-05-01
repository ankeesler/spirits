package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BattlePhase string

const (
	BattlePhasePending        BattlePhase = "Pending"
	BattlePhaseRunning        BattlePhase = "Running"
	BattlePhaseAwaitingAction BattlePhase = "AwaitingAction"
	BattlePhaseFinished       BattlePhase = "Finished"
	BattlePhaseError          BattlePhase = "Error"
)

// BattleSpec defines the desired state of Battle
type BattleSpec struct {
	// Spirits are the spirits involved in this Battle
	// +kubebuilder:validation:MinItems=2
	// +kubebuilder:validation:MaxItems=2
	Spirits []corev1.LocalObjectReference `json:"spirits"`
}

// BattleStatus defines the observed state of Battle
type BattleStatus struct {
	// Conditions represents the observations of a Battle's current state
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`

	// Phase describes what stage of the battle lifecycle this Battle is in currently.
	// +kubebuilder:default=Pending
	// +kubebuilder:validation:Enum=Pending;Running;Finished;Error
	Phase BattlePhase `json:"phase"`

	// Message describes the reason for the Battle's Phase
	// +optional
	Message string `json:"message,omitempty"`

	// InBattleSpirits holds the names of the Spirit's that are participating in this Battle
	// +optional
	InBattleSpirits []corev1.LocalObjectReference `json:"inBattleSpirits,omitempty"`

	// ActingSpirit holds the name of the Spirit that is currently acting in this Battle
	// +optional
	ActingSpirit corev1.LocalObjectReference `json:"actingSpirit,omitempty"`
}

// Battle is the Schema for the battles API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=spiritsworld
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:subresource:status
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

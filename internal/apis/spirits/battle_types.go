package spirits

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BattlePhase string

const (
	BattlePhasePending  BattlePhase = "Pending"
	BattlePhaseRunning  BattlePhase = "Running"
	BattlePhaseFinished BattlePhase = "Finished"
	BattlePhaseError    BattlePhase = "Error"
)

type BattleSpec struct {
	Spirits []corev1.LocalObjectReference `json:"spirits"`
}

type BattleStatus struct {
	Conditions      []metav1.Condition            `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
	Phase           BattlePhase                   `json:"phase"`
	Message         string                        `json:"message"`
	InBattleSpirits []corev1.LocalObjectReference `json:"inBattleSpirits,omitempty"`
	ActingSpirit    corev1.LocalObjectReference   `json:"actingSpirit,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Battle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BattleSpec   `json:"spec,omitempty"`
	Status BattleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type BattleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Battle `json:"items"`
}

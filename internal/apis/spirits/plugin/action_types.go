package plugin

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
)

type ActionRunSpec struct {
	From spiritsinternal.SpiritSpec `json:"from"`
	To   spiritsinternal.SpiritSpec `json:"to"`
}

type ActionRunStatus struct {
	From spiritsinternal.SpiritSpec `json:"from"`
	To   spiritsinternal.SpiritSpec `json:"to"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ActionRun struct {
	metav1.TypeMeta `json:",inline"`

	Spec   ActionRunSpec   `json:"spec,omitempty"`
	Status ActionRunStatus `json:"status,omitempty"`
}

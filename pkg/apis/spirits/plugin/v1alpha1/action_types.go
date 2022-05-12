package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
)

// ActionRunSpec defines the input on which an Action will be performed
type ActionRunSpec struct {
	// From is the Spirit performing this Action
	// +k8s:conversion-gen=false
	From spiritsv1alpha1.SpiritSpec `json:"from"`
	// To is the Spirit receiving this Action
	// +k8s:conversion-gen=false
	To spiritsv1alpha1.SpiritSpec `json:"to"`
}

// ActionRunStatus defines the output after an Action has been performed
type ActionRunStatus struct {
	// From is the Spirit that performed the Action
	// +k8s:conversion-gen=false
	From spiritsv1alpha1.SpiritSpec `json:"from"`
	// To is the Spirit that received the Action
	// +k8s:conversion-gen=false
	To spiritsv1alpha1.SpiritSpec `json:"to"`
}

// ActionRun is a request/response handled by an entity that defines an Action's execution
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ActionRun struct {
	metav1.TypeMeta `json:",inline"`

	Spec   ActionRunSpec   `json:"spec,omitempty"`
	Status ActionRunStatus `json:"status,omitempty"`
}

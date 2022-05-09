package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ActionCallResult string

const (
	ActionCallResultAccepted ActionCallResult = "Accepted"
	ActionCallResultRejected ActionCallResult = "Rejected"
)

// ActionCallSpec defines the desired state of an ActionCall
type ActionCallSpec struct {
	// Battle references the Battle to which this ActionCall should be applied
	Battle corev1.LocalObjectReference `json:"battle"`
	// Spirit references the Spirit to which this ActionCall should be applied
	Spirit corev1.LocalObjectReference `json:"spirit"`

	// ActionName is the name of the actual Action being requested to run.
	ActionName string `json:"actionName"`
}

// ActionStatus defines the observed state of an ActionCall
type ActionCallStatus struct {
	// Result contains a summary of the ActionCall processing
	// +optional
	Result ActionCallResult `json:"result,omitempty"`
	// Message contains an optional description of an ActionCall's Result
	// +optional
	Message string `json:"message,omitempty"`
}

// ActionCall is the Schema for the ActionCall API
// +genclient
// +genclient:onlyVerbs=create
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ActionCall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ActionCallSpec   `json:"spec,omitempty"`
	Status ActionCallStatus `json:"status,omitempty"`
}

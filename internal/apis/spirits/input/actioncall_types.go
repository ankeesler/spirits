package input

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ActionCallResult string

const (
	ActionCallResultAccepted ActionCallResult = "Accepted"
	ActionCallResultRejected ActionCallResult = "Rejected"
)

type ActionCallSpec struct {
	Battle     corev1.LocalObjectReference `json:"battle"`
	Spirit     corev1.LocalObjectReference `json:"spirit"`
	ActionName string                      `json:"actionName"`
}

type ActionCallStatus struct {
	Result  ActionCallResult `json:"result,omitempty"`
	Message string           `json:"message,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ActionCall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ActionCallSpec   `json:"spec,omitempty"`
	Status ActionCallStatus `json:"status,omitempty"`
}

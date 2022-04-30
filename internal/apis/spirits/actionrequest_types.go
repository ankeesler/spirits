package spirits

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ActionRequestResult string

const (
	ActionRequestResultAccepted ActionRequestResult = "Accepted"
	ActionRequestResultRejected ActionRequestResult = "Rejected"
)

type ActionRequestLocalObjectGenerationReference struct {
	Name       string `json:"name"`
	Generation int64  `json:"generation"`
}

type ActionRequestSpec struct {
	Battle ActionRequestLocalObjectGenerationReference `json:"battle"`
	Spirit ActionRequestLocalObjectGenerationReference `json:"spirit"`

	ActionName string `json:"actionName"`
}

type ActionRequestStatus struct {
	Result  ActionRequestResult `json:"result"`
	Message string              `json:"message,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ActionRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ActionRequestSpec   `json:"spec,omitempty"`
	Status ActionRequestStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ActionRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Battle `json:"items"`
}

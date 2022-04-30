package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ActionRequestResult string

const (
	ActionRequestResultAccepted ActionRequestResult = "Accepted"
	ActionRequestResultRejected ActionRequestResult = "Rejected"
)

// ActionRequestLocalObjectGenerationReference defines the information an ActionRequest
// needs to locate the specific generation of another object
type ActionRequestLocalObjectGenerationReference struct {
	// Name is the name of the object
	Name string `json:"name"`
	// Generation is the generation of the object
	Generation int64 `json:"generation"`
}

// ActionRequestSpec defines the desired state of an ActionRequest
type ActionRequestSpec struct {
	// Battle references the Battle to which this ActionRequest should be applied
	Battle ActionRequestLocalObjectGenerationReference `json:"battle"`
	// Spirit references the Spirit to which this ActionRequest should be applied
	Spirit ActionRequestLocalObjectGenerationReference `json:"spirit"`

	// ActionName is the name of the actual Action being requested to run.
	ActionName string `json:"actionName"`
}

// ActionStatus defines the observed state of an ActionRequest
type ActionRequestStatus struct {
	// Result contains a summary of the ActionRequest processing
	// +kubebuilder:validation:Enum=Accepted;Rejected
	Result ActionRequestResult `json:"result"`
	// Message contains an optional description of an ActionRequest's Result
	// +optional
	Message string `json:"message,omitempty"`
}

// ActionRequest is the Schema for the ActionRequest API
// +genclient
// +genclient:onlyVerbs=create
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ActionRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ActionRequestSpec   `json:"spec,omitempty"`
	Status ActionRequestStatus `json:"status,omitempty"`
}

// ActionRequestList contains a list of ActionRequest's
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ActionRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Battle `json:"items"`
}

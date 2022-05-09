package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SpiritWellKnownAction string

const (
	// SpiritWellKnownActionAttack uses the Spirit's Power, less the opponent's Armor,
	// to decrement the opponent's Health, i.e., to.health -= from.power - to.armor
	SpiritWellKnownActionAttack SpiritWellKnownAction = "Attack"

	// SpiritWellKnownActionNoop does nothing
	SpiritWellKnownActionNoop SpiritWellKnownAction = "Noop"
)

type SpiritActionChoicesIntelligence string

const (
	// SpiritActionChoicesIntelligenceRoundRobin is the default intelligence for a Spirit.
	// It describes a Spirit that performs actions in a sequential, deterministic order
	SpiritActionChoicesIntelligenceRoundRobin SpiritActionChoicesIntelligence = "RoundRobin"

	// SpiritActionChoicesIntelligenceRandom describes a Spirit that performs actions in a random order
	SpiritActionChoicesIntelligenceRandom SpiritActionChoicesIntelligence = "Random"

	// SpiritActionChoicesIntelligenceHuman describes a Spirit whose actions are driven by human interaction
	SpiritActionChoicesIntelligenceHuman SpiritActionChoicesIntelligence = "Human"
)

type SpiritPhase string

const (
	// SpiritPhasePending is the default phase for newly-created resources.
	SpiritPhasePending SpiritPhase = "Pending"

	// SpiritPhaseReady is the phase for a resource in a healthy state.
	SpiritPhaseReady SpiritPhase = "Ready"

	// SpiritPhaseError is the phase for a in an unhealthy state.
	SpiritPhaseError SpiritPhase = "Error"
)

// SpiritStats are quantitative properties of the Spirit.
// These are utilized and manipulated throughout the course of a Battle.
type SpiritStats struct {
	// Health is a quantitative representation of the energy of the Spirit.
	// When this drops to 0, the Spirit is no longer able to participate in a Battle.
	// +kubebuilder:validation:Minimum=0
	// +optional
	Health int64 `json:"health"`

	// Power is a quantitative representation of the attacking ability of the Spirit.
	// +kubebuilder:validation:Minimum=0
	// +optional
	Power int64 `json:"power"`

	// Armor is a quantitative representation of the defending ability of the Spirit.
	// +kubebuilder:validation:Minimum=0
	// +optional
	Armor int64 `json:"armor"`

	// Agility is a quantitative representation of the speed of the Spirit.
	// +kubebuilder:validation:Minimum=0
	// +optional
	Agility int64 `json:"agility"`
}

// SpiritAttributes are the list of qualities of a Spirit
type SpiritAttributes struct {
	// Stats are the current statistics that describe this Spirit
	// +optional
	Stats SpiritStats `json:"stats"`
}

// NamedSpiritAction is a type holding an Action and a unique name used to reference that Action
type NamedSpiritAction struct {
	// Name is a unique name used to reference this Action
	Name string `json:"name"`

	// Action is the SpiritAction describing this Action
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	Action SpiritAction `json:"action"`
}

// SpiritActionChoices is a list of SpiritAction's that a Spirit could chose to perform;
// A Spirit choses one of them to perform for their Action
type SpiritActionChoices struct {
	// Intelligence holds the strategy via which one of the below Action's will be chosen
	Intelligence SpiritActionChoicesIntelligence `json:"intelligence"`

	// Actions are the SpiritAction's from which the Spirit can chose
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=name
	Actions []NamedSpiritAction `json:"actions" patchStrategy:"merge" patchMergeKey:"name"`
}

// SpiritAction describes the Action that a Spirit performs
// +kubebuilder:validation:MaxProperties=1
// +kubebuilder:validation:MinProperties=1
type SpiritAction struct {
	// WellKnown is a string identifying a well-known (builtin) Action
	// +kubebuilder:validation:Enum=Attack;Noop
	// +optional
	WellKnown *SpiritWellKnownAction `json:"wellKnown,omitempty"`

	// Choices specify a list of Action's a Spirit could potentially perform
	// +optional
	Choices *SpiritActionChoices `json:"choices,omitempty"`

	// Script holds the source of a script used to implement a Spirit's Action
	// +optional
	Script *Script `json:"script,omitempty"`

	// Registry holds the HTTP information to GET an Action
	// +optional
	Registry *HTTP `json:"registry,omitempty"`
}

// SpiritSpec defines the desired state of Spirit
type SpiritSpec struct {
	// Attributes describe the in-battle spirit's qualities
	// +optional
	Attributes SpiritAttributes `json:"attributes,omitempty"`

	// Action contains the description of the Action this Spirit performs
	Action SpiritAction `json:"action"`
}

// SpiritStatus defines the observed state of Spirit
type SpiritStatus struct {
	// Conditions represents the observations of a Spirit's current in-API state
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`

	// Phase summarizes the overall status of the Spirit API object
	// +kubebuilder:default=Pending
	// +kubebuilder:validation:Enum=Pending;Ready;Error
	Phase SpiritPhase `json:"phase,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Spirit is the Schema for the spirits API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=spiritsworld
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Health",type=integer,JSONPath=`.spec.attributes.stats.health`
// +kubebuilder:subresource:status
type Spirit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpiritSpec   `json:"spec,omitempty"`
	Status SpiritStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SpiritList contains a list of Spirit
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SpiritList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Spirit `json:"items"`
}

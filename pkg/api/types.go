package api

type Phase string

const (
	// PhasePending is the default phase for newly-created resources.
	PhasePending Phase = "Pending"

	// PhaseReady is the phase for a resource in a healthy state.
	PhaseReady Phase = "Ready"

	// PhaseError is the phase for a in an unhealthy state.
	PhaseError Phase = "Error"
)

package battle

type State string

const (
	StatePending   State = "pending"
	StateStarted   State = "started"
	StateWaiting   State = "waiting"
	StateFinished  State = "finished"
	StateCancelled State = "cancelled"
	StateError     State = "error"
)

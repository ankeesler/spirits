package battle

type State string

const (
	StatePending   State = "pending"
	StateStarted   State = "started"
	StateRunning   State = "running"
	StateWaiting   State = "waiting"
	StateFinished  State = "finished"
	StateCancelled State = "cancelled"
	StateError     State = "error"
)

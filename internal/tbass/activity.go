package tbass

// Activity drives Teams of Actor's actions.
type Activity interface {
	// Play performs the Activity amongst the provided Teams.
	Play(Teams)
}

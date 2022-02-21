package tbass

// Actor performs some action on other Actor's.
//
// Actor's have Stat's, which are used to record actions on other Actor's.
type Actor interface {
	// Name returns the name of this Actor.
	Name() string
	// Act performs this Actor's action on the provided Team's.
	//
	// Teams[0] is the Actor's Team. Teams[1:] contains the rest of the Team's in the Activity,
	// including the Actor's Team.
	Act(Teams)
	// Stat gets the stats of this Actor.
	Stat(string) Stat

	// Clone returns a deep-copy of this Actor.
	Clone() Actor
}

// StatValue is the number type that is used for Stat's. It should be able to be changed to the
// desired number type of choice.
type StatValue int

// Stat measures some aspect of an Actor.
type Stat interface {
	// Name returns the name of the Stat.
	Name() string
	// Get loads the Stat's current value.
	Get() StatValue
	// Set stores the Stat's current value.
	Set(StatValue)
}

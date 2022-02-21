// Package activity provides a simple implementation of a tbass.Activity.
//
// Here are some invariants of this tbass.Activity implementation:
//   - At the end of an Activity, there are 0 or more winning Team's
//   - Only one Actor acts at one time
package activity

import "github.com/ankeesler/spirits/internal/tbass"

// Config describes the Activity.
type Config struct {
	// GetStrategy provides the Strategy for the Activity. This is called at the beginning of
	// Activity.Play() to create a fresh Strategy for each Activity.
	//
	// If left unset, the Activty will use a default Strategy, which loops through the Actor's across
	// the Teams once, and then returns all Teams as the winner.
	GetStrategy func(tbass.Teams) Strategy

	// Listener is an optional field that can receive callbacks from when an Activity is played.
	Listener Listener
}

// Strategy provides the core logic for an Activity.
type Strategy interface {
	// Next returns the next Actor (and associated Team) to act.
	Next() (tbass.Actor, tbass.Team)

	// Winner returns the winner of the Activity. This is called once at the top each Activity loop.
	// If Winner returns nil, then the Activity will continue.
	Winner() tbass.Teams
}

type defaultStrategy struct {
	tt tbass.Teams
	i  int
}

func (s *defaultStrategy) Next() (tbass.Actor, tbass.Team) {
	i := 0
	for _, t := range s.tt {
		for _, a := range t.Actors() {
			if i == s.i {
				s.i++
				return a, t
			}
		}
	}
	return nil, nil
}

func (s *defaultStrategy) Winner() tbass.Teams {
	return s.tt
}

// Listener is a set of callbacks to provide information from an in-progress Activity.
type Listener interface {
	// OnBegin is called begin any Actor's act. It is called with all of the Teams in the Activity.
	OnBegin(tbass.Teams)
	// OnTurn is called after each Actor acts.
	OnTurn(tbass.Actor)
	// OnEnd is called at the end of the Activity with the winning Teams of the Activity.
	OnEnd(tbass.Teams)
}

type defaultListener struct{}

func (l defaultListener) OnBegin(_ tbass.Teams) {}
func (l defaultListener) OnTurn(_ tbass.Actor)  {}
func (l defaultListener) OnEnd(_ tbass.Teams)   {}

type activity struct {
	getStrategy func(tbass.Teams) Strategy
	l           Listener
}

// New creates a new Activity with the provided Config.
func New(c *Config) tbass.Activity {
	a := &activity{}

	a.getStrategy = c.GetStrategy
	if a.getStrategy == nil {
		a.getStrategy = func(tt tbass.Teams) Strategy { return &defaultStrategy{tt: tt} }
	}

	a.l = c.Listener
	if a.l == nil {
		a.l = defaultListener{}
	}

	return a
}

func (a *activity) Play(tt tbass.Teams) {
	a.l.OnBegin(tt)

	s := a.getStrategy(tt)

	for {
		if winner := s.Winner(); winner != nil {
			a.l.OnEnd(winner)
			return
		}

		actor, t := s.Next()
		if actor == nil {
			break
		}

		actor.Act(append(tbass.Teams{t}, tt...))
		a.l.OnTurn(actor)

		// TODO: detect loop?
	}
}

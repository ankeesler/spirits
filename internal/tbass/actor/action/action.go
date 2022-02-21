// Package action provides helpful actor.Action utilities.
package action

import "github.com/ankeesler/spirits/internal/tbass"

// ActionFunc is a helper type that makes it easy to use a function to define an actor.Action.
type ActionFunc func(tbass.Actor, tbass.Team, tbass.Teams)

func (f ActionFunc) Act(me tbass.Actor, us tbass.Team, all tbass.Teams) { f(me, us, all) }

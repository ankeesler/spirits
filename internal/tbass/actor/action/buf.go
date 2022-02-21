package action

import (
	"github.com/ankeesler/spirits/internal/tbass"
	"github.com/ankeesler/spirits/internal/tbass/actor"
)

// Buf uses fromStat on the tbass.Actor performing this actor.Action to increase the toStat on an
// allied tbass.Actor/s.
//
// If all is true, then all allies will be targeted. Otherwise, the allied tbass.Actor with the
// lowest fromStat will be targeted.
func Buf(fromStat, toStat string, all bool) actor.Action {
	return ActionFunc(func(me tbass.Actor, us tbass.Team, them tbass.Teams) {
		power := me.Stat(fromStat).Get()

		var targets []tbass.Actor
		if all {
			targets = us.Actors()
		} else {
			targets = []tbass.Actor{findMin(us.Actors(), toStat)}
		}

		for _, target := range targets {
			s := target.Stat(toStat)
			s.Set(s.Get() + power)
		}
	})
}

// Debuf uses fromStat on the tbass.Actor performing this actor.Action to decrease the toStat on an
// allied tbass.Actor/s.
//
// If all is true, then all opponents will be targeted. Otherwise, the opposing tbass.Actor with
// the lowest fromStat will be targeted.
func Debuf(fromStat, toStat string, all bool) actor.Action {
	return ActionFunc(func(me tbass.Actor, us tbass.Team, them tbass.Teams) {
		power := me.Stat(fromStat).Get()

		var targets []tbass.Actor
		if all {
			targets = them.Actors()
		} else {
			targets = []tbass.Actor{findMin(them.Actors(), toStat)}
		}

		for _, target := range targets {
			s := target.Stat(toStat)
			s.Set(s.Get() - power)
		}
	})
}

func bufOrDebuf(fromStat, toStat string, team bool, buf bool) actor.Action {
	return nil
}

func findMin(aa []tbass.Actor, stat string) tbass.Actor {
	var minStat tbass.StatValue
	var minActor tbass.Actor
	for _, a := range aa {
		if minActor == nil || a.Stat(stat).Get() < minStat {
			minStat = a.Stat(stat).Get()
			minActor = a
		}
	}
	return minActor
}

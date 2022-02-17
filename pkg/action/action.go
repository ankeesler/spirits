// Package action provides common implementations of spirit.Action.
package action

import (
	"github.com/ankeesler/spirits/pkg/spirit"
)

type actionFunc func(*spirit.ActionContext)

func (f actionFunc) Run(ac *spirit.ActionContext) { f(ac) }

// Attack returns a spirit.Action that decreases a target's health.
func Attack(target Target) spirit.Action {
	return actionFunc(func(ac *spirit.ActionContext) {
		for _, targetSpirit := range target.find(ac) {
			damage := ac.Me.Power() - targetSpirit.Armour()
			targetSpirit.DecreaseHealth(damage)
		}
	})
}

// Heal returns a spirit.Action that increases a target's health.
func Heal(target Target) spirit.Action {
	return actionFunc(func(ac *spirit.ActionContext) {
		for _, targetSpirit := range target.find(ac) {
			targetSpirit.IncreaseHealth(ac.Me.Power())
		}
	})
}

// Buf returns a spirit.Action that increases a target's power, armour, and/or agility.
func Buf(target Target) spirit.Action {
	return actionFunc(func(ac *spirit.ActionContext) {
		for _, targetSpirit := range target.find(ac) {
			if target&TargetPower != 0 {
				targetSpirit.IncreasePower(ac.Me.Power())
			}
			if target&TargetArmour != 0 {
				targetSpirit.IncreaseArmour(ac.Me.Power())
			}
			if target&TargetAgility != 0 {
				targetSpirit.IncreaseAgility(ac.Me.Power())
			}
		}
	})
}

// Debuf returns a spirit.Action that decreases a target's power, armour, and/or agility.
func Debuf(target Target) spirit.Action {
	return actionFunc(func(ac *spirit.ActionContext) {
		for _, targetSpirit := range target.find(ac) {
			if target&TargetPower != 0 {
				targetSpirit.DecreasePower(ac.Me.Power())
			}
			if target&TargetArmour != 0 {
				targetSpirit.DecreaseArmour(ac.Me.Power())
			}
			if target&TargetAgility != 0 {
				targetSpirit.DecreaseAgility(ac.Me.Power())
			}
		}
	})
}

package battle

import (
	"fmt"

	"github.com/ankeesler/spirits/pkg/api"
)

type State string

const (
	StatePending   State = "pending"
	StateStarted         = "started"
	StateWaiting         = "waiting"
	StateFinished        = "finished"
	StateCancelled       = "cancelled"
	StateError           = "error"
)

func stateFromAPI(apiBattleState api.BattleState) State {
	switch apiBattleState {
	case api.BattleState_BATTLE_STATE_PENDING:
		return StatePending
	case api.BattleState_BATTLE_STATE_STARTED:
		return StateStarted
	case api.BattleState_BATTLE_STATE_WAITING:
		return StateWaiting
	case api.BattleState_BATTLE_STATE_FINISHED:
		return StateFinished
	case api.BattleState_BATTLE_STATE_CANCELLED:
		return StateCancelled
	case api.BattleState_BATTLE_STATE_ERROR:
		return StateError
	default:
		panic(fmt.Sprintf("unknown battle state type: %d", apiBattleState))
	}
}

func (s State) ToAPI() api.BattleState {
	switch s {
	case StatePending:
		return api.BattleState_BATTLE_STATE_PENDING
	case StateStarted:
		return api.BattleState_BATTLE_STATE_STARTED
	case StateWaiting:
		return api.BattleState_BATTLE_STATE_WAITING
	case StateFinished:
		return api.BattleState_BATTLE_STATE_FINISHED
	case StateCancelled:
		return api.BattleState_BATTLE_STATE_CANCELLED
	case StateError:
		return api.BattleState_BATTLE_STATE_ERROR
	default:
		panic(fmt.Sprintf("unknown battle state type: %s", s))
	}
}

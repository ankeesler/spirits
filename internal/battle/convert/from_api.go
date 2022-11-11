package convert

import (
	"errors"
	"fmt"

	battlepkg "github.com/ankeesler/spirits/internal/battle"
	convertmeta "github.com/ankeesler/spirits/internal/meta/convert"
	convertspirit "github.com/ankeesler/spirits/internal/spirit/convert"
	"github.com/ankeesler/spirits/pkg/api"
)

func FromAPI(apiBattle *api.Battle, actionSource battlepkg.ActionSource) (*battlepkg.Battle, error) {
	internalMeta := convertmeta.FromAPI(apiBattle.GetMeta())
	internalBattle := battlepkg.New(internalMeta)

	internalBattle.SetState(stateFromAPI(apiBattle.GetState()))
	if errorMessage := apiBattle.ErrorMessage; errorMessage != nil {
		internalBattle.SetErr(errors.New(*errorMessage))
	}

	for _, apiTeam := range apiBattle.GetTeams() {
		teamName := apiTeam.GetName()
		if internalBattle.Team(teamName) != nil {
			return nil, fmt.Errorf("duplicate team name %s", teamName)
		}

		internalBattle.AddTeam(teamName)

		for _, apiTeamSpirit := range apiTeam.GetSpirits() {
			internalSpirit, err := convertspirit.FromAPI(apiTeamSpirit.GetSpirit())
			if err != nil {
				return nil, fmt.Errorf("invalid spirit %s: %w", apiTeamSpirit.GetSpirit().GetMeta().GetId(), err)
			}

			internalBattle.AddTeamSpirit(
				teamName,
				internalSpirit,
				SpiritIntelligenceFromAPI(apiTeamSpirit.GetIntelligence()),
				apiTeamSpirit.GetSeed(),
				actionSource,
			)
		}

		internalBattle.SetTurns(apiBattle.GetTurns())
	}

	return internalBattle, nil
}

func stateFromAPI(apiBattleState api.BattleState) battlepkg.State {
	switch apiBattleState {
	case api.BattleState_BATTLE_STATE_PENDING:
		return battlepkg.StatePending
	case api.BattleState_BATTLE_STATE_STARTED:
		return battlepkg.StateStarted
	case api.BattleState_BATTLE_STATE_WAITING:
		return battlepkg.StateWaiting
	case api.BattleState_BATTLE_STATE_FINISHED:
		return battlepkg.StateFinished
	case api.BattleState_BATTLE_STATE_CANCELLED:
		return battlepkg.StateCancelled
	case api.BattleState_BATTLE_STATE_ERROR:
		return battlepkg.StateError
	default:
		panic(fmt.Sprintf("unknown api battle state type: %d", apiBattleState))
	}
}

func SpiritIntelligenceFromAPI(
	apiBattleSpiritIntelligence api.BattleTeamSpiritIntelligence,
) battlepkg.SpiritIntelligence {
	switch apiBattleSpiritIntelligence {
	case api.BattleTeamSpiritIntelligence_BATTLE_TEAM_SPIRIT_INTELLIGENCE_HUMAN:
		return battlepkg.SpiritIntelligenceHuman
	case api.BattleTeamSpiritIntelligence_BATTLE_TEAM_SPIRIT_INTELLIGENCE_RANDOM:
		return battlepkg.SpiritIntelligenceRandom
	default:
		panic(fmt.Sprintf("unknown api spirit intelligence type: %d", apiBattleSpiritIntelligence))
	}
}

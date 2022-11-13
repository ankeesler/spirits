package convert

import (
	"fmt"

	battlepkg "github.com/ankeesler/spirits/internal/battle"
	convertmeta "github.com/ankeesler/spirits/internal/meta/convert"
	convertspirit "github.com/ankeesler/spirits/internal/spirit/convert"
	spiritsv1 "github.com/ankeesler/spirits/pkg/api/spirits/v1"
)

func ToAPI(internalBattle *battlepkg.Battle) *spiritsv1.Battle {
	apiBattle := &spiritsv1.Battle{
		Meta:  convertmeta.ToAPI(internalBattle.Meta),
		State: stateToAPI(internalBattle.State()),
	}

	if err := internalBattle.Err(); err != nil {
		errorMessage := err.Error()
		apiBattle.ErrorMessage = &errorMessage
	}

	for _, teamName := range internalBattle.TeamNames() {
		teamToAPI(teamName, internalBattle.Team(teamName), &apiBattle.Teams)
		teamToAPI(teamName, internalBattle.InBattleTeam(teamName), &apiBattle.InBattleTeams)
	}

	if next := internalBattle.PeekNext(); next != nil {
		apiBattle.NextSpiritIds = []string{next.ID()}
	}
	apiBattle.Turns = internalBattle.Turns()

	return apiBattle
}

func teamToAPI(teamName string, internalTeamSpirits []*battlepkg.Spirit, apiTeams *[]*spiritsv1.BattleTeam) {
	apiBattleTeam := &spiritsv1.BattleTeam{
		Name: teamName,
	}

	for _, internalTeamSpirit := range internalTeamSpirits {
		apiBattleTeam.Spirits = append(apiBattleTeam.Spirits, &spiritsv1.BattleTeamSpirit{
			Spirit:       convertspirit.ToAPI(internalTeamSpirit.Spirit),
			Intelligence: spiritIntelligenceToAPI(internalTeamSpirit.Intelligence()),
			Seed:         internalTeamSpirit.Seed(),
		})
	}

	*apiTeams = append(*apiTeams, apiBattleTeam)
}

func stateToAPI(internalState battlepkg.State) spiritsv1.BattleState {
	switch internalState {
	case battlepkg.StatePending:
		return spiritsv1.BattleState_BATTLE_STATE_PENDING
	case battlepkg.StateStarted:
		return spiritsv1.BattleState_BATTLE_STATE_STARTED
	case battlepkg.StateWaiting:
		return spiritsv1.BattleState_BATTLE_STATE_WAITING
	case battlepkg.StateFinished:
		return spiritsv1.BattleState_BATTLE_STATE_FINISHED
	case battlepkg.StateCancelled:
		return spiritsv1.BattleState_BATTLE_STATE_CANCELLED
	case battlepkg.StateError:
		return spiritsv1.BattleState_BATTLE_STATE_ERROR
	default:
		panic(fmt.Sprintf("unknown battle state type: %s", internalState))
	}
}

func spiritIntelligenceToAPI(
	internalSpiritIntelligence battlepkg.SpiritIntelligence,
) spiritsv1.BattleTeamSpiritIntelligence {
	switch internalSpiritIntelligence {
	case battlepkg.SpiritIntelligenceHuman:
		return spiritsv1.BattleTeamSpiritIntelligence_BATTLE_TEAM_SPIRIT_INTELLIGENCE_HUMAN
	case battlepkg.SpiritIntelligenceRandom:
		return spiritsv1.BattleTeamSpiritIntelligence_BATTLE_TEAM_SPIRIT_INTELLIGENCE_RANDOM
	default:
		panic(fmt.Sprintf("unknown api spirit intelligence type: %s", internalSpiritIntelligence))
	}
}

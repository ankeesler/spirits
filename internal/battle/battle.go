package battle

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	metapkg "github.com/ankeesler/spirits/internal/meta"
	spiritpkg "github.com/ankeesler/spirits/internal/spirit"
	"github.com/ankeesler/spirits/pkg/api"
)

type Battle struct {
	apiBattle *api.Battle

	*metapkg.Meta

	state State
	err   error

	teams         map[string][]*spiritpkg.Spirit
	inBattleTeams map[string][]*spiritpkg.Spirit

	spiritTeams map[string]string

	queue queue

	turns int64
}

type ActionSource interface {
	Pend(context.Context, string, string, int64) (string, []string, error)
}

func FromAPI(
	ctx context.Context,
	apiBattle *api.Battle,
	actionRepo spiritpkg.ActionRepo,
	actionSource ActionSource,
) (*Battle, error) {
	internalBattle := &Battle{
		apiBattle: apiBattle,

		Meta: metapkg.FromAPI(apiBattle.GetMeta()),

		state: stateFromAPI(apiBattle.GetState()),

		teams:         make(map[string][]*spiritpkg.Spirit),
		inBattleTeams: make(map[string][]*spiritpkg.Spirit),

		spiritTeams: make(map[string]string),
	}

	if message := apiBattle.ErrorMessage; message != nil {
		internalBattle.err = errors.New(*message)
	}

	addAPIBattleTeams(
		ctx,
		apiBattle.GetTeams(),
		actionRepo,
		actionSource,
		internalBattle,
		internalBattle.teams,
	)

	addAPIBattleTeams(
		ctx,
		apiBattle.GetTeams(),
		actionRepo,
		actionSource,
		internalBattle,
		internalBattle.inBattleTeams,
	)

	for internalTeamName, internalSpirits := range internalBattle.teams {
		for _, internalSpirit := range internalSpirits {
			internalBattle.spiritTeams[internalSpirit.ID()] = internalTeamName
		}
	}

	for _, internalSpirits := range internalBattle.teams {
		for _, internalSpirit := range internalSpirits {
			internalBattle.queue.AddSpirit(internalSpirit)
		}
	}

	internalBattle.turns = apiBattle.GetTurns()

	return internalBattle, nil
}

func (b *Battle) State() State         { return b.state }
func (b *Battle) SetState(state State) { b.state = state }

func (b *Battle) SetError(err error) { b.err = err }

func (b *Battle) TeamNames() []string {
	var teamNames []string
	for teamName := range b.teams {
		teamNames = append(teamNames, teamName)
	}
	return teamNames
}

func (b *Battle) Team(name string) []*spiritpkg.Spirit {
	return b.teams[name]
}

func (b *Battle) AddTeam(name string) {
	b.teams[name] = make([]*spiritpkg.Spirit, 0)
}

func (b *Battle) AddTeamSpirit(name string, spirit *spiritpkg.Spirit) {
	b.teams[name] = append(b.teams[name], spirit)
}

func (b *Battle) InBattleTeam(name string) []*spiritpkg.Spirit {
	return b.inBattleTeams[name]
}

func (b *Battle) HasNext() bool {
	healthyTeams := 0
	for _, team := range b.inBattleTeams {
		healthySpirits := 0
		for _, spirit := range team {
			if spirit.Health() > 0 {
				healthySpirits++
			}
		}
		if healthySpirits > 0 {
			healthyTeams++
		}
	}
	return healthyTeams > 1
}

func (b *Battle) Next() (*spiritpkg.Spirit, []*spiritpkg.Spirit, [][]*spiritpkg.Spirit) {
	me := b.queue.Next()
	myTeamName := b.spiritTeams[me.ID()]
	us := b.inBattleTeams[myTeamName]
	var them [][]*spiritpkg.Spirit
	for teamName, teamSpirits := range b.inBattleTeams {
		if teamName != myTeamName {
			them = append(them, teamSpirits)
		}
	}
	return me, us, them
}

func (b *Battle) Turns() int64 { return b.turns }

func (b *Battle) ToAPI() *api.Battle {
	apiBattle := &api.Battle{
		Meta: b.Meta.ToAPI(),

		State: b.State().ToAPI(),

		NextSpiritIds: b.queue.NextIDs(),

		Turns: b.Turns(),
	}

	if b.err != nil {
		errorMessage := b.err.Error()
		apiBattle.ErrorMessage = &errorMessage
	}

	addInternalBattleTeam(b.teams, &apiBattle.Teams)
	addInternalBattleTeam(b.inBattleTeams, &apiBattle.InBattleTeams)

	return apiBattle
}

func addAPIBattleTeams(
	ctx context.Context,
	apiBattleTeams []*api.BattleTeam,
	actionRepo spiritpkg.ActionRepo,
	actionSource ActionSource,
	internalBattle *Battle,
	internalTeams map[string][]*spiritpkg.Spirit,
) error {
	for _, apiBattleTeam := range apiBattleTeams {
		for _, apiBattleTeamSpirit := range apiBattleTeam.GetSpirits() {
			var spiritActionSource spiritpkg.ActionSource
			switch apiBattleTeamSpirit.GetIntelligence() {
			case api.BattleTeamSpiritIntelligence_BATTLE_TEAM_SPIRIT_INTELLIGENCE_HUMAN:
				spiritActionSource = spiritActionSourceFunc(func(
					ctx context.Context,
					me *spiritpkg.Spirit,
					us []*spiritpkg.Spirit,
					them [][]*spiritpkg.Spirit,
				) (string, []string, error) {
					return actionSource.Pend(ctx, internalBattle.ID(), me.ID(), internalBattle.Turns())
				})
			case api.BattleTeamSpiritIntelligence_BATTLE_TEAM_SPIRIT_INTELLIGENCE_RANDOM:
				spiritActionSource = randomActionSource{
					r: rand.New(rand.NewSource(apiBattleTeamSpirit.GetSeed())),
				}
			}

			apiSpirit := apiBattleTeamSpirit.GetSpirit()
			internalSpirit, err := spiritpkg.FromAPI(
				ctx,
				apiSpirit,
				actionRepo,
				spiritActionSource,
			)
			if err != nil {
				return fmt.Errorf("convert spirit %s from API: %w", apiSpirit.GetMeta().GetId(), err)
			}

			teamName := apiBattleTeam.GetName()
			internalTeams[teamName] = append(internalTeams[teamName], internalSpirit)
		}
	}
	return nil
}

func addInternalBattleTeam(
	internalTeams map[string][]*spiritpkg.Spirit, apiTeam *[]*api.BattleTeam) {
	for teamName, internalTeamSpirits := range internalTeams {
		apiBattleTeam := &api.BattleTeam{
			Name: teamName,
		}

		for _, internalTeamSpirit := range internalTeamSpirits {
			apiBattleTeam.Spirits = append(apiBattleTeam.Spirits, &api.BattleTeamSpirit{
				Spirit: internalTeamSpirit.ToAPI(),
			})
		}

		*apiTeam = append(*apiTeam, apiBattleTeam)
	}
}

type spiritActionSourceFunc func(
	ctx context.Context,
	me *spiritpkg.Spirit,
	us []*spiritpkg.Spirit,
	them [][]*spiritpkg.Spirit,
) (string, []string, error)

func (f spiritActionSourceFunc) Pend(
	ctx context.Context,
	me *spiritpkg.Spirit,
	us []*spiritpkg.Spirit,
	them [][]*spiritpkg.Spirit,
) (string, []string, error) {
	return f(ctx, me, us, them)
}

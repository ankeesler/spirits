package battle

import (
	"context"
	"math/rand"

	metapkg "github.com/ankeesler/spirits/internal/meta"
	spiritpkg "github.com/ankeesler/spirits/internal/spirit"
)

type ActionSource interface {
	Pend(context.Context, string, string, int64) (string, []string, error)
}

type Battle struct {
	*metapkg.Meta

	state State
	err   error

	teams         map[string][]*Spirit
	inBattleTeams map[string][]*Spirit

	spiritTeams map[string]string

	queue queue

	turns int64
}

func New(meta *metapkg.Meta) *Battle {
	return &Battle{
		Meta: meta,

		teams:         make(map[string][]*Spirit),
		inBattleTeams: make(map[string][]*Spirit),

		spiritTeams: make(map[string]string),
	}
}

func (b *Battle) State() State         { return b.state }
func (b *Battle) SetState(state State) { b.state = state }

func (b *Battle) Err() error       { return b.err }
func (b *Battle) SetErr(err error) { b.err = err }

func (b *Battle) TeamNames() []string {
	var teamNames []string
	for teamName := range b.teams {
		teamNames = append(teamNames, teamName)
	}
	return teamNames
}

func (b *Battle) Team(name string) []*Spirit {
	return b.teams[name]
}

func (b *Battle) AddTeam(name string) {
	b.teams[name] = make([]*Spirit, 0)
	b.inBattleTeams[name] = make([]*Spirit, 0)
}

func (b *Battle) AddTeamSpirit(
	name string,
	spirit *spiritpkg.Spirit,
	intelligence SpiritIntelligence,
	seed int64,
	actionSource ActionSource,
) {
	battleSpirit := &Spirit{
		Spirit:       spirit,
		battle:       b,
		intelligence: intelligence,
		seed:         seed,
	}
	switch intelligence {
	case SpiritIntelligenceHuman:
		battleSpirit.actionSource = &humanActionSource{actionSource}
	case SpiritIntelligenceRandom:
		battleSpirit.actionSource = &randomActionSource{rand.New(rand.NewSource(seed))}
	}

	b.teams[name] = append(b.teams[name], battleSpirit)
	b.inBattleTeams[name] = append(b.inBattleTeams[name], battleSpirit)

	b.spiritTeams[battleSpirit.ID()] = name

	b.queue.AddSpirit(battleSpirit)
	b.queue.Init()
}

func (b *Battle) InBattleTeam(name string) []*Spirit {
	return b.inBattleTeams[name]
}

func (b *Battle) HasNext() bool {
	healthyTeams := 0
	for _, team := range b.inBattleTeams {
		healthySpirits := 0
		for _, internalSpirit := range team {
			if internalSpirit.Health() > 0 {
				healthySpirits++
			}
		}
		if healthySpirits > 0 {
			healthyTeams++
		}
	}
	return healthyTeams > 1
}

func (b *Battle) Next() (*Spirit, []*Spirit, [][]*Spirit) {
	me := b.queue.Next()
	myTeamName := b.spiritTeams[me.ID()]
	us := b.inBattleTeams[myTeamName]
	var them [][]*Spirit
	for teamName, teamSpirits := range b.inBattleTeams {
		if teamName != myTeamName {
			them = append(them, teamSpirits)
		}
	}
	return me, us, them
}

func (b *Battle) PeekNext() *Spirit {
	return b.queue.Peek()
}

func (b *Battle) Turns() int64         { return b.turns }
func (b *Battle) SetTurns(turns int64) { b.turns = turns }

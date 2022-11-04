package runner

import (
	"context"
	"errors"
	"fmt"

	battlepkg "github.com/ankeesler/spirits0/internal/battle"
	"github.com/ankeesler/spirits0/internal/spirit"
)

const maxTurns = 100

type Callback func(int, error)

type Runner struct {
	battle   *battlepkg.Battle
	callback Callback

	teams       map[string]*battlepkg.Team // Team name -> team
	spiritTeams map[string]*battlepkg.Team // Spirit ID -> team
}

func New(battle *battlepkg.Battle, callback Callback) *Runner {
	teams := make(map[string]*battlepkg.Team)
	spiritTeams := make(map[string]*battlepkg.Team)
	for _, team := range battle.Teams() {
		teams[team.Name()] = team
		for _, spirit := range team.Spirits() {
			spiritTeams[spirit.ID()] = team
		}
	}
	return &Runner{
		battle:   battle,
		callback: callback,

		teams:       teams,
		spiritTeams: spiritTeams,
	}
}

func (r *Runner) Run(ctx context.Context) {
	turn := 0
	for {
		var err error

		select {
		case <-ctx.Done():
			r.callback(turn, ctx.Err())
		default:
		}

		turn++
		if turn >= maxTurns {
			r.callback(turn, errors.New("too many turns"))
			break
		}

		if !r.battle.Queue().HasNext() {
			break
		}

		next := r.battle.Queue().Next()
		ctx, err = next.Act(ctx, r.myTeamSpirits(next), r.notMyTeamSpirits(next))
		if err != nil {
			r.callback(turn, fmt.Errorf("action errored: %w", err))
			break
		}

		r.callback(turn, nil)
	}
}

func (r *Runner) myTeamSpirits(next *spirit.Spirit) []*spirit.Spirit {
	return r.spiritTeams[next.ID()].Spirits()
}

func (r *Runner) notMyTeamSpirits(next *spirit.Spirit) [][]*spirit.Spirit {
	myTeam := r.spiritTeams[next.ID()]
	var teams [][]*spirit.Spirit
	for _, team := range r.teams {
		if team.Name() != myTeam.Name() {
			teams = append(teams, team.Spirits())
		}
	}
	return teams
}

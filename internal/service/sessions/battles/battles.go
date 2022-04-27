package battles

import (
	"context"
	"fmt"

	"github.com/ankeesler/spirits/internal/domain"
	"github.com/ankeesler/spirits/internal/domain/battle"
	internalbattle "github.com/ankeesler/spirits/internal/domain/battle"
	"github.com/ankeesler/spirits/internal/domain/session"
	"github.com/ankeesler/spirits/internal/domain/team"
	"github.com/ankeesler/spirits/internal/service"
	server "github.com/ankeesler/spirits/pkg/api/generated/server/api"
)

type Service struct {
	domain *domain.Domain
}

func New(domain *domain.Domain) *Service {
	return &Service{
		domain: domain,
	}
}

func (s *Service) CreateSessionBattle(ctx context.Context, sessionName string, battle server.Battle) (server.ImplResponse, error) {
	session, rsp, err := service.Find(ctx, sessionName, s.domain.Sessions)
	if session == nil {
		return rsp, err
	}
	return service.Create(ctx, &battle, session.Battles, s.converterFuncs(ctx, session))
}

func (s *Service) ListSessionBattles(ctx context.Context, sessionName string) (server.ImplResponse, error) {
	session, rsp, err := service.Find(ctx, sessionName, s.domain.Sessions)
	if session == nil {
		return rsp, err
	}
	return service.List(ctx, session.Battles, s.converterFuncs(ctx, session))
}

func (s *Service) GetSessionBattle(ctx context.Context, sessionName, battleName string) (server.ImplResponse, error) {
	session, rsp, err := service.Find(ctx, sessionName, s.domain.Sessions)
	if session == nil {
		return rsp, err
	}
	return service.Get(ctx, battleName, session.Battles, s.converterFuncs(ctx, session))
}

func (s *Service) DeleteSessionBattle(ctx context.Context, sessionName, battleName string) (server.ImplResponse, error) {
	session, rsp, err := service.Find(ctx, sessionName, s.domain.Sessions)
	if session == nil {
		return rsp, err
	}

	battle, rsp, err := service.Find(ctx, sessionName, session.Battles)
	if battle == nil {
		return rsp, err
	}

	if err := battle.Stop(); err != nil {
		return server.ImplResponse{}, fmt.Errorf("failed to stop battle: %w", err)
	}

	return service.Delete(ctx, battleName, session.Battles, s.converterFuncs(ctx, session))
}

func (s *Service) converterFuncs(ctx context.Context, session *session.Session) *service.ConverterFuncs[server.Battle, internalbattle.Battle] {
	return &service.ConverterFuncs[server.Battle, internalbattle.Battle]{
		AToB: func(battle *server.Battle) (*internalbattle.Battle, error) {
			return s.toInternalBattle(ctx, session, battle)
		},
		BToA: func(internalBattle *internalbattle.Battle) (*server.Battle, error) {
			teams := []string{}
			for _, inBattleTeam := range internalBattle.InBattleTeams {
				teams = append(teams, inBattleTeam.Name)
			}
			return &server.Battle{Name: internalBattle.Name, Teams: teams}, nil
		},
	}
}

func (s *Service) toInternalBattle(
	ctx context.Context,
	session *session.Session,
	battle *server.Battle,
) (*internalbattle.Battle, error) {
	var inBattleInternalTeams []*team.Team

	for _, teamName := range battle.Teams {
		inBattleInternalTeam, err := session.Teams.Get(ctx, teamName)
		if err != nil {
			return nil, fmt.Errorf("could not get internal team %q: %w", teamName, err)
		}
		inBattleInternalTeams = append(inBattleInternalTeams, inBattleInternalTeam)
	}

	// TODO: store spirits from teams in battle.Spirits
	// TODO: or should we have /battles/*/teams/*/spirits/*...
	//   /sessions/*/battles/*/teams/*/spirits/*/actions

	internalBattle := internalbattle.New(battle.Name, inBattleInternalTeams)
	if err := internalBattle.Start(s.battleCallback); err != nil {
		return nil, fmt.Errorf("failed to start internal battle: %w", err)
	}

	return internalBattle, nil
}

func (s *Service) battleCallback(internalBattle *battle.Battle, err error) {
}

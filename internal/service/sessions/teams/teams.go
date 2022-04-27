package teams

import (
	"context"

	"github.com/ankeesler/spirits/internal/domain"
	internalteam "github.com/ankeesler/spirits/internal/domain/team"
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

func (s *Service) CreateSessionTeam(ctx context.Context, sessionName string, team server.Team) (server.ImplResponse, error) {
	session, rsp, err := service.Find(ctx, sessionName, s.domain.Sessions)
	if session == nil {
		return rsp, err
	}
	return service.Create(ctx, &team, session.Teams, &converterFuncs)
}

func (s *Service) UpdateSessionTeam(ctx context.Context, sessionName, teamName string, team server.Team) (server.ImplResponse, error) {
	session, rsp, err := service.Find(ctx, sessionName, s.domain.Sessions)
	if session == nil {
		return rsp, err
	}
	return service.Update(ctx, &team, session.Teams, &converterFuncs)
}

func (s *Service) ListSessionTeams(ctx context.Context, sessionName string) (server.ImplResponse, error) {
	session, rsp, err := service.Find(ctx, sessionName, s.domain.Sessions)
	if session == nil {
		return rsp, err
	}
	return service.List(ctx, session.Teams, &converterFuncs)
}

func (s *Service) GetSessionTeam(ctx context.Context, sessionName, teamName string) (server.ImplResponse, error) {
	session, rsp, err := service.Find(ctx, sessionName, s.domain.Sessions)
	if session == nil {
		return rsp, err
	}
	return service.Get(ctx, teamName, session.Teams, &converterFuncs)
}

func (s *Service) DeleteSessionTeam(ctx context.Context, sessionName, teamName string) (server.ImplResponse, error) {
	session, rsp, err := service.Find(ctx, sessionName, s.domain.Sessions)
	if session == nil {
		return rsp, err
	}
	return service.Delete(ctx, teamName, session.Teams, &converterFuncs)
}

var converterFuncs = service.ConverterFuncs[server.Team, internalteam.Team]{
	From: func(team *server.Team) (*internalteam.Team, error) {
		return internalteam.New(team.Name), nil
	},
	To: func(internalTeam *internalteam.Team) (*server.Team, error) {
		return &server.Team{Name: internalTeam.Name}, nil
	},
}

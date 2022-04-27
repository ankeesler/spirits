package teams

import (
	"context"
	"errors"
	"net/http"

	"github.com/ankeesler/spirits/internal/domain"
	internalsession "github.com/ankeesler/spirits/internal/domain/session"
	internalteam "github.com/ankeesler/spirits/internal/domain/team"
	"github.com/ankeesler/spirits/internal/service"
	"github.com/ankeesler/spirits/internal/store"
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
	session, rsp, err := s.getSession(ctx, sessionName)
	if session == nil {
		return rsp, err
	}
	return service.Create(ctx, &team, session.Teams, &converterFuncs)
}

func (s *Service) UpdateSessionTeam(ctx context.Context, sessionName, teamName string, team server.Team) (server.ImplResponse, error) {
	session, rsp, err := s.getSession(ctx, sessionName)
	if session == nil {
		return rsp, err
	}
	return service.Update(ctx, &team, session.Teams, &converterFuncs)
}

func (s *Service) ListSessionTeams(ctx context.Context, sessionName string) (server.ImplResponse, error) {
	session, rsp, err := s.getSession(ctx, sessionName)
	if session == nil {
		return rsp, err
	}
	return service.List(ctx, session.Teams, &converterFuncs)
}

func (s *Service) GetSessionTeam(ctx context.Context, sessionName, teamName string) (server.ImplResponse, error) {
	session, rsp, err := s.getSession(ctx, sessionName)
	if session == nil {
		return rsp, err
	}
	return service.Get(ctx, teamName, session.Teams, &converterFuncs)
}

func (s *Service) DeleteSessionTeam(ctx context.Context, sessionName, teamName string) (server.ImplResponse, error) {
	session, rsp, err := s.getSession(ctx, sessionName)
	if session == nil {
		return rsp, err
	}
	return service.Delete(ctx, teamName, session.Teams, &converterFuncs)
}

func (s *Service) getSession(ctx context.Context, name string) (*internalsession.Session, server.ImplResponse, error) {
	session, err := s.domain.Sessions.Get(ctx, name)
	if err != nil {
		if errors.Is(err, &store.ErrNotFound{}) {
			return nil, server.ImplResponse{
				Code: http.StatusNotFound,
			}, nil
		}
		return nil, server.ImplResponse{}, err
	}
	return session, server.ImplResponse{}, nil
}

var converterFuncs = service.ConverterFuncs[server.Team, internalteam.Team]{
	From: func(team *server.Team) (*internalteam.Team, error) {
		return internalteam.New(team.Name), nil
	},
	To: func(internalTeam *internalteam.Team) (*server.Team, error) {
		return &server.Team{Name: internalTeam.Name}, nil
	},
}

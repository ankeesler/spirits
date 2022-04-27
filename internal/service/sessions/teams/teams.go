package teams

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/ankeesler/spirits/internal/domain"
	internalteam "github.com/ankeesler/spirits/internal/domain/team"
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
	internalTeam := toInternalTeam(&team)

	session, err := s.domain.Sessions.Get(ctx, sessionName)
	if err != nil {
		if errors.Is(err, &store.ErrNotFound{}) {
			return server.ImplResponse{
				Code: http.StatusNotFound,
			}, nil
		}
		return server.ImplResponse{}, err
	}

	internalTeam, err = session.Teams.Create(ctx, internalTeam)
	if err != nil {
		if errors.Is(err, &store.ErrAlreadyExists{}) {
			return server.ImplResponse{
				Code: http.StatusConflict,
			}, nil
		}
		return server.ImplResponse{}, err
	}

	return server.ImplResponse{
		Code: http.StatusCreated,
		Body: fromInternalTeam(internalTeam),
	}, nil
}

func (s *Service) UpdateSessionTeam(ctx context.Context, sessionName, teamName string, team server.Team) (server.ImplResponse, error) {
	internalTeam := toInternalTeam(&team)

	if teamName != team.Name {
		return server.ImplResponse{
			Code: http.StatusBadRequest,
			Body: fmt.Sprintf("path parameter  name %q does not match body name %q", teamName, team.Name),
		}, nil
	}

	session, err := s.domain.Sessions.Get(ctx, sessionName)
	if err != nil {
		if errors.Is(err, &store.ErrNotFound{}) {
			return server.ImplResponse{
				Code: http.StatusNotFound,
			}, nil
		}
		return server.ImplResponse{}, err
	}

	internalTeam, err = session.Teams.Update(ctx, internalTeam)
	if err != nil {
		if errors.Is(err, &store.ErrNotFound{}) {
			return server.ImplResponse{
				Code: http.StatusConflict,
			}, nil
		}
		return server.ImplResponse{}, err
	}

	return server.ImplResponse{
		Code: http.StatusCreated,
		Body: fromInternalTeam(internalTeam),
	}, nil
}

func (s *Service) ListSessionTeams(ctx context.Context, sessionName string) (server.ImplResponse, error) {
	session, err := s.domain.Sessions.Get(ctx, sessionName)
	if err != nil {
		if errors.Is(err, &store.ErrNotFound{}) {
			return server.ImplResponse{
				Code: http.StatusNotFound,
			}, nil
		}
		return server.ImplResponse{}, err
	}

	internalTeams, err := session.Teams.List(ctx)
	if err != nil {
		return server.ImplResponse{}, err
	}

	return server.ImplResponse{
		Code: http.StatusCreated,
		Body: fromInternalTeams(internalTeams),
	}, nil
}

func (s *Service) GetSessionTeam(ctx context.Context, sessionName, teamName string) (server.ImplResponse, error) {
	session, err := s.domain.Sessions.Get(ctx, sessionName)
	if err != nil {
		if errors.Is(err, &store.ErrNotFound{}) {
			return server.ImplResponse{
				Code: http.StatusNotFound,
			}, nil
		}
		return server.ImplResponse{}, err
	}

	internalSession, err := session.Teams.Get(ctx, teamName)
	if err != nil {
		if errors.Is(err, &store.ErrNotFound{}) {
			return server.ImplResponse{
				Code: http.StatusNotFound,
			}, nil
		}
		return server.ImplResponse{}, err
	}

	return server.ImplResponse{
		Code: http.StatusCreated,
		Body: fromInternalTeam(internalSession),
	}, nil
}

func (s *Service) DeleteSessionTeam(ctx context.Context, sessionName, teamName string) (server.ImplResponse, error) {
	session, err := s.domain.Sessions.Get(ctx, sessionName)
	if err != nil {
		if errors.Is(err, &store.ErrNotFound{}) {
			return server.ImplResponse{
				Code: http.StatusNotFound,
			}, nil
		}
		return server.ImplResponse{}, err
	}

	internalSession, err := session.Teams.Delete(ctx, teamName)
	if err != nil {
		if errors.Is(err, &store.ErrNotFound{}) {
			return server.ImplResponse{
				Code: http.StatusNotFound,
			}, nil
		}
		return server.ImplResponse{}, err
	}

	return server.ImplResponse{
		Code: http.StatusCreated,
		Body: fromInternalTeam(internalSession),
	}, nil
}

func toInternalTeam(team *server.Team) *internalteam.Team {
	return internalteam.New(team.Name)
}

func fromInternalTeams(internalTeams []*internalteam.Team) []*server.Team {
	teams := []*server.Team{}
	for _, internalTeam := range internalTeams {
		teams = append(teams, fromInternalTeam(internalTeam))
	}
	return teams
}

func fromInternalTeam(internalTeam *internalteam.Team) *server.Team {
	return &server.Team{Name: internalTeam.Name}
}

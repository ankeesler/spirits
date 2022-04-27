package sessions

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/ankeesler/spirits/internal/domain"
	internalsession "github.com/ankeesler/spirits/internal/domain/session"
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

func (s *Service) CreateSession(ctx context.Context, session server.Session) (server.ImplResponse, error) {
	internalSession := toInternalSession(&session)

	var err error
	internalSession, err = s.domain.Sessions.Create(ctx, internalSession)
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
		Body: fromInternalSession(internalSession),
	}, nil
}

func (s *Service) UpdateSession(ctx context.Context, name string, session server.Session) (server.ImplResponse, error) {
	internalSession := toInternalSession(&session)

	if name != session.Name {
		return server.ImplResponse{
			Code: http.StatusBadRequest,
			Body: fmt.Sprintf("path parameter  name %q does not match body name %q", name, session.Name),
		}, nil
	}

	var err error
	internalSession, err = s.domain.Sessions.Update(ctx, internalSession)
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
		Body: fromInternalSession(internalSession),
	}, nil
}

func (s *Service) ListSessions(ctx context.Context) (server.ImplResponse, error) {
	internalSessions, err := s.domain.Sessions.List(ctx)
	if err != nil {
		return server.ImplResponse{}, err
	}

	return server.ImplResponse{
		Code: http.StatusCreated,
		Body: fromInternalSessions(internalSessions),
	}, nil
}

func (s *Service) GetSession(ctx context.Context, name string) (server.ImplResponse, error) {
	internalSession, err := s.domain.Sessions.Get(ctx, name)
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
		Body: fromInternalSession(internalSession),
	}, nil
}

func (s *Service) DeleteSession(ctx context.Context, name string) (server.ImplResponse, error) {
	internalSession, err := s.domain.Sessions.Delete(ctx, name)
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
		Body: fromInternalSession(internalSession),
	}, nil
}

func toInternalSession(session *server.Session) *internalsession.Session {
	return internalsession.New(session.Name)
}

func fromInternalSessions(internalSessions []*internalsession.Session) []*server.Session {
	sessions := []*server.Session{}
	for _, internalSession := range internalSessions {
		sessions = append(sessions, fromInternalSession(internalSession))
	}
	return sessions
}

func fromInternalSession(internalSession *internalsession.Session) *server.Session {
	return &server.Session{Name: internalSession.Name}
}

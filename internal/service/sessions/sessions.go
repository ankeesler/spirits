package sessions

import (
	"context"

	"github.com/ankeesler/spirits/internal/domain"
	internalsession "github.com/ankeesler/spirits/internal/domain/session"
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

func (s *Service) CreateSession(ctx context.Context, session server.Session) (server.ImplResponse, error) {
	return service.Create(ctx, &session, s.domain.Sessions, &converterFuncs)
}

func (s *Service) UpdateSession(ctx context.Context, name string, session server.Session) (server.ImplResponse, error) {
	// TODO: check name matches session.Name
	return service.Update(ctx, &session, s.domain.Sessions, &converterFuncs)
}

func (s *Service) ListSessions(ctx context.Context) (server.ImplResponse, error) {
	return service.List(ctx, s.domain.Sessions, &converterFuncs)
}

func (s *Service) GetSession(ctx context.Context, name string) (server.ImplResponse, error) {
	return service.Get(ctx, name, s.domain.Sessions, &converterFuncs)
}

func (s *Service) DeleteSession(ctx context.Context, name string) (server.ImplResponse, error) {
	return service.Delete(ctx, name, s.domain.Sessions, &converterFuncs)
}

var converterFuncs = service.ConverterFuncs[server.Session, internalsession.Session]{
	From: func(session *server.Session) (*internalsession.Session, error) {
		return internalsession.New(session.Name), nil
	},
	To: func(internalSession *internalsession.Session) (*server.Session, error) {
		return &server.Session{Name: internalSession.Name}, nil
	},
}

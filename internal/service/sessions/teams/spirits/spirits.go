package spirits

import (
	"context"

	"github.com/ankeesler/spirits/internal/domain"
	internalspirit "github.com/ankeesler/spirits/internal/domain/spirit"
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

func (s *Service) CreateSessionTeamSpirit(ctx context.Context, sessionName, teamName string, spirit server.Spirit) (server.ImplResponse, error) {
	session, rsp, err := service.Find(ctx, sessionName, s.domain.Sessions)
	if session == nil {
		return rsp, err
	}
	team, rsp, err := service.Find(ctx, teamName, session.Teams)
	if team == nil {
		return rsp, err
	}
	return service.Create(ctx, &spirit, team.Spirits, &converterFuncs)
}

func (s *Service) UpdateSessionTeamSpirit(ctx context.Context, sessionName, teamName, spiritName string, spirit server.Spirit) (server.ImplResponse, error) {
	session, rsp, err := service.Find(ctx, sessionName, s.domain.Sessions)
	if session == nil {
		return rsp, err
	}
	team, rsp, err := service.Find(ctx, teamName, session.Teams)
	if team == nil {
		return rsp, err
	}
	// TODO: check spiritName matches spirit.Name
	return service.Update(ctx, &spirit, team.Spirits, &converterFuncs)
}

func (s *Service) ListSessionTeamSpirits(ctx context.Context, sessionName, teamName string) (server.ImplResponse, error) {
	session, rsp, err := service.Find(ctx, sessionName, s.domain.Sessions)
	if session == nil {
		return rsp, err
	}
	team, rsp, err := service.Find(ctx, teamName, session.Teams)
	if team == nil {
		return rsp, err
	}
	return service.List(ctx, team.Spirits, &converterFuncs)
}

func (s *Service) GetSessionTeamSpirit(ctx context.Context, sessionName, teamName, spiritName string) (server.ImplResponse, error) {
	session, rsp, err := service.Find(ctx, sessionName, s.domain.Sessions)
	if session == nil {
		return rsp, err
	}
	team, rsp, err := service.Find(ctx, teamName, session.Teams)
	if team == nil {
		return rsp, err
	}
	return service.Get(ctx, spiritName, team.Spirits, &converterFuncs)
}

func (s *Service) DeleteSessionTeamSpirit(ctx context.Context, sessionName, teamName, spiritName string) (server.ImplResponse, error) {
	session, rsp, err := service.Find(ctx, sessionName, s.domain.Sessions)
	if session == nil {
		return rsp, err
	}
	team, rsp, err := service.Find(ctx, teamName, session.Teams)
	if team == nil {
		return rsp, err
	}
	return service.Delete(ctx, spiritName, team.Spirits, &converterFuncs)
}

var converterFuncs = service.ConverterFuncs[server.Spirit, internalspirit.Spirit]{
	AToB: func(spirit *server.Spirit) (*internalspirit.Spirit, error) {
		return internalspirit.New(spirit.Name, 1), nil
	},
	BToA: func(internalSpirit *internalspirit.Spirit) (*server.Spirit, error) {
		return &server.Spirit{Name: internalSpirit.Name}, nil
	},
}

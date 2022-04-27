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

func (s *Service) ListSessionBattleSpirits(ctx context.Context, sessionName, battleName string) (server.ImplResponse, error) {
	session, rsp, err := service.Find(ctx, sessionName, s.domain.Sessions)
	if session == nil {
		return rsp, err
	}
	battle, rsp, err := service.Find(ctx, battleName, session.Battles)
	if battle == nil {
		return rsp, err
	}
	return service.List(ctx, battle.Spirits, &converterFuncs)
}

func (s *Service) GetSessionBattleSpirit(ctx context.Context, sessionName, battleName, spiritName string) (server.ImplResponse, error) {
	session, rsp, err := service.Find(ctx, sessionName, s.domain.Sessions)
	if session == nil {
		return rsp, err
	}
	battle, rsp, err := service.Find(ctx, battleName, session.Teams)
	if battle == nil {
		return rsp, err
	}
	return service.Get(ctx, spiritName, battle.Spirits, &converterFuncs)
}

// TODO: this is duplicated
var converterFuncs = service.ConverterFuncs[server.Spirit, internalspirit.Spirit]{
	AToB: func(spirit *server.Spirit) (*internalspirit.Spirit, error) {
		return internalspirit.New(spirit.Name, 1), nil
	},
	BToA: func(internalSpirit *internalspirit.Spirit) (*server.Spirit, error) {
		return &server.Spirit{Name: internalSpirit.Name}, nil
	},
}

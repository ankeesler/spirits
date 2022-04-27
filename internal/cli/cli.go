package cli

import (
	"context"
	"os"

	"github.com/ankeesler/spirits/internal/domain"
	"github.com/ankeesler/spirits/internal/log"
	internalserver "github.com/ankeesler/spirits/internal/server"
	"github.com/ankeesler/spirits/internal/service"
	"github.com/ankeesler/spirits/internal/service/sessions"
	"github.com/ankeesler/spirits/internal/service/sessions/battles"
	battlespirits "github.com/ankeesler/spirits/internal/service/sessions/battles/spirits"
	"github.com/ankeesler/spirits/internal/service/sessions/teams"
	teamspirits "github.com/ankeesler/spirits/internal/service/sessions/teams/spirits"
	"github.com/ankeesler/spirits/pkg/api"
	server "github.com/ankeesler/spirits/pkg/api/generated/server/api"
)

func Run() error {
	log.Info("spirits version " + api.Version)

	address := ":80"
	if port, ok := os.LookupEnv("PORT"); ok {
		address = ":" + port
	}

	defaultAPIService := service.NewDefault()
	defaultAPIController := server.NewDefaultApiController(defaultAPIService)

	domain := domain.New()

	sessionsAPIService := sessions.New(domain)
	sessionsAPIController := server.NewSessionApiController(sessionsAPIService)

	sessionTeamsAPIService := teams.New(domain)
	sessionTeamsAPIController := server.NewSessionTeamApiController(sessionTeamsAPIService)

	sessionTeamSpiritsAPIService := teamspirits.New(domain)
	sessionTeamSpiritsAPIController := server.NewSessionTeamSpiritApiController(sessionTeamSpiritsAPIService)

	sessionBattlesAPIService := battles.New(domain)
	sessionBattlesAPIController := server.NewSessionBattleApiController(sessionBattlesAPIService)

	sessionBattleSpiritsAPIService := battlespirits.New(domain)
	sessionBattleSpiritsAPIController := server.NewSessionBattleSpiritApiController(sessionBattleSpiritsAPIService)

	handler := server.NewRouter(
		defaultAPIController,
		sessionsAPIController,
		sessionTeamsAPIController,
		sessionTeamSpiritsAPIController,
		sessionBattlesAPIController,
		sessionBattleSpiritsAPIController,
	)

	return internalserver.Run(context.TODO(), address, handler)
}

package cli

import (
	"context"
	"os"

	"github.com/ankeesler/spirits/internal/log"
	internalserver "github.com/ankeesler/spirits/internal/server"
	"github.com/ankeesler/spirits/internal/service"
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
	handler := server.NewRouter(defaultAPIController)

	return internalserver.Run(context.TODO(), address, handler)
}

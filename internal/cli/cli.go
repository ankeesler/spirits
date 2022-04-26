package cli

import (
	"github.com/ankeesler/spirits/internal/log"
	"github.com/ankeesler/spirits/internal/server"
	"github.com/ankeesler/spirits/pkg/api"
)

func Run() error {
	log.Info("spirits version " + api.Version)
	return server.Run()
}

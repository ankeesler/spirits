package cli

import (
	"github.com/ankeesler/spirits/internal/log"
	"github.com/ankeesler/spirits/internal/server"
	"github.com/ankeesler/spirits/pkg/version"
)

func Run() error {
	log.Info("spirits version " + version.Version)
	return server.Run()
}

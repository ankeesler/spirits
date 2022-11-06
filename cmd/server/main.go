package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/ankeesler/spirits/internal/server"
)

var (
	spiritBuiltinDir = flag.String(
		"spirit-builtin-dir", "api/builtin/spirit", "Path to builtin spirit directory")
	actionBuiltinDir = flag.String(
		"action-builtin-dir", "api/builtin/action", "Path to builtin action directory")
)

func main() {
	flag.Parse()
	server, err := server.Wire(&server.Config{
		SpiritBuiltinDir: os.DirFS(*spiritBuiltinDir),
		ActionBuiltinDir: os.DirFS(*actionBuiltinDir),
	})
	if err != nil {
		log.Fatalf("failed to wire server: %v", err)
	}

	if err := server.Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

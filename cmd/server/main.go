package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/ankeesler/spirits/internal/server"
)

var (
	spiritBuiltinDir = flag.String(
		"spirit-builtin-dir", "api/builtin/spirit", "Path to builtin spirit directory")
	actionBuiltinDir = flag.String(
		"action-builtin-dir", "api/builtin/action", "Path to builtin action directory")
)

func main() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)

	flag.Parse()

	var port int
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		var err error
		port, err = strconv.Atoi(portEnv)
		if err != nil {
			log.Fatalf("invalid port: %v", err)
		}
	}

	server, err := server.Wire(&server.Config{
		Port: port,

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

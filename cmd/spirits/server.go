package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/ankeesler/spirits/pkg/server"
)

func runServer() error {
	var (
		port int
		help bool
	)

	flag.IntVar(&port, "port", 8675, "localhost port on which to listen")
	flag.Parse()

	if help {
		flag.Usage()
		return nil
	}

	address := fmt.Sprintf(":%d", port)
	log.Printf("starting server on %s", address)
	return http.ListenAndServe(address, server.New())
}

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ankeesler/spirits/api/internal/api"
)

func main() {
	var (
		webAssetsDir string
		help         bool
	)
	flag.StringVar(&webAssetsDir, "web-assets-dir", "public", "path to web assets directory")
	flag.BoolVar(&help, "help", false, "whether to print this help message")
	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	api := api.New()
	web := http.FileServer(http.Dir(webAssetsDir))
	mux := http.NewServeMux()
	mux.Handle("/", web)
	mux.Handle("/api/", http.StripPrefix("/api", api))

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "12345"
	}
	address := fmt.Sprintf(":%s", port)
	log.Printf("listening on address %s", address)
	if err := http.ListenAndServe(address, mux); err != nil {
		log.Printf("HTTP server stopped listening: %s\n", err.Error())
	}
}

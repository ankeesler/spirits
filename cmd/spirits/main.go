package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ankeesler/spirits/internal/api"
)

func main() {
	api := api.New()
	web := http.FileServer(http.Dir("public"))
	mux := http.NewServeMux()
	mux.Handle("/", web)
	mux.Handle("/api/", http.StripPrefix("/api", api))

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "12345"
	}
	address := fmt.Sprintf(":%s", port)
	log.Printf("listening on address %s", address)
	http.ListenAndServe(address, mux)
}

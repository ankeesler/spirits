package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ankeesler/spirits/internal/api"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "12345"
	}
	address := fmt.Sprintf(":%s", port)
	log.Printf("listening on address %s", address)
	http.ListenAndServe(address, api.New())
}

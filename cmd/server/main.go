package main

import (
	"log"
	"net/http"

	api "github.com/ankeesler/spirits/pkg/v0"
)

func main() {
	log.Print(http.ListenAndServe(":12345", api.New()))
}

package v0

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type server struct {
	rooms sync.Map // map from room_id string to *room{}
}

type room struct {
	Room
	manifests sync.Map // map from manifest_id string to *Manifest
}

// New returns an http.Handler that serves the spirits game server.
func New() http.Handler {
	var s server

	router := mux.NewRouter()
	router.HandleFunc("/", s.getRoot).Methods(http.MethodGet)

	router.HandleFunc("/rooms", s.createRoom).Methods(http.MethodPost)
	router.HandleFunc("/rooms", s.getRooms).Methods(http.MethodGet)
	router.HandleFunc("/rooms/{room_id}", s.getRoom).Methods(http.MethodGet)
	router.HandleFunc("/rooms/{room_id}/events", s.getRoomEvents).Methods(http.MethodGet)

	router.HandleFunc("/rooms/{room_id}/manifests", s.createManifest).Methods(http.MethodPost)
	router.HandleFunc("/rooms/{room_id}/manifests", s.getManifests).Methods(http.MethodGet)
	router.HandleFunc("/rooms/{room_id}/manifests/{manifest_id}", s.getManifest).Methods(http.MethodGet)
	return router
}

func errorBody(format string, args ...string) string {
	var e string
	if len(args) > 0 {
		e = fmt.Sprintf(format, args)
	} else {
		e = format
	}
	return fmt.Sprintf(`{"error":"%s"}`, e)
}

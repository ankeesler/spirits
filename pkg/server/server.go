// Package server provides an http.Handler that implements the spirits game server.
//
// The http.Handler in this package is serves a single WebSocket endpoint at /rooms/<room-id>.
//
// The API objects for interacting with this server can be found in pkg/api.
package server

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type server struct {
	rooms sync.Map
}

// New returns an http.Handler that serves the spirits game server.
func New() http.Handler {
	var s server

	router := mux.NewRouter()
	router.HandleFunc("/rooms/{id}", s.roomHandler).Methods(http.MethodGet)
	return router
}

var upgrader = websocket.Upgrader{
	// TODO: what to set in here...
}

func (s *server) roomHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("cannot upgrade websocket connection: %s", err.Error())
		return
	}
	defer ws.Close()
}

package v0

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func (s *server) createRoom(w http.ResponseWriter, r *http.Request) {
	var thisRoom room
	if err := json.NewDecoder(r.Body).Decode(&thisRoom.Room); err != nil {
		http.Error(w, errorBody(err.Error()), http.StatusInternalServerError)
		return
	}

	if thisRoom.Name == "" {
		http.Error(w, errorBody("room missing required 'name' field"), http.StatusBadRequest)
		return
	}

	if _, loaded := s.rooms.LoadOrStore(thisRoom.Name, &thisRoom); loaded {
		http.Error(w, errorBody("room with id %q already exists", thisRoom.Name), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(thisRoom.Room); err != nil {
		http.Error(w, errorBody(err.Error()), http.StatusInternalServerError)
		return
	}
}

func (s *server) getRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomID := vars["room_id"]
	thisRoom, ok := s.rooms.Load(roomID)
	if !ok {
		http.Error(w, errorBody("unknown room with id %q", roomID), http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(thisRoom.(*room).Room); err != nil {
		http.Error(w, errorBody(err.Error()), http.StatusInternalServerError)
		return
	}
}

func (s *server) getRooms(w http.ResponseWriter, r *http.Request) {
	rooms := []*Room{}
	s.rooms.Range(func(key interface{}, value interface{}) bool {
		rooms = append(rooms, &value.(*room).Room)
		return true
	})

	if err := json.NewEncoder(w).Encode(rooms); err != nil {
		http.Error(w, errorBody(err.Error()), http.StatusInternalServerError)
		return
	}
}

var upgrader = websocket.Upgrader{
	// TODO: what to set in here...
}

func (s *server) getRoomEvents(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("cannot upgrade websocket connection: %s", err.Error())
		return
	}
	defer ws.Close()
}

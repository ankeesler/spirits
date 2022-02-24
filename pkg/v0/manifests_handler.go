package v0

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *server) createManifest(w http.ResponseWriter, r *http.Request) {
	var manifest Manifest
	if err := json.NewDecoder(r.Body).Decode(&manifest); err != nil {
		http.Error(w, errorBody(err.Error()), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	roomID := vars["room_id"]
	thisRoom, ok := s.rooms.Load(roomID)
	if !ok {
		http.Error(w, errorBody("unknown room with id %q", roomID), http.StatusNotFound)
		return
	}

	manifestID, ok := manifest.Metadata["name"]
	if !ok {
		http.Error(w, errorBody("manifest missing required 'name' metadata"), http.StatusBadRequest)
		return
	}
	if _, loaded := thisRoom.(*room).manifests.LoadOrStore(manifestID, &manifest); loaded {
		http.Error(w, errorBody("manifest with id %q already exists", manifestID), http.StatusBadRequest)
		return
	}
	if err := json.NewEncoder(w).Encode(manifest); err != nil {
		http.Error(w, errorBody(err.Error()), http.StatusInternalServerError)
		return
	}
}

func (s *server) getManifest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomID := vars["room_id"]
	thisRoom, ok := s.rooms.Load(roomID)
	if !ok {
		http.Error(w, errorBody("unknown room with id %q", roomID), http.StatusNotFound)
		return
	}

	manifestID := vars["manifest_id"]
	thisManifest, ok := thisRoom.(*room).manifests.Load(manifestID)
	if !ok {
		http.Error(w, errorBody("unknown manifest with id %q", manifestID), http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(thisManifest); err != nil {
		http.Error(w, errorBody(err.Error()), http.StatusInternalServerError)
		return
	}
}

func (s *server) getManifests(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomID := vars["room_id"]
	thisRoom, ok := s.rooms.Load(roomID)
	if !ok {
		http.Error(w, errorBody("unknown room with id %q", roomID), http.StatusNotFound)
		return
	}

	manifests := []*Manifest{}
	thisRoom.(*room).manifests.Range(func(key interface{}, value interface{}) bool {
		manifests = append(manifests, value.(*Manifest))
		return true
	})

	if err := json.NewEncoder(w).Encode(manifests); err != nil {
		http.Error(w, errorBody(err.Error()), http.StatusInternalServerError)
		return
	}
}

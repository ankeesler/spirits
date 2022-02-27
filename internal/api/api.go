package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ankeesler/spirits/internal/battle"
	"github.com/ankeesler/spirits/internal/spirit"
	"github.com/ankeesler/spirits/internal/ui"
)

func New() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.URL.Path, "/battle"; got != want {
			log.Printf("url path %q != %q", got, want)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var spirits []*spirit.Spirit
		if err := json.NewDecoder(r.Body).Decode(&spirits); err != nil {
			http.Error(w, "cannot decode body: "+err.Error(), http.StatusBadRequest)
			return
		}

		if len(spirits) != 2 {
			http.Error(w, "must provide 2 spirits", http.StatusBadRequest)
			return
		}

		u := ui.New(w)
		battle.Run(spirits, u.OnSpirits)
	})
}

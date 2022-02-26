package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ankeesler/spirits/internal/battle"
	"github.com/ankeesler/spirits/internal/spirit"
	"github.com/ankeesler/spirits/internal/ui"
)

func New() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/battles" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var spirits []*spirit.Spirit
		if err := json.NewDecoder(r.Body).Decode(&spirits); err != nil {
			message := fmt.Sprintf(`{"error": "cannot decode body: %s"}`, err.Error())
			http.Error(w, message, http.StatusBadRequest)
			return
		}

		if len(spirits) != 2 {
			message := fmt.Sprintf(`{"error": "must provide 2 spirits"}`)
			http.Error(w, message, http.StatusBadRequest)
			return
		}

		u := ui.New(w)
		battle.Run(spirits, u.OnSpirits)
	})
}

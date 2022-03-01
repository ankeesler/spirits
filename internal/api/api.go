package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/ankeesler/spirits/internal/battle"
	"github.com/ankeesler/spirits/internal/spirit"
	"github.com/ankeesler/spirits/internal/spirit/generate"
	"github.com/ankeesler/spirits/internal/ui"
)

type Spirit struct {
	Name    string
	Health  int
	Power   int
	Agility int
	Armour  int
}

var handlers = map[string]http.HandlerFunc{
	"/battle": handleBattle,
	"/spirit": handleSpirit,
}

func New() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler, ok := handlers[r.URL.Path]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		handler(w, r)
	})
}

func handleBattle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var apiSpirits []*Spirit
	if err := json.NewDecoder(r.Body).Decode(&apiSpirits); err != nil {
		http.Error(w, "cannot decode body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(apiSpirits) != 2 {
		http.Error(w, "must provide 2 spirits", http.StatusBadRequest)
		return
	}

	internalSpirits := toInternalSpirits(apiSpirits)

	u := ui.New(w)
	battle.Run(internalSpirits, u.OnSpirits)
}

func toInternalSpirits(apiSpirits []*Spirit) []*spirit.Spirit {
	internalSpirits := make([]*spirit.Spirit, len(apiSpirits))
	for i := range apiSpirits {
		internalSpirits[i] = toModelSpirit(apiSpirits[i])
	}
	return internalSpirits
}

func toModelSpirit(apiSpirit *Spirit) *spirit.Spirit {
	return &spirit.Spirit{
		Name:    apiSpirit.Name,
		Health:  apiSpirit.Health,
		Power:   apiSpirit.Power,
		Agility: apiSpirit.Agility,
		Armour:  apiSpirit.Armour,
	}
}

func handleSpirit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var seed int64
	query := r.URL.Query()
	if query.Has("seed") {
		var err error
		seed, err = strconv.ParseInt(query.Get("seed"), 10, 64)
		if err != nil {
			http.Error(w, "invalid seed", http.StatusBadRequest)
			return
		}
	} else {
		seed = time.Now().Unix()
	}

	spirits := generate.Generate(int64(seed))
	if err := json.NewEncoder(w).Encode(spirits); err != nil {
		http.Error(w, "cannot encode response: "+err.Error(), http.StatusBadRequest)
		return
	}
}

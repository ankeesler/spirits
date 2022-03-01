package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ankeesler/spirits/internal/battle"
	"github.com/ankeesler/spirits/internal/spirit"
	"github.com/ankeesler/spirits/internal/spirit/action"
	"github.com/ankeesler/spirits/internal/spirit/generate"
	"github.com/ankeesler/spirits/internal/ui"
)

type Spirit struct {
	Name    string   `json:"name"`
	Health  int      `json:"health"`
	Power   int      `json:"power"`
	Agility int      `json:"agility"`
	Armour  int      `json:"armour"`
	Actions []string `json:"actions,omitempty"`
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

	internalSpirits, err := toInternalSpirits(apiSpirits)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u := ui.New(w)
	battle.Run(internalSpirits, u.OnSpirits)
}

func toInternalSpirits(apiSpirits []*Spirit) ([]*spirit.Spirit, error) {
	internalSpirits := make([]*spirit.Spirit, len(apiSpirits))
	var err error
	for i := range apiSpirits {
		internalSpirits[i], err = toInternalSpirit(apiSpirits[i])
		if err != nil {
			return nil, err
		}
	}
	return internalSpirits, nil
}

func toInternalSpirit(apiSpirit *Spirit) (*spirit.Spirit, error) {
	s := &spirit.Spirit{
		Name:    apiSpirit.Name,
		Health:  apiSpirit.Health,
		Power:   apiSpirit.Power,
		Agility: apiSpirit.Agility,
		Armour:  apiSpirit.Armour,
	}

	var err error
	s.Action, err = toInternalAction(apiSpirit.Actions)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func toInternalAction(ids []string) (spirit.Action, error) {
	if len(ids) == 0 {
		return action.Attack(), nil
	}

	if len(ids) > 1 {
		return nil, errors.New("must specify one action")
	}

	switch ids[0] {
	case "", "attack":
		return action.Attack(), nil
	case "bolster":
		return action.Bolster(), nil
	case "drain":
		return action.Drain(), nil
	case "charge":
		return action.Charge(), nil
	default:
		return nil, fmt.Errorf("unrecognized action: %q", ids[0])
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

	internalSpirits := generate.Generate(int64(seed))
	apiSpirits := fromInternalSpirits(internalSpirits)

	if err := json.NewEncoder(w).Encode(apiSpirits); err != nil {
		http.Error(w, "cannot encode response: "+err.Error(), http.StatusBadRequest)
		return
	}
}

func fromInternalSpirits(internalSpirits []*spirit.Spirit) []*Spirit {
	apiSpirits := make([]*Spirit, len(internalSpirits))
	for i := range apiSpirits {
		apiSpirits[i] = fromInternalSpirit(internalSpirits[i])
	}
	return apiSpirits
}

func fromInternalSpirit(internalSpirit *spirit.Spirit) *Spirit {
	return &Spirit{
		Name:    internalSpirit.Name,
		Health:  internalSpirit.Health,
		Power:   internalSpirit.Power,
		Agility: internalSpirit.Agility,
		Armour:  internalSpirit.Armour,

		// This worries me...we wire the action based on an API symbol (e.g., "attack"),
		// but rely on the underlying spirit.Action to give us the name back...that
		// seems asymetrical.
		Actions: []string{internalSpirit.Action.Name()},
	}
}

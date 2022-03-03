package api

import (
	"encoding/json"
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
	Name         string   `json:"name"`
	Health       int      `json:"health"`
	Power        int      `json:"power"`
	Agility      int      `json:"agility"`
	Armour       int      `json:"armour"`
	Actions      []string `json:"actions,omitempty"`
	Intelligence string   `json:"intelligence,omitempty"`
}

type actions struct {
	ids []string
	spirit.Action
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

	seed, ok := getSeed(r)
	if !ok {
		http.Error(w, "invalid seed", http.StatusBadRequest)
		return
	}

	internalSpirits, err := toInternalSpirits(apiSpirits, seed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u := ui.New(w)
	battle.Run(internalSpirits, u.OnSpirits)
}

func toInternalSpirits(apiSpirits []*Spirit, seed int64) ([]*spirit.Spirit, error) {
	internalSpirits := make([]*spirit.Spirit, len(apiSpirits))
	var err error
	for i := range apiSpirits {
		internalSpirits[i], err = toInternalSpirit(apiSpirits[i], seed)
		if err != nil {
			return nil, err
		}
	}
	return internalSpirits, nil
}

func toInternalSpirit(apiSpirit *Spirit, seed int64) (*spirit.Spirit, error) {
	s := &spirit.Spirit{
		Name:    apiSpirit.Name,
		Health:  apiSpirit.Health,
		Power:   apiSpirit.Power,
		Agility: apiSpirit.Agility,
		Armour:  apiSpirit.Armour,
	}

	var err error
	s.Action, err = toInternalAction(apiSpirit.Actions, apiSpirit.Intelligence, seed)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func toInternalAction(ids []string, intelligence string, seed int64) (spirit.Action, error) {
	var internalActions []spirit.Action
	if len(ids) == 0 {
		internalActions = []spirit.Action{action.Attack()}
	} else {
		for _, id := range ids {
			switch id {
			case "", "attack":
				internalActions = append(internalActions, action.Attack())
			case "bolster":
				internalActions = append(internalActions, action.Bolster())
			case "drain":
				internalActions = append(internalActions, action.Drain())
			case "charge":
				internalActions = append(internalActions, action.Charge())
			default:
				return nil, fmt.Errorf("unrecognized action: %q", ids[0])
			}
		}
	}

	var internalAction spirit.Action
	switch intelligence {
	case "", "roundrobin":
		internalAction = action.RoundRobin(internalActions)
	case "random":
		internalAction = action.Random(seed, internalActions)
	default:
		return nil, fmt.Errorf("unrecognized intelligence: %q", intelligence)
	}

	return &actions{ids: ids, Action: internalAction}, nil
}

func handleSpirit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	seed, ok := getSeed(r)
	if !ok {
		http.Error(w, "invalid seed", http.StatusBadRequest)
		return
	}

	wellKnownActions := []spirit.Action{
		&actions{
			ids:    []string{"attack"},
			Action: action.Attack(),
		},
		&actions{
			ids:    []string{"bolster"},
			Action: action.Bolster(),
		},
		&actions{
			ids:    []string{"drain"},
			Action: action.Drain(),
		},
		&actions{
			ids:    []string{"charge"},
			Action: action.Charge(),
		},
	}
	internalSpirits := generate.Generate(int64(seed), wellKnownActions, func(generatedActions []spirit.Action) spirit.Action {
		var ids []string
		for _, generatedAction := range generatedActions {
			ids = append(ids, generatedAction.(*actions).ids...)
		}
		return &actions{
			ids:    ids,
			Action: action.RoundRobin(generatedActions),
		}
	})
	apiSpirits, err := fromInternalSpirits(internalSpirits)
	if err != nil {
		http.Error(w, "cannot process spirits: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(apiSpirits); err != nil {
		http.Error(w, "cannot encode response: "+err.Error(), http.StatusBadRequest)
		return
	}
}

func fromInternalSpirits(internalSpirits []*spirit.Spirit) ([]*Spirit, error) {
	apiSpirits := make([]*Spirit, len(internalSpirits))
	for i := range apiSpirits {
		var err error
		apiSpirits[i], err = fromInternalSpirit(internalSpirits[i])
		if err != nil {
			return nil, fmt.Errorf("could not convert from internal spirit %s: %w", internalSpirits[i].Name, err)
		}
	}
	return apiSpirits, nil
}

func fromInternalSpirit(internalSpirit *spirit.Spirit) (*Spirit, error) {
	apiActions, ok := internalSpirit.Action.(*actions)
	if !ok {
		return nil, fmt.Errorf("invalid internal spirit action type: %T", internalSpirit.Action)
	}

	return &Spirit{
		Name:    internalSpirit.Name,
		Health:  internalSpirit.Health,
		Power:   internalSpirit.Power,
		Agility: internalSpirit.Agility,
		Armour:  internalSpirit.Armour,
		Actions: apiActions.ids,
	}, nil
}

func getSeed(r *http.Request) (int64, bool) {
	var seed int64
	query := r.URL.Query()
	if query.Has("seed") {
		var err error
		seed, err = strconv.ParseInt(query.Get("seed"), 10, 64)
		if err != nil {
			return 0, false
		}
	} else {
		seed = time.Now().Unix()
	}
	return seed, true
}

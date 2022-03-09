package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ankeesler/spirits/api/internal/spirit"
	"github.com/ankeesler/spirits/api/internal/spirit/action"
)

type Spirit struct {
	Name         string   `json:"name"`
	Health       int      `json:"health"`
	Power        int      `json:"power"`
	Agility      int      `json:"agility"`
	Armor        int      `json:"armor"`
	Actions      []string `json:"actions,omitempty"`
	Intelligence string   `json:"intelligence,omitempty"`
}

type actions struct {
	ids []string
	spirit.Action
}

func New() http.Handler {
	return http.HandlerFunc(serveWebsocket)
}

func toInternalSpirits(
	apiSpirits []*Spirit,
	seed int64,
	humanActionFunc func(ctx context.Context, s *Spirit) (spirit.Action, error),
) ([]*spirit.Spirit, error) {
	internalSpirits := make([]*spirit.Spirit, len(apiSpirits))
	var err error
	for i := range apiSpirits {
		internalSpirits[i], err = toInternalSpirit(apiSpirits[i], seed, humanActionFunc)
		if err != nil {
			return nil, err
		}
	}
	return internalSpirits, nil
}

func toInternalSpirit(apiSpirit *Spirit, seed int64, humanActionFunc func(ctx context.Context, s *Spirit) (spirit.Action, error)) (*spirit.Spirit, error) {
	s := &spirit.Spirit{
		Name:    apiSpirit.Name,
		Health:  apiSpirit.Health,
		Power:   apiSpirit.Power,
		Agility: apiSpirit.Agility,
		Armor:   apiSpirit.Armor,
	}

	var lazyActionFunc func(ctx context.Context) (spirit.Action, error)
	if humanActionFunc != nil {
		lazyActionFunc = func(ctx context.Context) (spirit.Action, error) {
			return humanActionFunc(ctx, apiSpirit)
		}
	}
	var err error
	s.Action, err = toInternalAction(apiSpirit.Actions, apiSpirit.Intelligence, seed, lazyActionFunc)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func toInternalAction(
	ids []string,
	intelligence string,
	seed int64,
	lazyActionFunc func(ctx context.Context) (spirit.Action, error),
) (spirit.Action, error) {
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
	case "human":
		if lazyActionFunc == nil {
			return nil, fmt.Errorf("unsupported intelligence value (hint: you must use websocket API): %q", intelligence)
		}
		internalAction = action.Lazy(lazyActionFunc)
	default:
		return nil, fmt.Errorf("unrecognized intelligence: %q", intelligence)
	}

	return &actions{ids: ids, Action: internalAction}, nil
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
		Armor:   internalSpirit.Armor,
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

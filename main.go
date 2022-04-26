package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"path"
	"sync"

	"github.com/gorilla/websocket"

	battlepkg "github.com/ankeesler/spirits/internal/battle"
	spiritpkg "github.com/ankeesler/spirits/internal/spirit"
	actionpkg "github.com/ankeesler/spirits/internal/spirit/action"
)

func main() {
	mux := http.NewServeMux()

	spiritsManager := newSpiritsManager()
	spiritsStore := newMemoryStore(&storeHooks[spirit]{
		create: spiritsManager.upsert,
		update: spiritsManager.upsert,
	})

	battlesManager := newBattlesManager(spiritsStore, spiritsManager)
	battlesStore := newMemoryStore(&storeHooks[battle]{
		create: battlesManager.create,
		delete: battlesManager.delete,
	})

	spiritsHandler := newHandler("spirits", spiritsStore, handlerCapsAll)
	battlesHandler := newHandler("battles", battlesStore, handlerCapsAll&^handlerCapsUpdate)

	spiritsHandler.install(mux)
	battlesHandler.install(mux)

	const address = ":12345"
	log.Printf("listening on %s", address)
	panic(http.ListenAndServe(address, mux))
}

type Nameable interface {
	GetName() string
}

// -------------------------------------------------------------------------------------------------
// store

type storeHooks[T Nameable] struct {
	create func(context.Context, *T) error
	update func(context.Context, *T) error
	delete func(context.Context, string) error
}

type store[T Nameable] struct {
	l     sync.Mutex
	m     map[string]T
	hooks *storeHooks[T]

	watchC   chan T
	watchers map[string][]chan<- T
}

func newMemoryStore[T Nameable](hooks *storeHooks[T]) *store[T] {
	return &store[T]{
		m:     make(map[string]T),
		hooks: hooks,

		watchers: make(map[string][]chan<- T),
	}
}

func (s *store[T]) Start(ctx context.Context) error {
	s.l.Lock()
	defer s.l.Unlock()

	if s.watchC != nil {
		return errors.New("already started")
	}

	s.watchC = make(chan T)
	go s.watch(ctx)

	return nil
}

func (s *store[T]) Create(ctx context.Context, t T) (T, error) {
	s.l.Lock()
	defer s.l.Unlock()

	if s.hooks != nil && s.hooks.create != nil {
		if err := s.hooks.create(ctx, &t); err != nil {
			return t, &storeErrInvalidObject{err}
		}
	}

	_, ok := s.m[t.GetName()]
	if ok {
		return t, &storeErrAlreadyExists{}
	}
	s.m[t.GetName()] = t

	return t, nil
}

func (s *store[T]) Update(ctx context.Context, t T) (T, error) {
	s.l.Lock()
	defer s.l.Unlock()

	if s.hooks != nil && s.hooks.create != nil {
		if err := s.hooks.update(ctx, &t); err != nil {
			return t, &storeErrInvalidObject{err}
		}
	}

	s.m[t.GetName()] = t

	return t, nil
}

func (s *store[T]) List(ctx context.Context) ([]T, error) {
	s.l.Lock()
	defer s.l.Unlock()

	var ts []T
	for _, v := range s.m {
		ts = append(ts, v)
	}
	return ts, nil
}

func (s *store[T]) Get(ctx context.Context, name string) (T, error) {
	s.l.Lock()
	defer s.l.Unlock()

	t, ok := s.m[name]
	if !ok {
		return t, &storeErrNotFound{}
	}

	return t, nil
}

func (s *store[T]) Watch(ctx context.Context, name string, ts chan<- T) error {
	watchers := s.watchers[name]
	watchers = append(watchers, ts)
	s.watchers[name] = watchers // In case watchers was nil
	return nil
}

func (s *store[T]) watch(ctx context.Context) {
	for {
		var t T

		select {
		case <-ctx.Done():
			close(s.watchC)
			s.closeWatchers()
			return
		case t = <-s.watchC:
		}

		for _, watcher := range s.watchers[t.GetName()] {
			select {
			case watcher <- t:
			default:
				delete(s.watchers, t.GetName())
			}
		}
	}
}

func (s *store[T]) closeWatchers() {
	for name, watchers := range s.watchers {
		for _, watcher := range watchers {
			close(watcher)
		}
		delete(s.watchers, name)
	}
}

func (s *store[T]) Delete(ctx context.Context, name string) (T, error) {
	s.l.Lock()
	defer s.l.Unlock()

	if s.hooks != nil && s.hooks.create != nil {
		if err := s.hooks.delete(ctx, name); err != nil {
			var t T
			return t, &storeErrInvalidObject{err}
		}
	}

	t, ok := s.m[name]
	if !ok {
		return t, &storeErrNotFound{}
	}

	delete(s.m, name)
	return t, nil
}

type storeErrInvalidObject struct {
	err error
}

func (e *storeErrInvalidObject) Error() string {
	return "invalid object: " + e.err.Error()
}

type storeErrAlreadyExists struct {
}

func (e *storeErrAlreadyExists) Error() string {
	return "already exists"
}

type storeErrNotFound struct {
}

func (e *storeErrNotFound) Error() string {
	return "not found"
}

// -------------------------------------------------------------------------------------------------
// handler

type handlerCaps int

const (
	handlerCapsCreate handlerCaps = (1 << iota)
	handlerCapsUpdate
	handlerCapsList
	handlerCapsGet
	handlerCapsWatch
	handlerCapsDelete

	handlerCapsAll = handlerCapsCreate | handlerCapsUpdate | handlerCapsList | handlerCapsGet | handlerCapsWatch | handlerCapsDelete
)

type handler[T Nameable] struct {
	name  string
	store *store[T]
	caps  handlerCaps

	upgrader *websocket.Upgrader
}

func newHandler[T Nameable](name string, store *store[T], caps handlerCaps) *handler[T] {
	return &handler[T]{
		name:  name,
		store: store,
		caps:  caps,

		upgrader: &websocket.Upgrader{},
	}
}

type serveMux interface {
	Handle(string, http.Handler)
}

func (h *handler[T]) install(mux serveMux) {
	mux.Handle("/"+h.name, h.handle(h.handleTs))
	mux.Handle("/"+h.name+"/", h.handle(h.handleT))
}

func (h *handler[T]) handle(handleFunc func(http.ResponseWriter, *http.Request, T) (any, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.log(fmt.Sprintf("%s %s", r.Method, r.URL))

		var reqBody T

		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		h.log(fmt.Sprintf("req body: %#v", reqBody))

		rspBody, err := handleFunc(w, r, reqBody)
		if err != nil {
			if handlerErr, ok := err.(*handlerErr); ok {
				http.Error(w, handlerErr.Error(), handlerErr.Code())
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		h.log(fmt.Sprintf("rsp body: %#v", rspBody))

		if rspBody != nil {
			if err := json.NewEncoder(w).Encode(rspBody); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

func (h *handler[T]) handleTs(w http.ResponseWriter, r *http.Request, body T) (any, error) {
	switch r.Method {
	case http.MethodPost:
		t, err := h.store.Create(r.Context(), body)
		if errors.Is(err, &storeErrInvalidObject{}) {
			return nil, &handlerErr{err: err, code: http.StatusBadRequest}
		}
		return t, nil
	case http.MethodGet:
		return h.store.List(r.Context())
	default:
		return nil, &handlerErr{code: http.StatusMethodNotAllowed}
	}
}

func (h *handler[T]) handleT(w http.ResponseWriter, r *http.Request, body T) (any, error) {
	name := path.Base(r.URL.Path)
	switch r.Method {
	case http.MethodPut:
		t, err := h.store.Update(r.Context(), body)
		if errors.Is(err, &storeErrInvalidObject{}) {
			return nil, &handlerErr{err: err, code: http.StatusBadRequest}
		}
		return t, nil
	case http.MethodGet:
		if websocket.IsWebSocketUpgrade(r) {
			return nil, h.handleWebsocket(w, r, name)
		}

		t, err := h.store.Get(r.Context(), name)
		if errors.Is(err, &storeErrNotFound{}) {
			return nil, &handlerErr{code: http.StatusNotFound}
		}
		return t, nil
	case http.MethodDelete:
		t, err := h.store.Delete(r.Context(), name)
		if errors.Is(err, &storeErrNotFound{}) {
			return nil, &handlerErr{code: http.StatusNotFound}
		}
		return t, nil
	default:
		return body, &handlerErr{code: http.StatusMethodNotAllowed}
	}
}

func (h *handler[T]) handleWebsocket(w http.ResponseWriter, r *http.Request, name string) error {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	watchC := make(chan T)
	defer close(watchC)

	if err := h.store.Watch(r.Context(), name, watchC); err != nil {
		return fmt.Errorf("could not start watch: %w", err)
	}

	for t := range watchC {
		if err := conn.WriteJSON(t); err != nil {
			return fmt.Errorf("could not write to websocket: %w", err)
		}
	}

	return errors.New("watch closed")
}

func (h *handler[T]) log(msg string) {
	log.Printf("%s handler: %s", h.name, msg)
}

type handlerErr struct {
	err  error
	code int
}

func (e *handlerErr) Error() string {
	msg := http.StatusText(e.code)
	if e.err != nil {
		msg += ": " + e.err.Error()
	}
	return msg
}

func (e *handlerErr) Code() int {
	return e.code
}

// -------------------------------------------------------------------------------------------------
// spirit

type spirit struct {
	Name         string      `json:"name"`
	Stats        spiritStats `json:"stats"`
	Actions      []string    `json:"actions"`
	Intelligence string      `json:"intelligence"`
}

type spiritStats struct {
	Health  int `json:"health"`
	Power   int `json:"power"`
	Armor   int `json:"armor"`
	Agility int `json:"agility"`
}

func (s spirit) GetName() string { return s.Name }

type spiritsManager struct {
	l sync.Mutex
	m map[string]*spiritpkg.Spirit
}

func newSpiritsManager() *spiritsManager {
	return &spiritsManager{
		m: make(map[string]*spiritpkg.Spirit),
	}
}

func (m *spiritsManager) upsert(ctx context.Context, s *spirit) error {
	m.l.Lock()
	defer m.l.Unlock()

	internalSpirit, err := toInternalSpirit(s, nil, func(ctx context.Context, s *spirit) (spiritpkg.Action, error) {
		// TODO: human intelligence
		return actionpkg.Attack(), nil
	})
	if err != nil {
		return err
	}

	m.m[s.Name] = internalSpirit

	return nil
}

func (m *spiritsManager) get(ctx context.Context, name string) *spiritpkg.Spirit {
	return m.m[name]
}

type actions struct {
	ids          []string
	intelligence string
	spiritpkg.Action
}

func toInternalSpirit(apiSpirit *spirit, r *rand.Rand, humanActionFunc func(ctx context.Context, s *spirit) (spiritpkg.Action, error)) (*spiritpkg.Spirit, error) {
	s := &spiritpkg.Spirit{
		Name:    apiSpirit.Name,
		Health:  apiSpirit.Stats.Health,
		Power:   apiSpirit.Stats.Power,
		Agility: apiSpirit.Stats.Agility,
		Armor:   apiSpirit.Stats.Armor,
	}

	var lazyActionFunc func(ctx context.Context) (spiritpkg.Action, error)
	if humanActionFunc != nil {
		lazyActionFunc = func(ctx context.Context) (spiritpkg.Action, error) {
			return humanActionFunc(ctx, apiSpirit)
		}
	}
	var err error
	s.Action, err = toInternalAction(apiSpirit.Actions, apiSpirit.Intelligence, r, lazyActionFunc)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func toInternalAction(
	ids []string,
	intelligence string,
	r *rand.Rand,
	lazyActionFunc func(ctx context.Context) (spiritpkg.Action, error),
) (spiritpkg.Action, error) {
	var internalActions []spiritpkg.Action
	if len(ids) == 0 {
		internalActions = []spiritpkg.Action{actionpkg.Attack()}
	} else {
		for _, id := range ids {
			switch id {
			case "", "attack":
				internalActions = append(internalActions, actionpkg.Attack())
			case "bolster":
				internalActions = append(internalActions, actionpkg.Bolster())
			case "drain":
				internalActions = append(internalActions, actionpkg.Drain())
			case "charge":
				internalActions = append(internalActions, actionpkg.Charge())
			default:
				return nil, fmt.Errorf("unrecognized action: %q", ids[0])
			}
		}
	}

	var internalAction spiritpkg.Action
	switch intelligence {
	case "", "roundrobin":
		internalAction = actionpkg.RoundRobin(internalActions)
	case "random":
		internalAction = actionpkg.Random(r, internalActions)
	case "human":
		if lazyActionFunc == nil {
			return nil, fmt.Errorf("unsupported intelligence value (hint: you must use websocket API): %q", intelligence)
		}
		internalAction = actionpkg.Lazy(lazyActionFunc)
	default:
		return nil, fmt.Errorf("unrecognized intelligence: %q", intelligence)
	}

	return &actions{ids: ids, intelligence: intelligence, Action: internalAction}, nil
}

func fromInternalSpirit(internalSpirit *spiritpkg.Spirit) (*spirit, error) {
	apiActions, ok := internalSpirit.Action.(*actions)
	if !ok {
		return nil, fmt.Errorf("invalid internal spirit action type: %T", internalSpirit.Action)
	}

	return &spirit{
		Name: internalSpirit.Name,
		Stats: spiritStats{Health: internalSpirit.Health,
			Power:   internalSpirit.Power,
			Agility: internalSpirit.Agility,
			Armor:   internalSpirit.Armor,
		},
		Actions:      apiActions.ids,
		Intelligence: apiActions.intelligence,
	}, nil
}

// -------------------------------------------------------------------------------------------------
// battle

type battle struct {
	Name    string   `json:"name"`
	Spirits []string `json:"spirits"`

	// The below fields are set by the battle handler

	State           string    `json:"state"`
	Message         string    `json:"message"`
	InBattleSpirits []*spirit `json:"inBattleSpirits,omitempty"`

	// The below fields are private to this battle

	cancel context.CancelFunc
}

func (b battle) GetName() string { return b.Name }

type spiritsStore interface {
	Create(context.Context, spirit) (spirit, error)
	Update(context.Context, spirit) (spirit, error)
}

type spiritsCache interface {
	get(context.Context, string) *spiritpkg.Spirit
}

type battlesManager struct {
	l            sync.Mutex
	m            map[string]*battle
	spiritsStore spiritsStore
	spiritsCache spiritsCache
}

func newBattlesManager(spiritsStore spiritsStore, spiritsCache spiritsCache) *battlesManager {
	return &battlesManager{
		m:            make(map[string]*battle),
		spiritsStore: spiritsStore,
		spiritsCache: spiritsCache,
	}
}

func (m *battlesManager) create(ctx context.Context, battle *battle) error {
	// Get internal spirits from cache
	var internalSpirits []*spiritpkg.Spirit
	for _, spiritName := range battle.Spirits {
		internalSpirit := m.spiritsCache.get(ctx, spiritName)
		if internalSpirit == nil {
			return fmt.Errorf("unknown spirit: %q", spiritName)
		}
		internalSpirits = append(internalSpirits, internalSpirit)
	}

	go m.run(battle, internalSpirits)

	return nil
}

func (m *battlesManager) delete(ctx context.Context, name string) error {
	m.l.Lock()
	defer m.l.Unlock()

	battle, ok := m.m[name]
	if ok {
		battle.cancel()
	}

	return nil
}

func (m *battlesManager) run(battle *battle, internalSpirits []*spiritpkg.Spirit) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var inBattleInternalSpirits []*spiritpkg.Spirit
	for _, internalSpirit := range internalSpirits {
		// Create in-battle spirit
		inBattleSpirit, err := fromInternalSpirit(internalSpirit)
		if err != nil {
			panic(err)
		}
		if _, err := m.spiritsStore.Create(ctx, *inBattleSpirit); err != nil {
			panic(err)
		}

		// Map to in-battle spirit name
		inBattleSpirit.Name = fmt.Sprintf("%s-%s", battle.Name, internalSpirit.Name)

		// Create in-battle spirit
		if _, err := m.spiritsStore.Create(ctx, *inBattleSpirit); err != nil {
			panic(err)
		}

		// Get in-battle internal spirit from cache
		inBattleInternalSpirit := m.spiritsCache.get(ctx, inBattleSpirit.Name)
		if inBattleInternalSpirit == nil {
			panic(fmt.Sprintf("cannot find in-battle internal spirit %q in cache", inBattleInternalSpirit.Name))
		}
		inBattleInternalSpirits = append(inBattleInternalSpirits, inBattleInternalSpirit)
	}

	battlepkg.Run(ctx, inBattleInternalSpirits, func(inBattleInternalSpirits []*spiritpkg.Spirit, err error) {
		if err != nil {
			panic(err)
		}

		// Update each in-battle spirit
		for _, inBattleInternalSpirit := range inBattleInternalSpirits {
			inBattleSpirit, err := fromInternalSpirit(inBattleInternalSpirit)
			if err != nil {
				panic(err)
			}
			if _, err := m.spiritsStore.Update(ctx, *inBattleSpirit); err != nil {
				panic(err)
			}
		}
	})
}

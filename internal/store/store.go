package store

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/ankeesler/spirits/internal/log"
)

type Store[T any] struct {
	l sync.Mutex
	m map[string]T

	nameFunc func(*T) string

	watchC   chan *T
	watchers map[string][]chan<- *T
}

func New[T any](nameFunc func(*T) string) *Store[T] {
	return &Store[T]{
		m: make(map[string]T),

		nameFunc: nameFunc,

		watchers: make(map[string][]chan<- *T),
	}
}

func (s *Store[T]) Start(ctx context.Context) error {
	log.Debug("store: start: begin")
	defer log.Debug("store: start: end")

	s.l.Lock()
	defer s.l.Unlock()

	if s.watchC != nil {
		return errors.New("already started")
	}

	s.watchC = make(chan *T)
	go s.watch(ctx)

	return nil
}

func (s *Store[T]) Create(ctx context.Context, t *T) (*T, error) {
	log.Debug(fmt.Sprintf("store: create %#v: begin", t))
	defer log.Debug(fmt.Sprintf("store: create %#v: end", t))

	s.l.Lock()
	defer s.l.Unlock()

	_, ok := s.m[s.nameFunc(t)]
	if ok {
		return t, &ErrAlreadyExists{}
	}
	s.m[s.nameFunc(t)] = *t

	return t, nil
}

func (s *Store[T]) Update(ctx context.Context, t *T) (*T, error) {
	log.Debug(fmt.Sprintf("store: update %#v: begin", t))
	defer log.Debug(fmt.Sprintf("store: update %#v: end", t))

	s.l.Lock()
	defer s.l.Unlock()

	if _, ok := s.m[s.nameFunc(t)]; !ok {
		return nil, &ErrNotFound{}
	}

	s.m[s.nameFunc(t)] = *t

	return t, nil
}

func (s *Store[T]) List(ctx context.Context) ([]*T, error) {
	log.Debug("store: list: begin")
	defer log.Debug("store: list: end")

	s.l.Lock()
	defer s.l.Unlock()

	var ts []*T
	for _, v := range s.m {
		ts = append(ts, &v)
	}
	return ts, nil
}

func (s *Store[T]) Get(ctx context.Context, name string) (*T, error) {
	log.Debug(fmt.Sprintf("store: get %q: begin", name))
	defer log.Debug(fmt.Sprintf("store: update %q: end", name))

	s.l.Lock()
	defer s.l.Unlock()

	t, ok := s.m[name]
	if !ok {
		return nil, &ErrNotFound{}
	}

	return &t, nil
}

func (s *Store[T]) Watch(ctx context.Context, name string, ts chan<- *T) error {
	log.Debug(fmt.Sprintf("store: watch %q: begin", name))
	defer log.Debug(fmt.Sprintf("store: watch %q: end", name))

	s.l.Lock()
	defer s.l.Unlock()

	watchers := s.watchers[name]
	watchers = append(watchers, ts)
	s.watchers[name] = watchers // In case watchers was nil
	return nil
}

func (s *Store[T]) Delete(ctx context.Context, name string) (*T, error) {
	log.Debug(fmt.Sprintf("store: delete %q: begin", name))
	defer log.Debug(fmt.Sprintf("store: delete %q: end", name))

	s.l.Lock()
	defer s.l.Unlock()

	t, ok := s.m[name]
	if !ok {
		return nil, &ErrNotFound{}
	}

	delete(s.m, name)
	return &t, nil
}

func (s *Store[T]) watch(ctx context.Context) {
	for {
		var t *T

		select {
		case <-ctx.Done():
			close(s.watchC)
			s.closeWatchers()
			return
		case t = <-s.watchC:
		}

		for _, watcher := range s.watchers[s.nameFunc(t)] {
			select {
			case watcher <- t:
			default:
				delete(s.watchers, s.nameFunc(t))
			}
		}
	}
}

func (s *Store[T]) closeWatchers() {
	for name, watchers := range s.watchers {
		for _, watcher := range watchers {
			close(watcher)
		}
		delete(s.watchers, name)
	}
}

type ErrAlreadyExists struct {
}

func (e *ErrAlreadyExists) Error() string {
	return "already exists"
}

type ErrNotFound struct {
}

func (e *ErrNotFound) Error() string {
	return "not found"
}

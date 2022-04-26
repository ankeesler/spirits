package service

import (
	"context"
	"errors"
	"sync"
)

type Nameable interface {
	GetName() string
}

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

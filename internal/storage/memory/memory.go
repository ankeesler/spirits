package memory

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Meta interface {
	ID() string
	SetID(string)

	CreatedTime() time.Time
	SetCreatedTime(time.Time)

	UpdatedTime() time.Time
	SetUpdatedTime(time.Time)
}

type watchContext[T Meta] struct {
	id *string
	c  chan<- T
}

type Storage[T Meta] struct {
	r *rand.Rand

	data    map[string]T
	watches sync.Map

	lock sync.Mutex
}

func New[T Meta](r *rand.Rand) *Storage[T] {
	return &Storage[T]{
		r: r,

		data: make(map[string]T),
	}
}

func (s *Storage[T]) Create(
	ctx context.Context,
	t T,
) (T, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	log.Printf("creating %+v", t)

	id := fmt.Sprintf("%x", s.r.Uint64())
	if _, ok := s.data[id]; ok {
		return t, status.Error(codes.AlreadyExists, "already exists")
	}

	t.SetID(id)
	t.SetCreatedTime(time.Now())
	t.SetUpdatedTime(t.CreatedTime())

	s.data[t.ID()] = t
	go s.notifyWatch(ctx, t)

	return t, nil
}

func (s *Storage[T]) Get(
	ctx context.Context,
	id string,
) (T, error) {
	log.Printf("waiting to get %s", id)

	s.lock.Lock()
	defer s.lock.Unlock()

	log.Printf("getting %s", id)

	t, ok := s.data[id]
	if !ok {
		return t, status.Error(codes.NotFound, "not found")
	}

	return t, nil
}

func (s *Storage[T]) Watch(
	ctx context.Context,
	id *string,
) (<-chan T, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	var t T
	log.Printf("opening watch for %T", t)

	c := make(chan T, 1)
	watchID := fmt.Sprintf("%x", s.r.Uint64())
	s.watches.Store(watchID, &watchContext[T]{
		id: id,
		c:  c,
	})

	if id != nil {
		if t, ok := s.data[*id]; ok {
			c <- t
		}
	}

	go func() {
		<-ctx.Done()
		s.watches.Delete(watchID)
		close(c)
	}()

	return c, nil
}

func (s *Storage[T]) List(
	ctx context.Context,
) ([]T, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	log.Print("listing")

	var ts []T
	for _, t := range s.data {
		ts = append(ts, t)
	}

	return ts, nil
}

func (s *Storage[T]) Update(
	ctx context.Context,
	t T,
) (T, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	log.Printf("updating %#v", t)

	id := t.ID()
	if _, ok := s.data[id]; !ok {
		return t, status.Error(codes.NotFound, "not found")
	}

	t.SetUpdatedTime(time.Now())

	s.data[id] = t
	go s.notifyWatch(ctx, t)

	return t, nil
}

func (s *Storage[T]) Delete(
	ctx context.Context,
	id string,
) (T, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	log.Printf("deleting %s", id)

	t, ok := s.data[id]
	if !ok {
		return t, status.Error(codes.NotFound, "not found")
	}

	delete(s.data, id)

	return t, nil
}

func (s *Storage[T]) notifyWatch(
	ctx context.Context,
	t T,
) {
	log.Printf("kicking watch for %#v", t)

	s.watches.Range(func(key, val any) bool {
		watchCtx := val.(*watchContext[T])
		if watchCtx.id == nil || *watchCtx.id == t.ID() {
			watchCtx.c <- t
		}
		return true
	})
}

package memory

import (
	"context"
	"fmt"
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

type Storage[T Meta] struct {
	r *rand.Rand

	data    map[string]T
	watches []chan<- T

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
	s.lock.Lock()
	defer s.lock.Unlock()

	t, ok := s.data[id]
	if !ok {
		return t, status.Error(codes.NotFound, "not found")
	}

	return t, nil
}

func (s *Storage[T]) Watch(
	ctx context.Context,
	c chan<- T,
) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.watches = append(s.watches, c)

	return nil
}

func (s *Storage[T]) List(
	ctx context.Context,
) ([]T, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

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
	for _, watch := range s.watches {
		watch <- t
	}
}

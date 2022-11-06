package memory

import (
	"context"
	"fmt"
	"math/rand"
	"sync"

	"github.com/ankeesler/spirits/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Meta interface {
	GetMeta() *api.Meta
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
	validateFunc func(T) error,
) (T, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	id := fmt.Sprintf("%x", s.r.Uint64())
	if _, ok := s.data[id]; ok {
		return t, status.Error(codes.AlreadyExists, "already exists")
	}

	t.GetMeta().Id = id
	t.GetMeta().CreatedTime = timestamppb.Now()
	t.GetMeta().UpdatedTime = t.GetMeta().CreatedTime

	if err := validateFunc(t); err != nil {
		return t, status.Error(codes.InvalidArgument, err.Error())
	}

	s.data[t.GetMeta().GetId()] = t
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
	validateFunc func(T) error,
) (T, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	id := t.GetMeta().GetId()
	if _, ok := s.data[id]; !ok {
		return t, status.Error(codes.NotFound, "not found")
	}

	t.GetMeta().UpdatedTime = timestamppb.Now()

	if err := validateFunc(t); err != nil {
		return t, status.Error(codes.InvalidArgument, err.Error())
	}

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

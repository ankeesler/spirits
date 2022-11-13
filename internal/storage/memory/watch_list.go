package memory

import (
	"container/list"
	"context"
	"log"
	"path/filepath"
	"runtime"
	"sync"
)

type watchContext[T Meta] struct {
	id *string
	c  chan<- T
}

type watchList[T Meta] struct {
	l    *list.List
	lock sync.Mutex
}

func newWatchList[T Meta]() *watchList[T] {
	return &watchList[T]{
		l: list.New(),
	}
}

func (l *watchList[T]) add(ctx context.Context, id *string, c chan<- T) {
	l.lock.Lock()
	defer l.lock.Unlock()

	_, callerFile, callerLine, ok := runtime.Caller(2)
	if !ok {
		callerFile = "???"
		callerLine = 0
	} else {
		callerFile = filepath.Base(callerFile)
	}

	var t T
	log.Printf("opening watch for %T and id %+v from %s:%d",
		t, id, callerFile, callerLine)

	e := l.l.PushBack(&watchContext[T]{id: id, c: c})
	go func() {
		<-ctx.Done()
		log.Printf("closing watch for %T and id %+v from %s:%d",
			t, id, callerFile, callerLine)

		func() {
			l.lock.Lock()
			defer l.lock.Unlock()

			l.l.Remove(e)
		}()

		close(c)
	}()
}

func (l *watchList[T]) notify(t T) {
	log.Printf("kicking watch for %#v", t)

	l.lock.Lock()
	defer l.lock.Unlock()

	for e := l.l.Front(); e != nil; e = e.Next() {
		watchCtx := e.Value.(*watchContext[T])
		if watchCtx.id == nil || *watchCtx.id == t.ID() {
			log.Printf("about to watch from id %+v", t.ID())
			watchCtx.c <- t
			log.Printf("watched from id %+v", t.ID())
		}
	}
}

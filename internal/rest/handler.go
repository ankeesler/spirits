package rest

import "net/http"

type Handler interface {
	Handle(http.ResponseWriter, *http.Request) error
}

type HandlerFunc func(http.ResponseWriter, *http.Request) error

var _ Handler = HandlerFunc(func(w http.ResponseWriter, r *http.Request) error { return nil })

func (h HandlerFunc) Handle(w http.ResponseWriter, r *http.Request) error {
	return h(w, r)
}

func ServeHTTP(h Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := h.Handle(w, r); err != nil {
			Error(w, err)
			return
		}
	})
}

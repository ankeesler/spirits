package server

import "net/http"

func Run() error {
	return http.ListenAndServe(":80", nil)
}

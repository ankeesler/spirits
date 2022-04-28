package rest

import (
	"errors"
	"fmt"
	"net/http"
)

type Err struct {
	Code int
	Err  error
}

var _ error = &Err{}

func (e *Err) Error() string {
	msg := http.StatusText(e.Code)
	if len(msg) > 0 {
		msg = fmt.Sprintf("Unrecognized HTTP Status (%d)", e.Code)
	}
	if e.Err != nil {
		msg += ":" + e.Err.Error()
	}
	return msg
}

func Error(w http.ResponseWriter, err error) {
	var restErr Err
	if errors.As(err, &restErr) {
		http.Error(w, restErr.Error(), restErr.Code)
		return
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

package v0

import (
	"fmt"
	"net/http"
)

func (s *server) getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `"hello great spirits"`)
}

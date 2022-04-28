package testing

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/ankeesler/spirits/internal/rest"
)

type Request struct {
	Method *string
	URL    *string
	Body   *string
}

type Handler struct {
	t        *testing.T
	Requests []*http.Request
}

var _ rest.Handler = &Handler{}

func New(t *testing.T) *Handler {
	return &Handler{t: t}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) error {
	h.Requests = append(h.Requests, r)
	return nil
}

func (h *Handler) Assert(wants []*Request) {
	// TODO: assert length is the same
	// TODO: for each, call Assert(h.t, got, want)
}

func Assert(t *testing.T, r *http.Request, args ...string) {
	var (
		wantMethod, wantURL, wantBody *string
	)

	switch {
	case len(args) > 2:
		wantBody = &args[2]
		fallthrough
	case len(args) > 1:
		wantURL = &args[1]
		fallthrough
	case len(args) > 0:
		wantMethod = &args[0]
		fallthrough
	default:
	}

	if wantMethod != nil {
		if want, got := *wantMethod, r.Method; want != got {
			t.Errorf("wanted method %q, got method %q", want, got)
		}
	}

	if wantURL != nil {
		if want, got := *wantURL, r.URL.String(); want != got {
			t.Errorf("wanted url %q, got url %q", want, got)
		}
	}

	if wantBody != nil {
		if want, got := *wantBody, readBody(t, r); want != got {
			t.Errorf("wanted body %q, got body %q", want, got)
		}
	}
}

func JSON(t *testing.T, o any) string {
	data, err := json.Marshal(o)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func readBody(t *testing.T, r *http.Request) string {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

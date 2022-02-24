package v0_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	api "github.com/ankeesler/spirits/pkg/v0"
	"github.com/stretchr/testify/require"
)

func TestManifests(t *testing.T) {
	s := httptest.NewServer(api.New())
	t.Cleanup(s.Close)

	c := &http.Client{}

	steps := []struct {
		name       string
		method     string
		header     http.Header
		wantStatus int
		wantHeader http.Header
		wantBody   []byte
	}{
		{
			name:       "DELETE not allowed",
			method:     http.MethodDelete,
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:       "PUT not allowed",
			method:     http.MethodPut,
			wantStatus: http.StatusMethodNotAllowed,
		},
	}
	for _, step := range steps {
		step := step
		req, err := http.NewRequest(step.method, s.URL+"/rooms/abc123/manifests", nil)
		require.NoError(t, err, "step: "+step.name)

		rsp, err := c.Do(req)
		require.NoError(t, err, "step: "+step.name)
		require.Equal(t, step.wantStatus, rsp.StatusCode, "step: "+step.name)
		if step.wantHeader != nil {
			step.wantHeader.Add("Content-Length", fmt.Sprintf("%d", len(step.wantBody)))
			rsp.Header.Del("Date")
			require.Equal(t, step.wantHeader, rsp.Header, "step: "+step.name)
		}
		if step.wantBody == nil {
			step.wantBody = []byte{}
		}
		require.Equal(t, step.wantBody, readBody(t, rsp.Body), "step: "+step.name)

	}
}

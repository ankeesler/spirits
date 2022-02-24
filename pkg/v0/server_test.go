package v0_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	api "github.com/ankeesler/spirits/pkg/v0"
	"github.com/stretchr/testify/require"
)

func TestServer(t *testing.T) {
	s := httptest.NewServer(api.New())
	t.Cleanup(s.Close)

	c := &http.Client{}

	t.Run("path not found", func(t *testing.T) {
		for _, path := range []string{
			"/rooms",
			"/rooms/events",
			"/roomz/abc123",
			"/rooms/abc123/foo",
			"/marshmallow",
		} {
			req, err := http.NewRequest(http.MethodGet, s.URL+"/"+path, nil)
			require.NoError(t, err)

			rsp, err := c.Do(req)
			require.NoError(t, err)
			require.Equal(t, http.StatusNotFound, rsp.StatusCode)
		}
	})
}

func readBody(t *testing.T, body io.Reader) []byte {
	bodyBytes, err := io.ReadAll(body)
	require.NoError(t, err)
	return bodyBytes
}

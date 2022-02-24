package v0_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	api "github.com/ankeesler/spirits/pkg/v0"
)

func TestRooms(t *testing.T) {
	s := httptest.NewServer(api.New())
	t.Cleanup(s.Close)

	c := &http.Client{}

	t.Run("method not allowed", func(t *testing.T) {
		for _, method := range []string{
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodHead,
			http.MethodOptions,
			http.MethodPatch,
		} {
			req, err := http.NewRequest(method, s.URL+"/rooms/abc123/events", nil)
			require.NoError(t, err)

			rsp, err := c.Do(req)
			require.NoError(t, err)
			require.Equal(t, http.StatusMethodNotAllowed, rsp.StatusCode)
		}
	})

	t.Run("http-only client", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, s.URL+"/rooms/abc123/events", nil)
		require.NoError(t, err)

		rsp, err := c.Do(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, rsp.StatusCode)
	})

	t.Run("2 clients", func(t *testing.T) {
		// ...
	})

	t.Run("10 clients", func(t *testing.T) {
		// ...
	})
}

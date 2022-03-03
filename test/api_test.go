package test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAPI(t *testing.T) {
	baseURL := serverBaseURL(t)

	steps := []struct {
		name           string
		req            *http.Request
		wantStatusCode int
		wantBody       string
	}{
		// /battle happy paths
		{
			name:           "same speed",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/good-spirits.json")),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/good-spirits.txt"),
		},
		{
			name:           "double speed",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/good-spirits-double-speed.json")),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/good-spirits-double-speed.txt"),
		},
		{
			name:           "triple speed",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/good-spirits-triple-speed.json")),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/good-spirits-triple-speed.txt"),
		},
		{
			name:           "3:2 speed",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/good-spirits-3-to-2-speed.json")),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/good-spirits-3-to-2-speed.txt"),
		},
		{
			name:           "with defense",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/good-spirits-with-defense.json")),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/good-spirits-with-defense.txt"),
		},
		{
			name:           "bolster",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/good-spirits-with-bolster.json")),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/good-spirits-with-bolster.txt"),
		},
		{
			name:           "drain",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/good-spirits-with-drain.json")),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/good-spirits-with-drain.txt"),
		},
		{
			name:           "charge",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/good-spirits-with-charge.json")),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/good-spirits-with-charge.txt"),
		},
		{
			name:           "multi-move roundrobin",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/good-spirits-with-roundrobin.json")),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/good-spirits-with-roundrobin.txt"),
		},
		{
			name:           "multi-move random",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle?seed=1", readFile(t, "testdata/good-spirits-with-random.json")),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/good-spirits-with-random.txt"),
		},
		// /battle sad paths
		{
			name:           "1 spirit",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/too-few-spirits.json")),
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "must provide 2 spirits\n",
		},
		{
			name:           "3 spirits",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/too-many-spirits.json")),
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "must provide 2 spirits\n",
		},
		{
			name:           "not found",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/nope", readFile(t, "testdata/good-spirits.json")),
			wantStatusCode: http.StatusNotFound,
		},
		{
			name:           "method not allowed",
			req:            newRequest(t, http.MethodPut, baseURL+"/api/battle", readFile(t, "testdata/good-spirits.json")),
			wantStatusCode: http.StatusMethodNotAllowed,
		},
		{
			name:           "invalid body",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", "42"),
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "cannot decode body: json: cannot unmarshal number into Go value of type []*api.Spirit\n",
		},
		{
			name:           "infinite loop",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/powerless-spirits.json")),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/powerless-spirits.txt"),
		},
		{
			name:           "unrecognized action",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/unrecognized-action.json")),
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "unrecognized action: \"tuna\"\n",
		},
		{
			name:           "unrecognized intelligence",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/unrecognized-intelligence.json")),
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "unrecognized intelligence: \"tuna\"\n",
		},
		{
			name:           "/battle bad seed",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/spirit?seed=tuna", ""),
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "invalid seed\n",
		},
		// /spirit happy paths
		{
			name:           "generated spirits with seed 1",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/spirit?seed=1", ""),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/generated-spirits-seed-1.json"),
		},
		{
			name:           "generated spirits with seed 2",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/spirit?seed=2", ""),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/generated-spirits-seed-2.json"),
		},
		// /spirit sad paths
		{
			name:           "/spirit wrong method",
			req:            newRequest(t, http.MethodPut, baseURL+"/api/spirit?seed=2", ""),
			wantStatusCode: http.StatusMethodNotAllowed,
		},
		{
			name:           "/spirit bad seed",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/spirit?seed=tuna", ""),
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "invalid seed\n",
		},
	}
	for _, step := range steps {
		t.Logf("step: %s", step.name)
		t.Logf("req: %s %s", step.req.Method, step.req.URL)
		rsp, err := http.DefaultClient.Do(step.req)
		require.NoError(t, err)

		gotBody, err := io.ReadAll(rsp.Body)
		require.Equalf(t, step.wantStatusCode, rsp.StatusCode, "body: %q", string(gotBody))
		require.NoError(t, err)
		require.Equal(t, step.wantBody, string(gotBody))
	}

	t.Log("step: generated spirits are valid")
	for i := 0; i < 20; i++ {
		// Generate spirits.
		req := newRequest(t, http.MethodPost, baseURL+"/api/spirit", "")
		rsp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)

		gotBody, err := io.ReadAll(rsp.Body)
		require.Equalf(t, http.StatusOK, rsp.StatusCode, "body: %q", string(gotBody))
		require.NoError(t, err)

		// Make sure spirits are valid.
		req = newRequest(t, http.MethodPost, baseURL+"/api/battle", string(gotBody))
		rsp, err = http.DefaultClient.Do(req)
		require.NoError(t, err)

		gotBody, err = io.ReadAll(rsp.Body)
		require.Equalf(t, http.StatusOK, rsp.StatusCode, "body: %q", string(gotBody))
		require.NoError(t, err)
	}
}

func newRequest(t *testing.T, method, url string, body string) *http.Request {
	buf := bytes.NewBuffer([]byte(body))

	req, err := http.NewRequest(method, url, buf)
	require.NoError(t, err)

	return req
}

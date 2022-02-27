package test

import (
	"bytes"
	"io"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAPI(t *testing.T) {
	baseURL := serverBaseURL(t)

	steps := []struct {
		name           string
		req            *http.Request
		wantStatusCode int
		wantBodySuffix string
	}{
		{
			name:           "run battle",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battles", readFile(t, "testdata/good-spirits.json")),
			wantStatusCode: http.StatusOK,
			wantBodySuffix: readFile(t, "testdata/good-spirits.txt"),
		},
		{
			name:           "1 spirit",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battles", readFile(t, "testdata/too-few-spirits.json")),
			wantStatusCode: http.StatusBadRequest,
			wantBodySuffix: "must provide 2 spirits\n",
		},
		{
			name:           "3 spirits",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battles", readFile(t, "testdata/too-many-spirits.json")),
			wantStatusCode: http.StatusBadRequest,
			wantBodySuffix: "must provide 2 spirits\n",
		},
		{
			name:           "not found",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/nope", readFile(t, "testdata/good-spirits.json")),
			wantStatusCode: http.StatusNotFound,
		},
		{
			name:           "method not allowed",
			req:            newRequest(t, http.MethodPut, baseURL+"/api/battles", readFile(t, "testdata/good-spirits.json")),
			wantStatusCode: http.StatusMethodNotAllowed,
		},
		{
			name:           "invalid body",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battles", "42"),
			wantStatusCode: http.StatusBadRequest,
			wantBodySuffix: "cannot decode body: json: cannot unmarshal number into Go value of type []*spirit.Spirit\n",
		},
		{
			name:           "infinite loop",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battles", readFile(t, "testdata/powerless-spirits.json")),
			wantStatusCode: http.StatusOK,
			wantBodySuffix: "> error: too many turns\n",
		},
	}
	for _, step := range steps {
		t.Logf("step: %s", step.name)
		t.Logf("req: %s %s", step.req.Method, step.req.URL)
		rsp, err := http.DefaultClient.Do(step.req)
		require.NoError(t, err)
		require.Equal(t, step.wantStatusCode, rsp.StatusCode)

		gotBody, err := io.ReadAll(rsp.Body)
		require.NoError(t, err)
		require.Truef(
			t,
			strings.HasSuffix(string(gotBody), step.wantBodySuffix),
			"body %q does not end in %q",
			string(gotBody),
			step.wantBodySuffix,
		)
	}
}

func build(t *testing.T) string {
	t.Helper()

	execPath := filepath.Join(t.TempDir(), "spirits")
	output, err := exec.Command(
		"go",
		"build",
		"-o",
		execPath,
		"../../cmd/spirits",
	).CombinedOutput()
	require.NoErrorf(t, err, "output: %s", string(output))
	return execPath
}

func newRequest(t *testing.T, method, url string, body string) *http.Request {
	buf := bytes.NewBuffer([]byte(body))

	req, err := http.NewRequest(method, url, buf)
	require.NoError(t, err)

	return req
}

package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/ankeesler/spirits/internal/spirit"
	"github.com/stretchr/testify/require"
)

func TestAPI(t *testing.T) {
	if !serverUnderTest.remote {
		spiritsExecPath := build(t)
		start(t, spiritsExecPath)
	}

	steps := []struct {
		name           string
		req            *http.Request
		wantStatusCode int
		wantBodySuffix string
	}{
		{
			name: "run battle",
			req: newRequest(t, http.MethodPost, "/api/battles", []*spirit.Spirit{
				{Name: "a", Health: 3, Power: 1},
				{Name: "b", Health: 3, Power: 2},
			}),
			wantStatusCode: http.StatusOK,
			wantBodySuffix: `> summary
  a: 3
  b: 3
> summary
  a: 3
  b: 2
> summary
  a: 1
  b: 2
> summary
  a: 1
  b: 1
> summary
  a: 0
  b: 1
`,
		},
		{
			name: "1 spirit",
			req: newRequest(t, http.MethodPost, "/api/battles", []*spirit.Spirit{
				{Name: "a", Health: 3, Power: 1},
			}),
			wantStatusCode: http.StatusBadRequest,
			wantBodySuffix: "must provide 2 spirits\n",
		},
		{
			name: "3 spirits",
			req: newRequest(t, http.MethodPost, "/api/battles", []*spirit.Spirit{
				{Name: "a", Health: 3, Power: 1},
				{Name: "b", Health: 3, Power: 1},
				{Name: "c", Health: 3, Power: 1},
			}),
			wantStatusCode: http.StatusBadRequest,
			wantBodySuffix: "must provide 2 spirits\n",
		},
		{
			name: "not found",
			req: newRequest(t, http.MethodPost, "/api/nope", []*spirit.Spirit{
				{Name: "a", Health: 3, Power: 1},
				{Name: "b", Health: 3, Power: 2},
			}),
			wantStatusCode: http.StatusNotFound,
		},
		{
			name: "method not allowed",
			req: newRequest(t, http.MethodPut, "/api/battles", []*spirit.Spirit{
				{Name: "a", Health: 3, Power: 1},
				{Name: "b", Health: 3, Power: 2},
			}),
			wantStatusCode: http.StatusMethodNotAllowed,
		},
		{
			name:           "invalid body",
			req:            newRequest(t, http.MethodPost, "/api/battles", 42),
			wantStatusCode: http.StatusBadRequest,
			wantBodySuffix: "cannot decode body: json: cannot unmarshal number into Go value of type []*spirit.Spirit\n",
		},
		{
			name: "infinite loop",
			req: newRequest(t, http.MethodPost, "/api/battles", []*spirit.Spirit{
				{Name: "a", Health: 3, Power: 0},
				{Name: "b", Health: 3, Power: 0},
			}),
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

func start(t *testing.T, execPath string) {
	t.Helper()

	stdout, stderr := bytes.NewBuffer([]byte{}), bytes.NewBuffer([]byte{})
	cmd := exec.Command(execPath)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Start()
	require.NoError(t, err)

	require.Eventually(t, func() bool {
		_, err := http.Get(serverUnderTest.baseURL)
		return err == nil
	}, time.Second*10, time.Second*1)

	cmdErrChan := make(chan error)
	t.Cleanup(func() {
		if err := cmd.Process.Kill(); err != nil {
			t.Errorf("could not kill process: %s", err.Error())
		}
		if err := <-cmdErrChan; err != nil {
			t.Logf("process returned error: %s", err.Error())
		}
		if t.Failed() {
			t.Logf("process stdout:\n%s", stdout.String())
			t.Logf("process stderr:\n%s", stderr.String())
		}
	})
	go func() { cmdErrChan <- cmd.Wait() }()
}

func newRequest(t *testing.T, method, urlPath string, body interface{}) *http.Request {
	buf := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(buf).Encode(body)
	require.NoError(t, err)

	req, err := http.NewRequest(method, serverUnderTest.baseURL+urlPath, buf)
	require.NoError(t, err)

	return req
}

package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/ankeesler/spirits/internal/spirit"
	"github.com/stretchr/testify/require"
)

const spiritsTestBaseURLEnvVar = "SPIRITS_TEST_BASE_URL"

var serverUnderTest = struct {
	baseURL string
	remote  bool
}{
	baseURL: "http://localhost:12345",
	remote:  false,
}

func TestMain(m *testing.M) {
	if val := os.Getenv(spiritsTestBaseURLEnvVar); len(val) > 0 {
		serverUnderTest.baseURL = val
		serverUnderTest.remote = true
	}

	os.Exit(m.Run())
}

func TestSpirits(t *testing.T) {
	t.Log("sup")
	if !serverUnderTest.remote {
		spiritsExecPath := build(t)
		start(t, spiritsExecPath)
	}

	steps := []struct {
		name           string
		req            *http.Request
		wantStatusCode int
		wantBody       string
	}{
		{
			name: "run battle",
			req: newRequest(t, http.MethodPost, "/battles", []*spirit.Spirit{
				{Name: "a", Health: 3, Power: 1},
				{Name: "b", Health: 3, Power: 2},
			}),
			wantStatusCode: http.StatusOK,
			wantBody: `> summary
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
			req: newRequest(t, http.MethodPost, "/battles", []*spirit.Spirit{
				{Name: "a", Health: 3, Power: 1},
			}),
			wantStatusCode: http.StatusBadRequest,
			wantBody: `{"error": "must provide 2 spirits"}
`,
		},
		{
			name: "3 spirits",
			req: newRequest(t, http.MethodPost, "/battles", []*spirit.Spirit{
				{Name: "a", Health: 3, Power: 1},
				{Name: "b", Health: 3, Power: 1},
				{Name: "c", Health: 3, Power: 1},
			}),
			wantStatusCode: http.StatusBadRequest,
			wantBody: `{"error": "must provide 2 spirits"}
`,
		},
		{
			name: "not found",
			req: newRequest(t, http.MethodPost, "/nope", []*spirit.Spirit{
				{Name: "a", Health: 3, Power: 1},
				{Name: "b", Health: 3, Power: 2},
			}),
			wantStatusCode: http.StatusNotFound,
		},
		{
			name: "method not allowed",
			req: newRequest(t, http.MethodPut, "/battles", []*spirit.Spirit{
				{Name: "a", Health: 3, Power: 1},
				{Name: "b", Health: 3, Power: 2},
			}),
			wantStatusCode: http.StatusMethodNotAllowed,
		},
		{
			name:           "invalid body",
			req:            newRequest(t, http.MethodPost, "/battles", 42),
			wantStatusCode: http.StatusBadRequest,
			wantBody: `{"error": "cannot decode body: json: cannot unmarshal number into Go value of type []*spirit.Spirit"}
`,
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
		require.Equal(t, step.wantBody, string(gotBody))
	}
}

func build(t *testing.T) string {
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

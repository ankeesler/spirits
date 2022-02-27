package test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

var buildServerOnce sync.Once
var runServerOncePerTest = make(map[*testing.T]struct{})

func serverBaseURL(t *testing.T) string {
	t.Helper()

	// Check to see if we are running against existing server.
	if envVar := os.Getenv("SPIRITS_TEST_BASE_URL"); len(envVar) > 0 {
		return envVar
	}

	// Build binary, but only once per test run.
	serverExe := filepath.Join(testDir, "spirits-under-test")
	buildServerOnce.Do(func() {
		output, err := exec.Command(
			"go",
			"build",
			"-o",
			serverExe,
			"../../cmd/spirits",
		).CombinedOutput()
		require.NoErrorf(t, err, "output: %s", string(output))
	})

	// Actually start the binary, but only if it hasn't been started by this test.
	const port = 12345
	if _, ok := runServerOncePerTest[t]; !ok {
		startProcess(t, checkPortFunc(port), serverExe, "-web-assets-dir", "../../public")
		runServerOncePerTest[t] = struct{}{}
	}

	return fmt.Sprintf("http://localhost:%d", port)
}

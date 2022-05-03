package client

import (
	"os/exec"
	"path"
	"testing"

	"github.com/ankeesler/spirits/test/lib"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	// Build the client
	clientPath := build(t)

	// Run the default command (generate battle+spirits)
	// TODO: lib.NewTestConfig(t)
	tc := lib.NewTestConfig(t)
	output, err := exec.Command(clientPath, "-namespace", tc.Namespace.Name).CombinedOutput()
	require.NoErrorf(t, err, "output: %q", string(output))

	// Make sure we have a battle created
}

func build(t *testing.T) string {
	t.Helper()
	clientPath := path.Join(t.TempDir(), "client")
	output, err := exec.Command("go", "build", "-o", clientPath, "../../cmd/client").CombinedOutput()
	require.NoErrorf(t, err, "output: %q", string(output))
	return clientPath
}

package test

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCLI(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	t.Log("building cli...")
	cliPath := buildCLI(ctx, t)
	t.Logf("cli built to %q", cliPath)

	t.Run("battle", func(t *testing.T) {
		output, err := exec.CommandContext(
			ctx,
			cliPath,
			"battle",
			"-manifest",
			"testdata/good-manifest.yaml",
			"-teams",
			"my-team,their-team",
		).CombinedOutput()
		require.NoError(t, err, string(output))
	})
}

func buildCLI(ctx context.Context, t *testing.T) string {
	t.Helper()

	dir, err := os.MkdirTemp("", "spirits-integration-*")
	require.NoError(t, err)

	cliPath := filepath.Join(dir, "spirits")

	var output []byte
	output, err = exec.CommandContext(
		ctx,
		"go",
		"build",
		"-o",
		cliPath,
		"../../cmd/spirits",
	).CombinedOutput()
	require.NoError(t, err, string(output))

	t.Cleanup(func() {
		require.NoError(t, os.RemoveAll(dir))
	})

	return cliPath
}

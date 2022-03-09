package test

import (
	"os"
	"path/filepath"
	"testing"
)

var testDir = filepath.Join(os.TempDir(), "spirits-test")

func TestMain(m *testing.M) {
	defer func() {
		if err := os.RemoveAll(testDir); err != nil {
			panic(err)
		}
	}()

	os.Exit(m.Run())
}

package test

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	defer func() {
		if err := os.RemoveAll(testDir); err != nil {
			panic(err)
		}
	}()

	os.Exit(m.Run())
}

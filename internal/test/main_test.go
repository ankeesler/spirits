package test

import (
	"os"
	"testing"
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

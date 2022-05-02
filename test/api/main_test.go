package api

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	if _, ok := os.LookupEnv("SPIRITS_TEST_INTEGRATION"); !ok {
		fmt.Println(strings.Repeat("!", 80))
		fmt.Println("  WARNING")
		fmt.Println("    skipping integration tests because env var 'SPIRITS_TEST_INTEGRATION' not set")
		fmt.Println("  WARNING")
		fmt.Println(strings.Repeat("!", 80))
		os.Exit(0)
	}
	os.Exit(m.Run())
}

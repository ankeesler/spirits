package test

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func startProcess(t *testing.T, readyFunc func() bool, exePath string, exeArgs ...string) {
	t.Helper()

	stdout, stderr := bytes.NewBuffer([]byte{}), bytes.NewBuffer([]byte{})
	cmd := exec.Command(exePath, exeArgs...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Start()
	require.NoError(t, err)
	t.Logf("started process %s %s", exePath, strings.Join(exeArgs, " "))

	require.Eventually(t, readyFunc, time.Second*5, time.Second*1)

	cmdErrChan := make(chan error)
	t.Cleanup(func() {
		if err := cmd.Process.Kill(); err != nil {
			t.Errorf("could not kill process %q: %s", exePath, err.Error())
		}
		if err := <-cmdErrChan; err != nil {
			t.Logf("process %q returned error: %s", exePath, err.Error())
		}
		if t.Failed() {
			t.Logf("process %q stdout:\n%s", exePath, stdout.String())
			t.Logf("process %q stderr:\n%s", exePath, stderr.String())
		}
	})
	go func() { cmdErrChan <- cmd.Wait() }()
}

func checkPortFunc(port int) func() bool {
	return func() bool {
		_, err := http.Get(fmt.Sprintf("http://localhost:%d", port))
		return err == nil
	}
}

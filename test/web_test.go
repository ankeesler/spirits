package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func TestWeb(t *testing.T) {
	const chromedriverPort = 4444
	startProcess(
		t,
		checkPortFunc(chromedriverPort),
		"chromedriver",
		fmt.Sprintf("--port=%d", chromedriverPort),

		// Uncomment this to get more verbose chrome logs.
		// "--verbose",
	)
	wd := startBrowser(t, chromedriverPort)

	steps := []struct {
		name string
		do   func(t *testing.T, wd selenium.WebDriver)
	}{
		{
			name: "get the homepage with seed set to 1",
			do: func(t *testing.T, wd selenium.WebDriver) {
				homepage := serverBaseURL(t) + "/?seed=1"
				err := wd.Get(homepage)
				require.NoError(t, err)
			},
		},
		{
			name: "first happy path",
			do: func(t *testing.T, wd selenium.WebDriver) {
				input, err := wd.FindElement(selenium.ByID, "input")
				require.NoError(t, err)

				err = input.SendKeys(readFile(t, "testdata/good-spirits.json"))
				require.NoError(t, err)

				testHappyPath(t, wd)
			},
		},
		{
			name: "sad path",
			do: func(t *testing.T, wd selenium.WebDriver) {
				input, err := wd.FindElement(selenium.ByID, "input")
				require.NoError(t, err)

				err = input.SendKeys(selenium.BackspaceKey)
				require.NoError(t, err)

				testSadPath(t, wd)
			},
		},
		{
			name: "second happy path",
			do: func(t *testing.T, wd selenium.WebDriver) {
				input, err := wd.FindElement(selenium.ByID, "input")
				require.NoError(t, err)

				err = input.SendKeys("]")
				require.NoError(t, err)

				testHappyPath(t, wd)
			},
		},
		{
			name: "generate spirits with seed",
			do: func(t *testing.T, wd selenium.WebDriver) {
				input, err := wd.FindElement(selenium.ByID, "input")
				require.NoError(t, err)
				output, err := wd.FindElement(selenium.ByID, "output")
				require.NoError(t, err)
				generateSpirits, err := wd.FindElement(selenium.ByID, "generate-spirits")
				require.NoError(t, err)

				err = generateSpirits.Click()
				require.NoError(t, err)

				inputTextFunc := func() (string, error) { return input.GetAttribute("value") }
				outputTextFunc := func() (string, error) { return output.GetAttribute("value") }
				eventually(t, inputTextFunc, readFile(t, "testdata/generated-spirits-seed-1.json"), true, true)
				eventually(t, outputTextFunc, readFile(t, "testdata/generated-spirits-seed-1.txt"), true, true)
			},
		},
		{
			name: "get the homepage without a seed query",
			do: func(t *testing.T, wd selenium.WebDriver) {
				err := wd.Get(serverBaseURL(t))
				require.NoError(t, err)
			},
		},
		{
			name: "run generate spirits again and make sure it is random",
			do: func(t *testing.T, wd selenium.WebDriver) {
				input, err := wd.FindElement(selenium.ByID, "input")
				require.NoError(t, err)
				generateSpirits, err := wd.FindElement(selenium.ByID, "generate-spirits")
				require.NoError(t, err)

				err = generateSpirits.Click()
				require.NoError(t, err)

				inputTextFunc := func() (string, error) { return input.GetAttribute("value") }
				eventually(t, inputTextFunc, readFile(t, "testdata/generated-spirits-seed-1.json"), false, true)
			},
		},
	}
	for _, step := range steps {
		t.Logf("step: %s", step.name)
		step.do(t, wd)
	}
}

func startBrowser(t *testing.T, port int) selenium.WebDriver {
	// Uncomment this to get more verbose selenium logs.
	// selenium.SetDebug(true)

	capabilities := selenium.Capabilities{}
	capabilities.AddChrome(chrome.Capabilities{
		// Comment out this line to see the tests happen in a visible browser window.
		Args: []string{"--headless"},
	})
	wd, err := selenium.NewRemote(capabilities, fmt.Sprintf("http://localhost:%d", port))
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := wd.Close(); err != nil {
			t.Logf("could not close webdriver: %s", err.Error())
		}
	})

	err = wd.SetImplicitWaitTimeout(time.Second * 5)
	require.NoError(t, err)

	return wd
}

func testHappyPath(t *testing.T, wd selenium.WebDriver) {
	t.Helper()

	output, err := wd.FindElement(selenium.ByID, "output")
	require.NoError(t, err)
	errorMessage, err := wd.FindElement(selenium.ByID, "error-message")
	require.NoError(t, err)

	outputTextFunc := func() (string, error) { return output.GetAttribute("value") }
	eventually(t, outputTextFunc, "⌛", true, true)
	eventually(t, outputTextFunc, readFile(t, "testdata/good-spirits.txt"), true, true)
	eventually(t, errorMessage.Text, "", false, false)
}

func testSadPath(t *testing.T, wd selenium.WebDriver) {
	t.Helper()

	output, err := wd.FindElement(selenium.ByID, "output")
	require.NoError(t, err)
	errorMessage, err := wd.FindElement(selenium.ByID, "error-message")
	require.NoError(t, err)

	outputTextFunc := func() (string, error) { return output.GetAttribute("value") }
	eventually(t, outputTextFunc, "⌛", true, true)
	eventually(t, errorMessage.Text, "invalid spirits json", true, true)
}

func eventually(t *testing.T, textFunc func() (string, error), wantText string, once, positive bool) {
	t.Helper()

	requireFunc := require.Never
	condition := "continuous"
	if once {
		requireFunc = require.Eventually
		condition = "once"
	}

	operator := "!="
	if positive {
		operator = "=="
	}

	conditionFunc := func() bool {
		gotText, err := textFunc()
		require.NoError(t, err)

		t.Logf("%s: %q %s %q", condition, wantText, operator, gotText)
		if positive {
			return (wantText == gotText)
		}
		return wantText != gotText
	}

	requireFunc(t, conditionFunc, time.Second*3, time.Second)
}

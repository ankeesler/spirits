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

	// Get the homepage.
	homepage := serverBaseURL(t) + "/"
	err := wd.Get(homepage)
	require.NoError(t, err)

	// Get the input box and error message.
	input, err := wd.FindElement(selenium.ByID, "input")
	require.NoError(t, err)
	errorMessage, err := wd.FindElement(selenium.ByID, "error-message")
	require.NoError(t, err)

	happyPath := func() {
		eventuallyOutputText(t, wd, "⌛")
		eventuallyOutputText(t, wd, readFile(t, "testdata/good-spirits.txt"))
		require.Never(t, func() bool {
			gotErrorMessage, err := errorMessage.Text()
			require.NoError(t, err)
			if len(gotErrorMessage) > 0 {
				t.Logf("never: wanted error message to be empty, got %q", gotErrorMessage)
				return true
			}
			return false
		}, time.Second*3, time.Second)
	}

	// When we pass a good spirits JSON, we get a success.
	err = input.SendKeys(readFile(t, "testdata/good-spirits.json"))
	require.NoError(t, err)
	happyPath()

	// When we paste a bad spirits JSON, we get an error.
	err = input.SendKeys(selenium.BackspaceKey)
	require.NoError(t, err)
	eventuallyOutputText(t, wd, "⌛")
	wantErrorMessage := "invalid spirits json"
	require.Eventually(t, func() bool {
		gotErrorMessage, err := errorMessage.Text()
		require.NoError(t, err)
		if wantErrorMessage != gotErrorMessage {
			t.Logf("eventually: want %q, got %q", wantErrorMessage, gotErrorMessage)
			return false
		}
		return true
	}, time.Second*3, time.Second)

	// We should be able to go back to the happy path.
	err = input.SendKeys("]")
	require.NoError(t, err)
	happyPath()
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

func eventuallyOutputText(t *testing.T, wd selenium.WebDriver, wantOutputText string) {
	t.Helper()

	output, err := wd.FindElement(selenium.ByID, "output")
	require.NoError(t, err)

	require.Eventually(t, func() bool {
		gotOutputText, err := output.GetAttribute("value")
		require.NoError(t, err)
		if wantOutputText != gotOutputText {
			t.Logf("eventually: want %q, got %q", wantOutputText, gotOutputText)
			return false
		}
		return true
	}, time.Second*3, time.Second)
}

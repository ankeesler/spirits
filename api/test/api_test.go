package test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/ankeesler/spirits/api/internal/api"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	name string
	msgs []testMsg
}

type testMsg struct {
	in    bool // true: tx, false: rx
	reset bool // true to reset connection
	msg   interface{}
}

func TestAPI(t *testing.T) {
	baseURL := serverBaseURL(t)
	u, err := url.Parse(baseURL)
	require.NoError(t, err)

	t.Run("http connection to websocket path", func(t *testing.T) {
		rsp, err := http.DefaultClient.Get("http://" + u.Host + "/api/battle")
		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, rsp.StatusCode)
		io.Copy(io.Discard, rsp.Body)
	})

	t.Run("invalid seed", func(t *testing.T) {
		_, rsp, err := websocket.DefaultDialer.Dial("ws://"+u.Host+"/api/battle?seed=tuna", nil)
		require.EqualError(t, err, "websocket: bad handshake")
		require.Equal(t, http.StatusBadRequest, rsp.StatusCode)

		data, err := io.ReadAll(rsp.Body)
		require.NoError(t, err)
		require.Equal(t, "invalid seed\n", string(data))
	})

	var c *websocket.Conn
	dial := func() {
		var err error
		c, _, err = websocket.DefaultDialer.Dial("ws://"+u.Host+"/api/battle?seed=1", nil)
		require.NoError(t, err)
		t.Cleanup(func() {
			c.Close()
		})
	}
	dial()

	steps := []testCase{
		{
			name: "when the message type is invalid it doesn't take down the server",
			msgs: []testMsg{
				{
					in: true,
					msg: api.Message{
						Type: "invalid",
					},
				},
			},
		},
		{
			name: "when the message type is the wrong schema it doesn't take down the server",
			msgs: []testMsg{
				{
					in: true,
					msg: map[string]interface{}{
						"type": 1,
					},
				},
			},
		},
		{
			name: "when the message details are the wrong schema it doesn't take down the server",
			msgs: []testMsg{
				{
					in: true,
					msg: map[string]interface{}{
						"type":    api.MessageTypeActionReq,
						"details": 1,
					},
				},
			},
		},
		// Happy path
		{
			name: "battle-start with 2 spirits with single human interaction",
			msgs: []testMsg{
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeBattleStart,
						Details: &api.MessageDetailsBattleStart{
							Spirits: readSpirits(t, "testdata/good-spirits-with-single-human-interaction.json"),
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeActionReq,
						Details: &api.MessageDetailsActionReq{
							Output: "> summary\n  a: 3\n  b: 3\n",
							Spirit: api.Spirit{
								Name:         "a",
								Health:       3,
								Power:        1,
								Agility:      1,
								Intelligence: "human",
							},
						},
					},
				},
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeActionRsp,
						Details: &api.MessageDetailsActionRsp{
							Spirit: api.Spirit{
								Name:         "a",
								Health:       3,
								Power:        1,
								Agility:      1,
								Intelligence: "human",
							},
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeActionReq,
						Details: &api.MessageDetailsActionReq{
							Output: "> summary\n  a: 3\n  b: 2\n> summary\n  a: 1\n  b: 2\n",
							Spirit: api.Spirit{
								Name:         "a",
								Health:       3,
								Power:        1,
								Agility:      1,
								Intelligence: "human",
							},
						},
					},
				},
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeActionRsp,
						Details: &api.MessageDetailsActionRsp{
							Spirit: api.Spirit{
								Name:         "a",
								Health:       1,
								Power:        1,
								Agility:      1,
								Intelligence: "human",
							},
							ID: "attack",
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeBattleStop,
						Details: &api.MessageDetailsBattleStop{
							Output: "> summary\n  a: 1\n  b: 1\n> summary\n  a: 0\n  b: 1\n",
						},
					},
				},
			},
		},
		{
			name: "battle-start with 2 spirits with double human interaction",
			msgs: []testMsg{
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeBattleStart,
						Details: &api.MessageDetailsBattleStart{
							Spirits: readSpirits(t, "testdata/good-spirits-with-double-human-interaction.json"),
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeActionReq,
						Details: &api.MessageDetailsActionReq{
							Output: "> summary\n  a: 3\n  b: 3\n",
							Spirit: api.Spirit{
								Name:         "a",
								Health:       3,
								Power:        1,
								Agility:      1,
								Intelligence: "human",
								Actions:      []string{"charge"},
							},
						},
					},
				},
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeActionRsp,
						Details: &api.MessageDetailsActionRsp{
							Spirit: api.Spirit{
								Name: "a",
							},
							ID: "", // defaults to first entry (charge)
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeActionReq,
						Details: &api.MessageDetailsActionReq{
							Output: "> summary\n  a: 2\n  b: 1\n",
							Spirit: api.Spirit{
								Name:         "b",
								Health:       3,
								Power:        2,
								Agility:      1,
								Intelligence: "human",
								Actions:      []string{"charge", "bolster"},
							},
						},
					},
				},
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeActionRsp,
						Details: &api.MessageDetailsActionRsp{
							Spirit: api.Spirit{
								Name: "b",
							},
							ID: "bolster",
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeActionReq,
						Details: &api.MessageDetailsActionReq{
							Output: "> summary\n  a: 1\n  b: 1\n",
							Spirit: api.Spirit{
								Name:         "a",
								Health:       3,
								Power:        1,
								Agility:      1,
								Intelligence: "human",
								Actions:      []string{"charge"},
							},
						},
					},
				},
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeActionRsp,
						Details: &api.MessageDetailsActionRsp{
							Spirit: api.Spirit{
								Name: "a",
							},
							ID: "charge",
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeBattleStop,
						Details: &api.MessageDetailsBattleStop{
							Output: "> summary\n  a: 1\n  b: 0\n",
						},
					},
				},
			},
		},
		{
			name: "battle-start with 2 spirits with human interaction and closed connection bounces back",
			msgs: []testMsg{
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeBattleStart,
						Details: &api.MessageDetailsBattleStart{
							Spirits: readSpirits(t, "testdata/good-spirits-with-single-human-interaction.json"),
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeActionReq,
						Details: &api.MessageDetailsActionReq{
							Output: "> summary\n  a: 3\n  b: 3\n",
							Spirit: api.Spirit{
								Name:         "a",
								Health:       3,
								Power:        1,
								Agility:      1,
								Intelligence: "human",
							},
						},
					},
				},
				{
					reset: true,
				},
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeBattleStart,
						Details: &api.MessageDetailsBattleStart{
							Spirits: readSpirits(t, "testdata/good-spirits-with-single-human-interaction.json"),
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeActionReq,
						Details: &api.MessageDetailsActionReq{
							Output: "> summary\n  a: 3\n  b: 3\n",
							Spirit: api.Spirit{
								Name:         "a",
								Health:       3,
								Power:        1,
								Agility:      1,
								Intelligence: "human",
							},
						},
					},
				},
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeActionRsp,
						Details: &api.MessageDetailsActionRsp{
							Spirit: api.Spirit{
								Name:         "a",
								Health:       3,
								Power:        1,
								Agility:      1,
								Intelligence: "human",
							},
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeActionReq,
						Details: &api.MessageDetailsActionReq{
							Output: "> summary\n  a: 3\n  b: 2\n> summary\n  a: 1\n  b: 2\n",
							Spirit: api.Spirit{
								Name:         "a",
								Health:       3,
								Power:        1,
								Agility:      1,
								Intelligence: "human",
							},
						},
					},
				},
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeActionRsp,
						Details: &api.MessageDetailsActionRsp{
							Spirit: api.Spirit{
								Name:         "a",
								Health:       1,
								Power:        1,
								Agility:      1,
								Intelligence: "human",
							},
							ID: "attack",
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeBattleStop,
						Details: &api.MessageDetailsBattleStop{
							Output: "> summary\n  a: 1\n  b: 1\n> summary\n  a: 0\n  b: 1\n",
						},
					},
				},
			},
		},
		{
			name: "battle-start with 2 spirits with human interaction and battle stop bounces back",
			msgs: []testMsg{
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeBattleStart,
						Details: &api.MessageDetailsBattleStart{
							Spirits: readSpirits(t, "testdata/good-spirits-with-single-human-interaction.json"),
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeActionReq,
						Details: &api.MessageDetailsActionReq{
							Output: "> summary\n  a: 3\n  b: 3\n",
							Spirit: api.Spirit{
								Name:         "a",
								Health:       3,
								Power:        1,
								Agility:      1,
								Intelligence: "human",
							},
						},
					},
				},
				{
					in: true,
					msg: api.Message{
						Type:    api.MessageTypeBattleStop,
						Details: &api.MessageDetailsBattleStop{},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeBattleStop,
						Details: &api.MessageDetailsBattleStop{
							Output: "> error: action errored: action canceled: context canceled\n",
						},
					},
				},
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeBattleStart,
						Details: &api.MessageDetailsBattleStart{
							Spirits: readSpirits(t, "testdata/good-spirits-with-single-human-interaction.json"),
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeActionReq,
						Details: &api.MessageDetailsActionReq{
							Output: "> summary\n  a: 3\n  b: 3\n",
							Spirit: api.Spirit{
								Name:         "a",
								Health:       3,
								Power:        1,
								Agility:      1,
								Intelligence: "human",
							},
						},
					},
				},
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeActionRsp,
						Details: &api.MessageDetailsActionRsp{
							Spirit: api.Spirit{
								Name:         "a",
								Health:       3,
								Power:        1,
								Agility:      1,
								Intelligence: "human",
							},
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeActionReq,
						Details: &api.MessageDetailsActionReq{
							Output: "> summary\n  a: 3\n  b: 2\n> summary\n  a: 1\n  b: 2\n",
							Spirit: api.Spirit{
								Name:         "a",
								Health:       3,
								Power:        1,
								Agility:      1,
								Intelligence: "human",
							},
						},
					},
				},
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeActionRsp,
						Details: &api.MessageDetailsActionRsp{
							Spirit: api.Spirit{
								Name:         "a",
								Health:       1,
								Power:        1,
								Agility:      1,
								Intelligence: "human",
							},
							ID: "attack",
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeBattleStop,
						Details: &api.MessageDetailsBattleStop{
							Output: "> summary\n  a: 1\n  b: 1\n> summary\n  a: 0\n  b: 1\n",
						},
					},
				},
			},
		},
		{
			name: "spirit-req",
			msgs: []testMsg{
				{
					in: true,
					msg: api.Message{
						Type:    api.MessageTypeSpiritReq,
						Details: &api.MessageDetailsSpiritReq{},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeSpiritRsp,
						Details: &api.MessageDetailsSpiritRsp{
							Spirits: readSpirits(t, "testdata/generated-spirits-seed-1.json"),
						},
					},
				},
			},
		},
		// Error cases
		{
			name: "battle-start with 0 spirits",
			msgs: []testMsg{
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeBattleStart,
						Details: &api.MessageDetailsBattleStart{
							Spirits: []*api.Spirit{},
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeError,
						Details: &api.MessageDetailsError{
							Reason: "must provide 2 spirits",
						},
					},
				},
			},
		},
		{
			name: "battle-start with 1 spirit",
			msgs: []testMsg{
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeBattleStart,
						Details: &api.MessageDetailsBattleStart{
							Spirits: []*api.Spirit{{Name: "a"}},
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeError,
						Details: &api.MessageDetailsError{
							Reason: "must provide 2 spirits",
						},
					},
				},
			},
		},
		{
			name: "battle-start with 3 spirits",
			msgs: []testMsg{
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeBattleStart,
						Details: &api.MessageDetailsBattleStart{
							Spirits: []*api.Spirit{{Name: "a"}, {Name: "b"}, {Name: "c"}},
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeError,
						Details: &api.MessageDetailsError{
							Reason: "must provide 2 spirits",
						},
					},
				},
			},
		},
		{
			name: "battle-stop without battle-start",
			msgs: []testMsg{
				{
					in: true,
					msg: api.Message{
						Type:    api.MessageTypeBattleStop,
						Details: &api.MessageDetailsBattleStop{},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeError,
						Details: &api.MessageDetailsError{
							Reason: "unexpected battle-stop: no battle running",
						},
					},
				},
			},
		},
		{
			name: "action-request without battle-start",
			msgs: []testMsg{
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeActionReq,
						Details: &api.MessageDetailsActionReq{
							Spirit: api.Spirit{Name: "a"},
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeError,
						Details: &api.MessageDetailsError{
							Reason: `unexpected action-req for spirit: "a"`,
						},
					},
				},
			},
		},
		{
			name: "action-response without battle-start",
			msgs: []testMsg{
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeActionRsp,
						Details: &api.MessageDetailsActionRsp{
							Spirit: api.Spirit{Name: "a"},
							ID:     "whatever",
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeError,
						Details: &api.MessageDetailsError{
							Reason: `unexpected action-rsp with ID "whatever" for spirit: "a"`,
						},
					},
				},
			},
		},
		{
			name: "unsolicited spirit-rsp",
			msgs: []testMsg{
				{
					in: true,
					msg: api.Message{
						Type:    api.MessageTypeSpiritRsp,
						Details: &api.MessageDetailsSpiritRsp{},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeError,
						Details: &api.MessageDetailsError{
							Reason: "unexpected spirit-rsp",
						},
					},
				},
			},
		},
		{
			name: "battle-start with 2 spirits and a battle already exists",
			msgs: []testMsg{
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeBattleStart,
						Details: &api.MessageDetailsBattleStart{
							Spirits: readSpirits(t, "testdata/good-spirits-with-single-human-interaction.json"),
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeActionReq,
						Details: &api.MessageDetailsActionReq{
							Output: "> summary\n  a: 3\n  b: 3\n",
							Spirit: api.Spirit{
								Name:         "a",
								Health:       3,
								Power:        1,
								Agility:      1,
								Intelligence: "human",
							},
						},
					},
				},
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeBattleStart,
						Details: &api.MessageDetailsBattleStart{
							Spirits: readSpirits(t, "testdata/good-spirits-with-single-human-interaction.json"),
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeError,
						Details: &api.MessageDetailsError{
							Reason: "battle already running",
						},
					},
				},
				{
					in: true,
					msg: api.Message{
						Type:    api.MessageTypeBattleStop,
						Details: &api.MessageDetailsBattleStop{},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeBattleStop,
						Details: &api.MessageDetailsBattleStop{
							Output: "> error: action errored: action canceled: context canceled\n",
						},
					},
				},
			},
		},
		{
			name: "unknown action id",
			msgs: []testMsg{
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeBattleStart,
						Details: &api.MessageDetailsBattleStart{
							Spirits: readSpirits(t, "testdata/good-spirits-with-single-human-interaction.json"),
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeActionReq,
						Details: &api.MessageDetailsActionReq{
							Output: "> summary\n  a: 3\n  b: 3\n",
							Spirit: api.Spirit{
								Name:         "a",
								Health:       3,
								Power:        1,
								Agility:      1,
								Intelligence: "human",
							},
						},
					},
				},
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeActionRsp,
						Details: &api.MessageDetailsActionRsp{
							Spirit: api.Spirit{
								Name:         "a",
								Health:       3,
								Power:        1,
								Agility:      1,
								Intelligence: "human",
							},
							ID: "invalid",
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeBattleStop,
						Details: &api.MessageDetailsBattleStop{
							Output: `> error: action errored: unknown action "invalid" for spirit "a"` + "\n",
						},
					},
				},
			},
		},
		{
			name: "unrecognized intelligence type",
			msgs: []testMsg{
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeBattleStart,
						Details: &api.MessageDetailsBattleStart{
							Spirits: readSpirits(t, "testdata/unrecognized-intelligence.json"),
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeError,
						Details: &api.MessageDetailsError{
							Reason: `unrecognized intelligence: "tuna"`,
						},
					},
				},
			},
		},
		{
			name: "unrecognized action type",
			msgs: []testMsg{
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeBattleStart,
						Details: &api.MessageDetailsBattleStart{
							Spirits: readSpirits(t, "testdata/unrecognized-action.json"),
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeError,
						Details: &api.MessageDetailsError{
							Reason: `unrecognized action: "tuna"`,
						},
					},
				},
			},
		},
	}
	steps = append(steps, getAutoTests(t)...)
	for _, step := range steps {
		t.Logf("step: %s", step.name)
		for _, msg := range step.msgs {
			if msg.in {
				err := c.WriteJSON(msg.msg)
				require.NoError(t, err)
				continue
			}

			if msg.reset {
				err := c.Close()
				require.NoError(t, err)
				dial()
				continue
			}

			msgIn, err := readMessage(c)
			require.NoErrorf(t, err, "waiting for %#v", msg.msg.(api.Message))
			require.Equal(t, msg.msg.(api.Message), *msgIn)
		}
	}

	t.Run("generated spirits are valid", func(t *testing.T) {
		for i := 0; i < 20; i++ {
			err := c.WriteJSON(&api.Message{
				Type: api.MessageTypeSpiritReq,
			})
			require.NoError(t, err)

			var m api.Message
			err = c.ReadJSON(&m)
			require.NoError(t, err)
			require.Equalf(t, m.Type, api.MessageTypeSpiritRsp, "wanted spirit-rsp, got %#v", &m)

			err = c.WriteJSON(&api.Message{
				Type: api.MessageTypeBattleStart,
				Details: &api.MessageDetailsBattleStart{
					Spirits: m.Details.(*api.MessageDetailsSpiritRsp).Spirits,
				},
			})
			require.NoError(t, err)

			err = c.ReadJSON(&m)
			require.NoError(t, err)
			require.Equalf(t, m.Type, api.MessageTypeBattleStop, "wanted battle-stop, got %#v", &m)
		}
	})
}

func getAutoTests(t *testing.T) []testCase {
	t.Helper()

	testCases := []testCase{}

	dirName := filepath.Join("testdata", "auto")
	entries, err := os.ReadDir(dirName)
	require.NoError(t, err)
	for _, entry := range entries {
		name := entry.Name()
		if !strings.HasSuffix(name, ".json") {
			t.Logf("ignoring non-json file %q", name)
			continue
		}

		otherName := strings.ReplaceAll(name, ".json", ".txt")
		_, err := os.Stat(filepath.Join(dirName, otherName))
		require.NoErrorf(t, err, "wanted file %q for %q", otherName, name)

		testCases = append(testCases, testCase{
			name: name,
			msgs: []testMsg{
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeBattleStart,
						Details: &api.MessageDetailsBattleStart{
							Spirits: readSpirits(t, filepath.Join(dirName, name)),
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeBattleStop,
						Details: &api.MessageDetailsBattleStop{
							Output: readFile(t, filepath.Join(dirName, otherName)),
						},
					},
				},
			},
		})
	}

	return testCases
}

func readSpirits(t *testing.T, path string) []*api.Spirit {
	t.Helper()

	data, err := os.ReadFile(path)
	require.NoError(t, err)

	var spirits []*api.Spirit
	err = json.Unmarshal(data, &spirits)
	require.NoError(t, err)

	return spirits
}

func readFile(t *testing.T, path string) string {
	data, err := os.ReadFile(path)
	require.NoError(t, err)
	return string(data)
}

func readMessage(c *websocket.Conn) (*api.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	var msg api.Message
	errCh := make(chan error)
	go func() { errCh <- c.ReadJSON(&msg) }()
	select {
	case err := <-errCh:
		return &msg, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

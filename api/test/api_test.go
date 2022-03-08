package test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/ankeesler/spirits/api/internal/api"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
)

func TestAPI(t *testing.T) {
	baseURL := serverBaseURL(t)
	t.Run("HTTP", func(t *testing.T) {
		testHTTPAPI(t, baseURL)
	})
	t.Run("Websocket", func(t *testing.T) {
		testWebsocketAPI(t, baseURL)
	})
}

func testHTTPAPI(t *testing.T, baseURL string) {
	tests := []struct {
		name           string
		req            *http.Request
		wantStatusCode int
		wantBody       string
	}{
		// /battle happy paths
		{
			name:           "same speed",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/good-spirits.json")),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/good-spirits.txt"),
		},
		{
			name:           "double speed",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/good-spirits-double-speed.json")),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/good-spirits-double-speed.txt"),
		},
		{
			name:           "triple speed",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/good-spirits-triple-speed.json")),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/good-spirits-triple-speed.txt"),
		},
		{
			name:           "3:2 speed",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/good-spirits-3-to-2-speed.json")),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/good-spirits-3-to-2-speed.txt"),
		},
		{
			name:           "with defense",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/good-spirits-with-defense.json")),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/good-spirits-with-defense.txt"),
		},
		{
			name:           "bolster",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/good-spirits-with-bolster.json")),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/good-spirits-with-bolster.txt"),
		},
		{
			name:           "drain",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/good-spirits-with-drain.json")),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/good-spirits-with-drain.txt"),
		},
		{
			name:           "charge",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/good-spirits-with-charge.json")),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/good-spirits-with-charge.txt"),
		},
		{
			name:           "multi-move roundrobin",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/good-spirits-with-roundrobin.json")),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/good-spirits-with-roundrobin.txt"),
		},
		{
			name:           "multi-move random",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle?seed=1", readFile(t, "testdata/good-spirits-with-random.json")),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/good-spirits-with-random.txt"),
		},
		// /battle sad paths
		{
			name:           "1 spirit",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/too-few-spirits.json")),
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "must provide 2 spirits\n",
		},
		{
			name:           "3 spirits",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/too-many-spirits.json")),
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "must provide 2 spirits\n",
		},
		{
			name:           "not found",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/nope", readFile(t, "testdata/good-spirits.json")),
			wantStatusCode: http.StatusNotFound,
		},
		{
			name:           "method not allowed",
			req:            newRequest(t, http.MethodPut, baseURL+"/api/battle", readFile(t, "testdata/good-spirits.json")),
			wantStatusCode: http.StatusMethodNotAllowed,
		},
		{
			name:           "invalid body",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", "42"),
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "cannot decode body: json: cannot unmarshal number into Go value of type []*api.Spirit\n",
		},
		{
			name:           "infinite loop",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/powerless-spirits.json")),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/powerless-spirits.txt"),
		},
		{
			name:           "unrecognized action",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/unrecognized-action.json")),
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "unrecognized action: \"tuna\"\n",
		},
		{
			name:           "unrecognized intelligence",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/unrecognized-intelligence.json")),
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "unrecognized intelligence: \"tuna\"\n",
		},
		{
			name:           "/battle bad seed",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle?seed=tuna", readFile(t, "testdata/good-spirits.json")),
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "invalid seed\n",
		},
		{
			name:           "/battle with human interaction requested",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/battle", readFile(t, "testdata/good-spirits-with-single-human-interaction.json")),
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "unsupported intelligence value (hint: you must use websocket API): \"human\"\n",
		},
		// /spirit happy paths
		{
			name:           "generated spirits with seed 1",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/spirit?seed=1", ""),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/generated-spirits-seed-1.json"),
		},
		{
			name:           "generated spirits with seed 2",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/spirit?seed=2", ""),
			wantStatusCode: http.StatusOK,
			wantBody:       readFile(t, "testdata/generated-spirits-seed-2.json"),
		},
		// /spirit sad paths
		{
			name:           "/spirit wrong method",
			req:            newRequest(t, http.MethodPut, baseURL+"/api/spirit?seed=2", ""),
			wantStatusCode: http.StatusMethodNotAllowed,
		},
		{
			name:           "/spirit bad seed",
			req:            newRequest(t, http.MethodPost, baseURL+"/api/spirit?seed=tuna", ""),
			wantStatusCode: http.StatusBadRequest,
			wantBody:       "invalid seed\n",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Logf("req: %s %s", test.req.Method, test.req.URL)
			rsp, err := http.DefaultClient.Do(test.req)
			require.NoError(t, err)

			gotBody, err := io.ReadAll(rsp.Body)
			require.Equalf(t, test.wantStatusCode, rsp.StatusCode, "body: %q", string(gotBody))
			require.NoError(t, err)
			require.Equal(t, test.wantBody, string(gotBody))
		})
	}

	t.Run("generated spirits are valid", func(t *testing.T) {
		for i := 0; i < 20; i++ {
			// Generate spirits.
			req := newRequest(t, http.MethodPost, baseURL+"/api/spirit", "")
			rsp, err := http.DefaultClient.Do(req)
			require.NoError(t, err)

			gotBody, err := io.ReadAll(rsp.Body)
			require.Equalf(t, http.StatusOK, rsp.StatusCode, "body: %q", string(gotBody))
			require.NoError(t, err)

			// Make sure spirits are valid.
			req = newRequest(t, http.MethodPost, baseURL+"/api/battle", string(gotBody))
			rsp, err = http.DefaultClient.Do(req)
			require.NoError(t, err)

			gotBody, err = io.ReadAll(rsp.Body)
			require.Equalf(t, http.StatusOK, rsp.StatusCode, "body: %q", string(gotBody))
			require.NoError(t, err)
		}
	})
}

func newRequest(t *testing.T, method, url string, body string) *http.Request {
	buf := bytes.NewBuffer([]byte(body))

	req, err := http.NewRequest(method, url, buf)
	require.NoError(t, err)

	return req
}

type testMsg struct {
	in    bool // true: tx, false: rx
	reset bool // true to reset connection
	msg   interface{}
}

func testWebsocketAPI(t *testing.T, baseURL string) {
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

	steps := []struct {
		name string
		msgs []testMsg
	}{
		{
			name: "when the error type is invalid it doesn't take down the server",
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
			name: "battle-start with 2 spirits without human interaction",
			msgs: []testMsg{
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeBattleStart,
						Details: &api.MessageDetailsBattleStart{
							Spirits: readSpirits(t, "testdata/good-spirits-without-human-interaction.json"),
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeBattleStop,
						Details: &api.MessageDetailsBattleStop{
							Output: readFile(t, "testdata/good-spirits-without-human-interaction.txt"),
						},
					},
				},
			},
		},
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
			name: "infinite loop",
			msgs: []testMsg{
				{
					in: true,
					msg: api.Message{
						Type: api.MessageTypeBattleStart,
						Details: &api.MessageDetailsBattleStart{
							Spirits: readSpirits(t, "testdata/powerless-spirits.json"),
						},
					},
				},
				{
					msg: api.Message{
						Type: api.MessageTypeBattleStop,
						Details: &api.MessageDetailsBattleStop{
							Output: readFile(t, "testdata/powerless-spirits.txt"),
						},
					},
				},
			},
		},
	}
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

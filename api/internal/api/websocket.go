package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

type MessageType string

const (
	MessageTypeBattleStart MessageType = "battle-start"
	MessageTypeBattleStop  MessageType = "battle-stop"
	MessageTypeActionReq   MessageType = "action-req"
	MessageTypeActionRsp   MessageType = "action-rsp"
	MessageTypeSpiritReq   MessageType = "spirit-req"
	MessageTypeSpiritRsp   MessageType = "spirit-rsp"
	MessageTypeError       MessageType = "error"
)

type MessageDetails interface{}

type Message struct {
	Type    MessageType    `json:"type"`
	Details MessageDetails `json:"details"`
}

func (m *Message) String() string {
	return fmt.Sprintf("api.Message{Type:%q, Details:%i}", m.Type, m.Details)
}

func (m *Message) UnmarshalJSON(b []byte) error {
	mTmp := struct {
		Type    MessageType            `json:"type"`
		Details map[string]interface{} `json:"details"`
	}{}

	if err := json.Unmarshal(b, &mTmp); err != nil {
		return fmt.Errorf("invalid message base: %w", err)
	}

	m.Type = mTmp.Type

	details := []struct {
		t MessageType
		d interface{}
	}{
		{t: MessageTypeBattleStart, d: &MessageDetailsBattleStart{}},
		{t: MessageTypeBattleStop, d: &MessageDetailsBattleStop{}},
		{t: MessageTypeActionReq, d: &MessageDetailsActionReq{}},
		{t: MessageTypeActionRsp, d: &MessageDetailsActionRsp{}},
		{t: MessageTypeSpiritReq, d: &MessageDetailsSpiritReq{}},
		{t: MessageTypeSpiritRsp, d: &MessageDetailsSpiritRsp{}},
		{t: MessageTypeError, d: &MessageDetailsError{}},
	}
	for _, detail := range details {
		if m.Type == detail.t {
			if err := mapstructure.Decode(mTmp.Details, detail.d); err != nil {
				return fmt.Errorf("invalid message details: %w", err)
			}

			m.Details = detail.d

			return nil
		}
	}

	return fmt.Errorf("unrecognized message type: %q", m.Type)
}

type MessageDetailsBattleStart struct {
	Spirits []*Spirit `json:"spirits" mapstructure:"spirits"`
}

type MessageDetailsBattleStop struct {
	Output string `json:"output" mapstructure:"output"`
}

type MessageDetailsActionReq struct {
	Spirit Spirit `json:"spirit" mapstructure:"spirit"`
	Output string `json:"output" mapstructure:"output"`
}

type MessageDetailsActionRsp struct {
	Spirit Spirit `json:"spirit" mapstructure:"spirit"`
	ID     string `json:"id" mapstructure:"id"`
}

type MessageDetailsSpiritReq struct {
}

type MessageDetailsSpiritRsp struct {
	Spirits []*Spirit `json:"spirits" mapstructure:"spirits"`
}

type MessageDetailsError struct {
	Reason string `json:"reason" mapstruture:"reason"`
}

func serveWebsocket(w http.ResponseWriter, r *http.Request) {
	log.Print("serving websocket")

	seed, ok := getSeed(r)
	if !ok {
		http.Error(w, "invalid seed", http.StatusBadRequest)
		return
	}

	upgrader := websocket.Upgrader{
		Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
			log.Printf("could not upgrade connection: %d %s", status, reason.Error())
		},
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("origin")
			if len(origin) == 0 {
				log.Printf("origin is empty, allowing")
				return true
			}

			originURL, err := url.Parse(origin)
			if err != nil {
				log.Printf("failed parsing origin url (%s), disallowing", err.Error())
				return false
			}

			if len(r.Host) == 0 {
				log.Printf("host is empty, disallowing")
				return false
			}

			// net.SplitHostPort() doesn't allow host[:port]...
			// This logic surely doesn't work for all allowable values...
			originSplit := strings.Split(originURL.Host, ":")
			hostSplit := strings.Split(r.Host, ":")
			if (originSplit[0] == "127.0.0.1" && hostSplit[0] == "127.0.0.1") ||
				(originSplit[0] == "localhost" && hostSplit[0] == "localhost") {
				log.Printf("origin and host are loopback, allowing")
				return true
			}

			if !strings.EqualFold(originURL.Host, r.Host) {
				log.Printf("origin (%s) and host (%s) mismatch, disallowing", origin, r.Host)
				return false
			}

			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "could not upgrade connection: "+err.Error(), http.StatusBadRequest)
		return
	}
	log.Print("upgraded connection")
	defer func() {
		log.Printf("closing connection")
		conn.Close()
	}()

	// Set a really long timeout, just in case...
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	log.Print("starting event handler")
	newEventHandler(conn, seed).run(ctx)
}

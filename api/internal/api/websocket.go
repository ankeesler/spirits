package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

type MessageType string

const (
	MessageTypeBattleStart MessageType = "battle-start"
	MessageTypeBattleStop              = "battle-stop"
	MessageTypeActionReq               = "action-req"
	MessageTypeActionRsp               = "action-rsp"
	MessageTypeSpiritReq               = "spirit-req"
	MessageTypeSpiritRsp               = "spirit-rsp"
	MessageTypeError                   = "error"
)

type MessageDetails interface{}

type Message struct {
	Type    MessageType    `json:"type"`
	Details MessageDetails `json:"details"`
}

func (m *Message) UnmarshalJSON(b []byte) error {
	mTmp := struct {
		Type    MessageType            `json:"type"`
		Details map[string]interface{} `json:"details"`
	}{}

	if err := json.Unmarshal(b, &mTmp); err != nil {
		return &messageInvalidError{error: fmt.Errorf("invalid message base: %w", err)}
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
				return &messageInvalidError{error: fmt.Errorf("invalid message details: %w", err)}
			}

			m.Details = detail.d

			return nil
		}
	}

	return &messageInvalidError{error: fmt.Errorf("invalid message type: %q", m.Type)}
}

type messageInvalidError struct {
	error
}

func (e *messageInvalidError) Is(target error) bool {
	_, ok := target.(*messageInvalidError)
	return ok
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
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "could not upgrade connection: "+err.Error(), http.StatusBadRequest)
		return
	}
	log.Print("upgraded connection")

	// Set a really long timeout, just in case...
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)

	go func() {
		defer conn.Close()
		<-ctx.Done()
	}()

	inMsgCh := make(chan *Message)
	go func() {
		defer cancel()
		defer close(inMsgCh)

		for {
			var msg Message
			if err := conn.ReadJSON(&msg); err != nil {
				log.Printf("error receiving message: %s %#v", err.Error(), err)
				if errors.Is(err, &messageInvalidError{}) {
					e := &messageInvalidError{}
					errors.As(err, &e)
					log.Printf("received an invalid message: %s", e.Error())
					continue
				}
				return
			}
			inMsgCh <- &msg
		}
	}()

	outMsgCh := make(chan *Message)
	go func() {
		defer cancel()
		defer close(outMsgCh)

		for outMsg := range outMsgCh {
			if err := conn.WriteJSON(outMsg); err != nil {
				log.Printf("error sending message: %s", err.Error())
				return
			}
		}
	}()

	go newEventHandler(inMsgCh, outMsgCh, seed).run(ctx)

	log.Print("started event handler")
}

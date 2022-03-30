package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"

	"github.com/ankeesler/spirits/api/internal/battle"
	"github.com/ankeesler/spirits/api/internal/spirit"
	"github.com/ankeesler/spirits/api/internal/spirit/action"
	"github.com/ankeesler/spirits/api/internal/spirit/generate"
	"github.com/ankeesler/spirits/api/internal/ui"
	"github.com/gorilla/websocket"
)

type conn interface {
	ReadMessage() (int, []byte, error)
	WriteMessage(int, []byte) error
}

type battleRunner struct {
	mu sync.Mutex

	s      atomic.Value // []*spirit.Spirit
	o      syncBuffer
	cancel func()
}

// start starts a battle. start will panic if a battle is already running.
func (b *battleRunner) start(ctx context.Context, spirits []*spirit.Spirit) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.runningLocked() {
		panic("battle already running")
	}

	var battleCtx context.Context
	battleCtx, b.cancel = context.WithCancel(ctx)
	go func() {
		defer func() {
			b.mu.Lock()
			defer b.mu.Unlock()

			b.cancel()
			b.cancel = nil
		}()

		b.o.reset()
		u := ui.New(&b.o)
		battle.Run(battleCtx, spirits, func(spirits []*spirit.Spirit, err error) {
			u.OnSpirits(spirits, err)
			b.s.Store(spirits)
		})
	}()
}

// stop stops a battle that is currently running. stop will panic if a battle is not running.
func (b *battleRunner) stop() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if !b.runningLocked() {
		panic("battle not running")
	}

	b.cancel()
}

// running returns whether a battle is currently being run.
func (b *battleRunner) running() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.runningLocked()
}

func (b *battleRunner) runningLocked() bool {
	return b.cancel != nil
}

// output returns the output up to this point. output will panic if a battle is not running.
func (b *battleRunner) output() string {
	return b.o.read()
}

// spirits returns the spirits that are currently involved in the battle. spirits will panic if a battle is not running.
func (b *battleRunner) spirits() []*spirit.Spirit {
	if !b.running() {
		panic("battle not running")
	}
	return b.s.Load().([]*spirit.Spirit)
}

type eventHandler struct {
	conn conn
	r    *rand.Rand

	wantActionResponse atomic.Value // *sync.Once
	actionResponse     chan *MessageDetailsActionRsp

	inBattle     atomic.Value // bool
	battleDone   chan struct{}
	battleOutput syncBuffer
}

func newEventHandler(conn conn, seed int64) *eventHandler {
	return &eventHandler{
		conn: conn,
		r:    rand.New(rand.NewSource(seed)),
	}
}

func (e *eventHandler) run(ctx context.Context) {
	e.actionResponse = make(chan *MessageDetailsActionRsp)
	defer close(e.actionResponse)
	e.battleDone = make(chan struct{})
	defer close(e.battleDone)

	alreadyUsedOnce := sync.Once{}
	alreadyUsedOnce.Do(func() {})
	e.wantActionResponse.Store(&alreadyUsedOnce)

	e.inBattle.Store(false)

	for {
		select {
		case <-ctx.Done():
			log.Printf("context closed: %s", ctx.Err().Error())
			return
		default:
		}

		_, data, err := e.conn.ReadMessage()
		if err != nil {
			log.Printf("error reading message: %s", err.Error())
			return
		}

		var m Message
		if err := json.Unmarshal(data, &m); err != nil {
			e.sendError(err.Error())
			continue
		}

		e.process(&m)
	}
}

func (e *eventHandler) process(msg *Message) {
	log.Printf("processing message: %#v", msg)

	switch msg.Type {
	case MessageTypeBattleStart:
		details := msg.Details.(*MessageDetailsBattleStart)
		e.onBattleStart(details)
	case MessageTypeBattleStop:
		details := msg.Details.(*MessageDetailsBattleStop)
		e.onBattleStop(details)
	case MessageTypeActionReq:
		details := msg.Details.(*MessageDetailsActionReq)
		e.onActionReq(details)
	case MessageTypeActionRsp:
		details := msg.Details.(*MessageDetailsActionRsp)
		e.onActionRsp(details)
	case MessageTypeSpiritReq:
		details := msg.Details.(*MessageDetailsSpiritReq)
		e.onSpiritReq(details)
	case MessageTypeSpiritRsp:
		details := msg.Details.(*MessageDetailsSpiritRsp)
		e.onSpiritRsp(details)
	default:
		e.sendError(fmt.Sprintf("unrecognized event type: %q", msg.Type))
	}
}

func (e *eventHandler) onBattleStart(details *MessageDetailsBattleStart) {
	if e.inBattle.Load().(bool) {
		e.sendError("battle already running")
		return
	}

	if len(details.Spirits) != 2 {
		e.sendError("must provide 2 spirits")
		return
	}

	internalSpirits, err := toInternalSpirits(details.Spirits, e.r, e.getAction)
	if err != nil {
		e.sendError(err.Error())
		return
	}

	e.inBattle.Store(true)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		defer e.inBattle.Store(false)

		select {
		case <-e.battleDone:
		case <-ctx.Done():
		}
	}()

	go func() {
		defer cancel()
		defer e.battleOutput.reset()

		u := ui.New(&e.battleOutput)
		battle.Run(ctx, internalSpirits, u.OnSpirits)
		e.send(&Message{
			Type: MessageTypeBattleStop,
			Details: &MessageDetailsBattleStop{
				Output: e.battleOutput.read(),
			},
		})
	}()
}

func (e *eventHandler) onBattleStop(details *MessageDetailsBattleStop) {
	if !e.inBattle.Load().(bool) {
		e.sendError("unexpected battle-stop: no battle running")
		return
	}

	e.battleDone <- struct{}{}
}

func (e *eventHandler) onActionReq(details *MessageDetailsActionReq) {
	e.sendError(fmt.Sprintf("unexpected action-req for spirit: %q", details.Spirit.Name))
}

func (e *eventHandler) onActionRsp(details *MessageDetailsActionRsp) {
	var handled bool
	e.wantActionResponse.Load().(*sync.Once).Do(func() {
		e.actionResponse <- details
		handled = true
	})

	if !handled {
		e.sendError(fmt.Sprintf("unexpected action-rsp with ID %q for spirit: %q", details.ID, details.Spirit.Name))
	}
}

func (e *eventHandler) getAction(ctx context.Context, s *Spirit) (spirit.Action, error) {
	e.wantActionResponse.Store(&sync.Once{})

	// Send the request.
	e.send(&Message{
		Type: MessageTypeActionReq,
		Details: MessageDetailsActionReq{
			Spirit: *s,
			Output: e.battleOutput.read(),
		},
	})

	// Wait for the response.
	var actionResponse *MessageDetailsActionRsp
	var ok bool
	select {
	case actionResponse, ok = <-e.actionResponse:
		if !ok {
			return nil, errors.New("never got action response from client")
		}
	case <-ctx.Done():
		return nil, fmt.Errorf("action canceled: %s", ctx.Err().Error())
	}

	// Make sure the action ID is valid.
	var actionID string
	spiritActions := s.Actions
	if len(spiritActions) == 0 {
		spiritActions = []string{"attack"}
	}
	if actionResponse.ID == "" {
		actionID = spiritActions[0]
	} else {
		for _, action := range spiritActions {
			if action == actionResponse.ID {
				actionID = action
				break
			}
		}
	}

	if len(actionID) == 0 {
		return nil, fmt.Errorf("unknown action %q for spirit %q", actionResponse.ID, s.Name)
	}

	log.Printf("running action %q for spirit %q", actionID, s.Name)
	return toInternalAction([]string{actionID}, "", e.r, nil)
}

func (e *eventHandler) onSpiritReq(details *MessageDetailsSpiritReq) {
	wellKnownActions := []spirit.Action{
		&actions{
			ids:    []string{"attack"},
			Action: action.Attack(),
		},
		&actions{
			ids:    []string{"bolster"},
			Action: action.Bolster(),
		},
		&actions{
			ids:    []string{"drain"},
			Action: action.Drain(),
		},
		&actions{
			ids:    []string{"charge"},
			Action: action.Charge(),
		},
	}
	internalSpirits := generate.Generate(e.r, wellKnownActions, func(generatedActions []spirit.Action) spirit.Action {
		var ids []string
		for _, generatedAction := range generatedActions {
			ids = append(ids, generatedAction.(*actions).ids...)
		}
		return &actions{
			ids:    ids,
			Action: action.RoundRobin(generatedActions),
		}
	})
	apiSpirits, err := fromInternalSpirits(internalSpirits)
	if err != nil {
		e.sendError("could not generate spirits: " + err.Error())
		return
	}

	e.send(&Message{
		Type: MessageTypeSpiritRsp,
		Details: &MessageDetailsSpiritRsp{
			Spirits: apiSpirits,
		},
	})
}

func (e *eventHandler) onSpiritRsp(details *MessageDetailsSpiritRsp) {
	e.sendError("unexpected spirit-rsp")
}

func (e *eventHandler) sendError(reason string) {
	e.send(&Message{
		Type: MessageTypeError,
		Details: &MessageDetailsError{
			Reason: reason,
		},
	})
}

func (e *eventHandler) send(m *Message) {
	data, err := json.Marshal(m)
	if err != nil {
		log.Printf("marshal message failed: %s", err.Error())
		return
	}

	if err := e.conn.WriteMessage(websocket.TextMessage, data); err != nil {
		log.Printf("write message failed: %s", err.Error())
		return
	}
}

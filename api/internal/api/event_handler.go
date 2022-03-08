package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"sync/atomic"

	"github.com/ankeesler/spirits/api/internal/battle"
	"github.com/ankeesler/spirits/api/internal/spirit"
	"github.com/ankeesler/spirits/api/internal/spirit/action"
	"github.com/ankeesler/spirits/api/internal/spirit/generate"
	"github.com/ankeesler/spirits/api/internal/ui"
)

type eventHandler struct {
	inMsgCh  <-chan *Message
	outMsgCh chan<- *Message

	seed int64

	wantActionResponse atomic.Value // *sync.Once
	actionResponse     chan *MessageDetailsActionRsp

	inBattle     atomic.Value // bool
	battleDone   chan struct{}
	battleOutput syncBuffer
}

func newEventHandler(inMsgCh <-chan *Message, outMsgCh chan<- *Message, seed int64) *eventHandler {
	return &eventHandler{
		inMsgCh:  inMsgCh,
		outMsgCh: outMsgCh,

		seed: seed,
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

	keepGoing := true
	for keepGoing {
		select {
		case inMsg, ok := <-e.inMsgCh:
			if !ok {
				keepGoing = false
				break
			}
			e.process(inMsg)
		case <-ctx.Done():
			keepGoing = false
		}
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
		e.outMsgCh <- newErrorMsg(fmt.Sprintf("unrecognized event type: %q", msg.Type))
	}
}

func (e *eventHandler) onBattleStart(details *MessageDetailsBattleStart) {
	if e.inBattle.Load().(bool) {
		e.outMsgCh <- newErrorMsg("battle already running")
		return
	}

	if len(details.Spirits) != 2 {
		e.outMsgCh <- newErrorMsg("must provide 2 spirits")
		return
	}

	internalSpirits, err := toInternalSpirits(details.Spirits, e.seed, e.getAction)
	if err != nil {
		e.outMsgCh <- newErrorMsg(err.Error())
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
		e.outMsgCh <- &Message{
			Type: MessageTypeBattleStop,
			Details: &MessageDetailsBattleStop{
				Output: e.battleOutput.read(),
			},
		}
	}()
}

func (e *eventHandler) onBattleStop(details *MessageDetailsBattleStop) {
	if !e.inBattle.Load().(bool) {
		e.outMsgCh <- newErrorMsg("unexpected battle-stop: no battle running")
		return
	}

	e.battleDone <- struct{}{}
}

func (e *eventHandler) onActionReq(details *MessageDetailsActionReq) {
	e.outMsgCh <- newErrorMsg(fmt.Sprintf("unexpected action-req for spirit: %q", details.Spirit.Name))
}

func (e *eventHandler) onActionRsp(details *MessageDetailsActionRsp) {
	var handled bool
	e.wantActionResponse.Load().(*sync.Once).Do(func() {
		e.actionResponse <- details
		handled = true
	})

	if !handled {
		e.outMsgCh <- newErrorMsg(fmt.Sprintf("unexpected action-rsp with ID %q for spirit: %q", details.ID, details.Spirit.Name))
	}
}

func (e *eventHandler) getAction(ctx context.Context, s *Spirit) (spirit.Action, error) {
	e.wantActionResponse.Store(&sync.Once{})

	// Send the request.
	e.outMsgCh <- &Message{
		Type: MessageTypeActionReq,
		Details: MessageDetailsActionReq{
			Spirit: *s,
			Output: e.battleOutput.read(),
		},
	}

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
	return toInternalAction([]string{actionID}, "", e.seed, nil)
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
	internalSpirits := generate.Generate(e.seed, wellKnownActions, func(generatedActions []spirit.Action) spirit.Action {
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
		e.outMsgCh <- newErrorMsg("could not generate spirits: " + err.Error())
		return
	}

	e.outMsgCh <- &Message{
		Type: MessageTypeSpiritRsp,
		Details: &MessageDetailsSpiritRsp{
			Spirits: apiSpirits,
		},
	}
}

func (e *eventHandler) onSpiritRsp(details *MessageDetailsSpiritRsp) {
	e.outMsgCh <- newErrorMsg("unexpected spirit-rsp")
}

func newErrorMsg(reason string) *Message {
	return &Message{
		Type: MessageTypeError,
		Details: &MessageDetailsError{
			Reason: reason,
		},
	}
}

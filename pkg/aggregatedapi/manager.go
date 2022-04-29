package aggregatedapi

import "context"

type ActionSink interface {
	Post(battleName, battleGeneration, spiritName, spiritGeneration, actionName string) error
}

type Manager struct {
	ActionSink ActionSink
}

func (m *Manager) Start(ctx context.Context) {

}

package spirit

import (
	"context"
	"fmt"

	"github.com/ankeesler/spirits0/internal/api"
	"google.golang.org/protobuf/proto"
)

type Action interface {
	Run(context.Context, *Spirit, []*Spirit)
}

type Spirit struct {
	API *api.Spirit

	Action Action
}

func FromAPI(apiSpirit *api.Spirit) (*Spirit, error) {
	internalSpirit := Spirit{
		API: proto.Clone(apiSpirit).(*api.Spirit),
	}

	var err error
	internalSpirit.Action, err = actionFromAPI(apiSpirit.GetActions())
	if err != nil {
		return nil, fmt.Errorf("action from api: %w", err)
	}

	return &internalSpirit, nil
}

func actionFromAPI(apiActions []*api.SpiritAction) (Action, error) {
	return nil, nil
}

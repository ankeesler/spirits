package service

import (
	"context"
	"net/http"

	"github.com/ankeesler/spirits/pkg/api"
	server "github.com/ankeesler/spirits/pkg/api/generated/server/api"
)

type Default struct {
}

func NewDefault() *Default {
	return &Default{}
}

func (d *Default) RootGet(ctx context.Context) (server.ImplResponse, error) {
	return server.ImplResponse{
		Code: http.StatusOK,
		Body: api.Object,
	}, nil
}

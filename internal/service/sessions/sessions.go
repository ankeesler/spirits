package sessions

import server "github.com/ankeesler/spirits/pkg/api/generated/server/api"

type Service struct {
	server.SessionsApiServicer
}

func New() *Service {
	return &Service{
		SessionsApiServicer: server.NewSessionsApiService(),
	}
}

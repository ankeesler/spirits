package sessions

import server "github.com/ankeesler/spirits/pkg/api/generated/server/api"

type Service struct {
	server.SessionApiServicer
}

func New() *Service {
	return &Service{
		SessionApiServicer: server.NewSessionApiService(),
	}
}

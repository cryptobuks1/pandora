package api

import (
	"pandora/pkg/cfg"
	"pandora/pkg/utils/pndrhttp"
)

type Service struct {
	healthController *HealthController

	httpServer *pndrhttp.Server
}

func NewService() *Service {
	s := &Service{}

	s.healthController = NewHealthController()

	httpServer := pndrhttp.NewServer(cfg.Cfg.API.Port, s.routes())
	s.httpServer = httpServer

	return s
}

func (s *Service) Start() error {
	return s.httpServer.Start()
}

func (s *Service) Stop() {
	s.httpServer.Stop()
}

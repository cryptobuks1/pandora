package api

import (
	"pandora/pkg/api/utils"
	"pandora/pkg/cfg"
)

type Server struct {
	healthController *HealthController

	httpSrv *utils.Server
}

func NewServer() *Server {
	srv := &Server{}

	srv.healthController = NewHealthController()

	srv.httpSrv = utils.NewServer(cfg.Cfg.API.Port, srv.routes())

	return srv
}

func (srv *Server) Start() error {
	return srv.httpSrv.Start()
}

func (srv *Server) Stop() {
	srv.httpSrv.Stop()
}

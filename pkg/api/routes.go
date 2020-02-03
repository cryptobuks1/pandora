package api

import (
	"net/http"

	"pandora/pkg/api/utils"
)

func (srv *Server) routes() utils.Routes {
	return utils.Routes{
		{Path: "/live", Method: http.MethodGet, Handler: srv.healthController.LiveH},
	}
}

package api

import (
	"net/http"
)

func (srv *Server) routes() routes {
	return routes{
		{Path: "/health", Method: http.MethodGet, Handler: srv.healthController.HealthH},
	}
}

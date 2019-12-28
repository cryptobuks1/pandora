package api

import (
	"net/http"

	"pandora/pkg/utils/pndrhttp"
)

func (s *Service) routes() pndrhttp.Routes {
	return pndrhttp.Routes{
		{Path: "/live", Method: http.MethodGet, Handler: s.healthController.LiveH},
	}
}

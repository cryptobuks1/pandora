package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"pandora/pkg/config"
	"pandora/pkg/constants"
)

type Server struct {
	healthController *HealthController

	httpSrv *http.Server

	logger zerolog.Logger
}

type route struct {
	Path    string
	Method  string
	Handler func(http.ResponseWriter, *http.Request, httprouter.Params)
}
type routes []route

func NewServer() *Server {
	srv := &Server{
		logger: log.Logger.With().Str(constants.LoggerComponentKey, "api").Logger(),
	}

	endpoint := fmt.Sprintf("127.0.0.1:%d", config.Cfg.API.Port)

	router := httprouter.New()

	for _, route := range srv.routes() {
		router.Handle(route.Method, route.Path, route.Handler)
	}

	var h http.Handler
	h = handlers.LoggingHandler(os.Stdout, router)
	h = handlers.CORS(handlers.AllowedOrigins([]string{endpoint}))(h)

	srv.httpSrv = &http.Server{
		Handler: h,
		Addr:    endpoint,
	}

	srv.healthController = NewHealthController()

	return srv
}

func (srv *Server) Start() error {
	srv.logger.Info().Msgf("start http server on: %s", srv.httpSrv.Addr)
	return srv.httpSrv.ListenAndServe()
}

func (srv *Server) Stop() {
	srv.logger.Debug().Msg("shutdown http server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	cancel()
	if err := srv.httpSrv.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("shutdown http server error")
	}
}

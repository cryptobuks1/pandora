package pndrhttp

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
)

type Route struct {
	Path    string
	Method  string
	Handler func(http.ResponseWriter, *http.Request, httprouter.Params)
}
type Routes []Route

type Server struct {
	srv *http.Server
}

func NewServer(port int, routes Routes) *Server {
	endpoint := fmt.Sprintf("127.0.0.1:%d", port)

	router := httprouter.New()

	for _, route := range routes {
		router.Handle(route.Method, route.Path, route.Handler)
	}

	var h http.Handler
	h = handlers.LoggingHandler(os.Stdout, router)
	h = handlers.CORS(handlers.AllowedOrigins([]string{endpoint}))(h)

	srv := &http.Server{
		Handler: h,
		Addr:    endpoint,
	}

	return &Server{srv: srv}
}

func (s *Server) Start() error {
	log.Debug().Msgf("listen http server on %s", s.srv.Addr)
	return s.srv.ListenAndServe()
}

func (s *Server) Stop() {
	if err := s.srv.Shutdown(context.TODO()); err != nil {
		log.Err(err).Msg("shutdown http server error")
	}
}

package httpserver

import (
	"context"
	"net"
	"net/http"
	"strconv"
)

type Server struct {
	srv *http.Server
	cfg Config
}

func NewServer(cfg Config, handler http.Handler) *Server {
	srv := &http.Server{
		Addr:         net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)),
		Handler:      handler,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}
	return &Server{
		srv: srv,
		cfg: cfg,
	}
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, s.cfg.ShutdownTimeout)
	defer cancel()

	return s.srv.Shutdown(ctx)
}

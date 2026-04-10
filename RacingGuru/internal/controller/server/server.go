package server

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	server *http.Server
	logger *log.Logger
}

func NewServer(addr string, handler http.HandlerFunc, logger *log.Logger) *Server {
	return &Server{server: &http.Server{Addr: addr, Handler: handler}, logger: logger}
}

func (s *Server) Run() error {
	go func() {
		s.logger.Printf("Server starting on %s", s.server.Addr)
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Fatalf("Server failed to start: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	return nil
}

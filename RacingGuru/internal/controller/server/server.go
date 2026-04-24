package server

import (
	"RacingGuru/internal/logger"
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	server *http.Server
	logger *logger.Logger
}

func NewServer(addr string, handler http.Handler, logger *logger.Logger) *Server {
	return &Server{
		server: &http.Server{
			Addr:     addr,
			Handler:  handler,
			ErrorLog: logger.StdLogger(),
		},
		logger: logger,
	}
}

func (s *Server) Run() error {
	serverErr := make(chan error, 1)

	go func() {
		s.logger.Infof("server starting on %s", s.server.Addr)
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		s.logger.Errorf("server failed to start: %v", err)
		return err
	case sig := <-quit:
		s.logger.Infof("shutdown signal received: %s", sig)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Errorf("server shutdown failed: %v", err)
		return err
	}

	s.logger.Info("server stopped")
	return nil
}

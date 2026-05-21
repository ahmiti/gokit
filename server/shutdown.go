package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (s *Server) WaitForShutdown(drainPeriod time.Duration) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	s.logger.Info("received shutdown signal")

	ctx, cancel := context.WithTimeout(context.Background(), drainPeriod)
	defer cancel()

	if err := s.Stop(ctx); err != nil {
		s.logger.Error("error during shutdown", "error", err)
	}
}

package server

import (
	"context"
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Config struct {
	Addr         string        `yaml:"addr"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

type Server struct {
	app    *fiber.App
	cfg    Config
	logger *slog.Logger
}

func New(cfg Config, logger *slog.Logger) *Server {
	app := fiber.New(fiber.Config{
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	})

	app.Use(recover.New())

	srv := &Server{
		app:    app,
		cfg:    cfg,
		logger: logger,
	}

	srv.setupHealth()

	return srv
}

func (s *Server) setupHealth() {
	s.app.Get("/health/live", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "alive"})
	})

	s.app.Get("/health/ready", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ready"})
	})
}

func (s *Server) App() *fiber.App {
	return s.app
}

func (s *Server) Start() error {
	s.logger.Info("starting server", "addr", s.cfg.Addr)
	return s.app.Listen(s.cfg.Addr)
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Info("shutting down server")
	return s.app.ShutdownWithContext(ctx)
}

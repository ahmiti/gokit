package middleware

import (
	"log/slog"
	"time"

	"github.com/ahmiti/gokit/logger"
	"github.com/gofiber/fiber/v2"
)

func Logger(log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start)

		requestID, _ := c.Locals("request_id").(string)
		ctx := logger.WithLogger(c.UserContext(), log)
		ctx = logger.With(ctx, "request_id", requestID)
		c.SetUserContext(ctx)

		log.Info("request",
			"method", c.Method(),
			"path", c.Path(),
			"status", c.Response().StatusCode(),
			"duration_ms", duration.Milliseconds(),
			"request_id", requestID,
		)
		return err
	}
}

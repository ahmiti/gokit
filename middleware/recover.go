package middleware

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func Recover(log *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				log.Error("panic recovered", "panic", r)
				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": fiber.Map{
						"code":    "internal_server_error",
						"message": "internal server error",
					},
				})
			}
		}()
		return c.Next()
	}
}

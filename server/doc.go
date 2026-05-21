// Package server provides a pre-configured HTTP server with:
//   - Standard middleware stack (recover, request ID, logger, CORS, metrics)
//   - Health endpoints (/health/live, /health/ready)
//   - Prometheus metrics (/metrics)
//   - Graceful shutdown
//
// Example:
//
//   cfg := server.Config{
//       Addr:         ":8080",
//       ReadTimeout:  15 * time.Second,
//       WriteTimeout: 30 * time.Second,
//   }
//   srv := server.New(cfg, logger)
//   srv.App().Get("/hello", func(c *fiber.Ctx) error {
//       return c.SendString("world")
//   })
//   srv.Start()
package server

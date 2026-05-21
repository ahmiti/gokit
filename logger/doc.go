// Package logger provides a structured logger wrapper around slog.
//
// It enforces the company's logging standards:
//   - JSON format in production, text in development
//   - UTC timestamps in RFC3339Nano
//   - Mandatory fields: time, level, service, env, msg, request_id, trace_id
//   - Request-scoped logger attached to context
//
// Example:
//
//	cfg := logger.Config{
//	    Level:       "info",
//	    Format:      "json",
//	    ServiceName: "orders-api",
//	    Env:         "prod",
//	}
//	log := logger.New(cfg)
//	log.Info("order created", "order_id", "ord_123")
package logger

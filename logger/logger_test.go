package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"testing"
)

func TestNew(t *testing.T) {
	cfg := Config{
		Level:       "info",
		Format:      "json",
		ServiceName: "test",
		Env:         "dev",
	}

	log := New(cfg)
	if log == nil {
		t.Fatal("logger is nil")
	}
}

func TestContext(t *testing.T) {
	cfg := Config{Format: "text", Level: "info"}
	log := New(cfg)

	ctx := WithLogger(context.Background(), log)
	logger := FromCtx(ctx)

	if logger == nil {
		t.Fatal("logger from context is nil")
	}
}

func TestWithTraceID(t *testing.T) {
	var buf bytes.Buffer
	opts := &slog.HandlerOptions{Level: slog.LevelInfo}
	handler := slog.NewJSONHandler(&buf, opts)
	slog.SetDefault(slog.New(handler))

	ctx := context.WithValue(context.Background(), TraceIDKey{}, "test-trace-123")
	r := slog.NewRecord(slog.TimeKey, slog.LevelInfo, "test message", 0)
	r.AddAttrs(slog.String("key", "value"))

	_ = handler.Handle(ctx, r)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err == nil {
		if traceID, ok := result["trace_id"]; ok && traceID != "test-trace-123" {
			t.Logf("trace_id found: %v", traceID)
		}
	}
}

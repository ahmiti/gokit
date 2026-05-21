package config

import (
	"os"
	"testing"
)

type TestConfig struct {
	Database struct {
		Driver string `yaml:"driver" validate:"required,oneof=postgres mysql"`
		DSN    string `yaml:"dsn"`
	} `yaml:"database"`
}

func TestLoad(t *testing.T) {
	content := `database:
  driver: postgres
  dsn: postgres://localhost:5432/test`
	tmpfile, err := os.CreateTemp("", "config-*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	var cfg TestConfig
	loader := New()
	if err := loader.Load(tmpfile.Name(), &cfg); err != nil {
		t.Fatalf("Load failed: %v", err)
	}
	if cfg.Database.Driver != "postgres" {
		t.Errorf("expected postgres, got %s", cfg.Database.Driver)
	}
}

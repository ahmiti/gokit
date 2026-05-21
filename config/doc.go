// Package config provides a typed configuration loader.
//
// Load order (highest priority last):
//   1. Struct defaults
//   2. config/config.yaml
//   3. config/config.local.yaml (gitignored)
//   4. Environment variables (DATABASE_DSN -> database.dsn)
//   5. Command-line flags
//
// Example:
//
//	type AppConfig struct {
//	    Database struct {
//	        Driver string `yaml:"driver" validate:"required,oneof=postgres mysql"`
//	        DSN    string `yaml:"dsn" validate:"required"`
//	    } `yaml:"database"`
//	}
//
//	var cfg AppConfig
//	loader := config.New()
//	if err := loader.Load("config/config.yaml", &cfg); err != nil {
//	    log.Fatal(err)
//	}
package config

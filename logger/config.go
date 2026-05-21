package logger

type Config struct {
	Level       string `yaml:"level"`        // debug, info, warn, error
	Format      string `yaml:"format"`       // json, text
	ServiceName string `yaml:"service_name"` // nom du service
	Env         string `yaml:"env"`          // dev, staging, prod
}

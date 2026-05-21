package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Loader struct {
	v *viper.Viper
}

func New() *Loader {
	v := viper.New()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	return &Loader{v: v}
}

func (l *Loader) Load(configFile string, cfg interface{}) error {
	l.v.SetConfigFile(configFile)
	if err := l.v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("read config: %w", err)
		}
	}
	if err := l.v.Unmarshal(cfg); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}
	return nil
}

func (l *Loader) Get(key string) interface{} {
	return l.v.Get(key)
}

func (l *Loader) SetDefault(key string, value interface{}) {
	l.v.SetDefault(key, value)
}

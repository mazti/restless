package config

import (
	"github.com/spf13/viper"

	"github.com/tiennv147/mazti-commons/config"
)

// Config stores configuration variables.
type Config struct {
	Name       *string
	Database   *config.Database
	HTTP       *config.HTTP
	GRPC       *GRPC
	Release    *config.Release
	SwaggerDir *string
}

// New returns a new config instance.
func New() (*Config, error) {
	cfg := &Config{}

	v := viper.New()
	v.SetConfigType("yml")
	v.SetConfigName("config")
	v.AddConfigPath("./")
	v.AddConfigPath("$GOPATH/src/github.com/mazti/restless/meta/config/.")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := v.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

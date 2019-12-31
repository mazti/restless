package config

import (
	"github.com/spf13/viper"

	"github.com/tiennv147/mazti-commons/config"
)

// Config stores configuration variables.
type Config struct {
	Database   *config.Database
	HTTP       *config.HTTP
	GRPC       *GRPC
	Release    *config.Release
	Meta       string
	SwaggerDir *string
}

// New returns a new config instance.
func New() (*Config, error) {
	cfg := &Config{}

	viper.New()
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.AddConfigPath("$GOPATH/src/github.com/tiennv147/restless/base/config/.")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

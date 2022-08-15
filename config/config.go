package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Host string `envconfig:"HOST" default:"localhost"`
	Port string `envconfig:"PORT" default:"8080"`
}

var settings Config

func Init() error {
	err := envconfig.Process("", &settings)
	if err != nil {
		return err
	}
	return nil
}

func GetSettings() *Config {
	return &settings
}

package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Host string `envconfig:"PARSER_HOST" default:"localhost"`
	Port string `envconfig:"PARSER_PORT" default:"8080"`
	// mongo
	MongoHost     string `envconfig:"MONGO_HOST" default:"localhost"`
	MongoPort     string `envconfig:"MONGO_PORT" default:"27017"`
	MongoDBName   string `envconfig:"MONGO_DB_NAME" default:"parser_test"`
	MongoUserName string `envconfig:"MONGO_USERNAME" default:"admin"`
	MongoPassword string `envconfig:"MONGO_PASSWORD" default:"admin"`
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

package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppEnv struct {
	ServiceName string `default:"vaultly" required:"true" split_words:"true"`
	PORT        int    `required:"true"`
	PostgresUrl string `required:"true" split_words:"true"`
}

func LoadEnv() (*AppEnv, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	env := &AppEnv{}

	if err := envconfig.Process("", env); err != nil {
		return nil, err
	}

	return env, nil
}

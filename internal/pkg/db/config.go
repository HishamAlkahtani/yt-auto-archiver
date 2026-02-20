package db

import "github.com/kelseyhightower/envconfig"

type config struct {
	User     string `envconfig:"POSTGRES_USER" required:"true"`
	Password string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	DB       string `envconfig:"POSTGRES_DB" required:"true"`
}

func newConfig() (*config, error) {
	var config config

	err := envconfig.Process("", &config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}

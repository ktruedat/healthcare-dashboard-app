package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

func NewConfig() (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig("../config.yml", &cfg); err != nil {
		return nil, errors.Wrap(err, "failed to read config")
	}

	return &cfg, nil
}

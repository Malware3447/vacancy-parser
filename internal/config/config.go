package config

import (
	"github.com/x3a-tech/configo"
)

type Config struct {
	App        configo.App      `yaml:"app" env-required:"true"`
	DatabasePg configo.Database `yaml:"postgres" env-required:"true"`
	Sentry     configo.Sentry   `yaml:"sentry"`
	Logger     configo.Logger   `yaml:"logger"`
}

func (c Config) Env() string {
	return c.App.Env
}

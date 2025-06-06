package config

import (
	"github.com/x3a-tech/configo"
	"vacancy-parser/internal/config/parser"
)

type Config struct {
	App        configo.App      `yaml:"app" env-required:"true"`
	DatabasePg configo.Database `yaml:"postgres" env-required:"true"`
	Sentry     configo.Sentry   `yaml:"sentry"`
	Logger     configo.Logger   `yaml:"logger"`
	Parser     parser.Parser    `yaml:"parser"`
}

func (c Config) Env() string {
	return c.App.Env
}

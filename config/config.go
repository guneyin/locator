package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
)

const (
	defaultPort = 8080
)

type Config struct {
	DbHost       string `env:"DB_HOST"`
	DbPort       int    `env:"DB_PORT"`
	DbUser       string `env:"DB_USER"`
	DbPassword   string `env:"DB_PASSWORD"`
	DbName       string `env:"DB_NAME"`
	Port         int    `env:"PORT"`
	MaxRateLimit int    `env:"MAX_RATE_LIMIT"`
}

func New() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		_ = cleanenv.ReadEnv(&cfg)
		slog.Info(err.Error())
	}

	err = cfg.validate()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *Config) validate() error {
	switch {
	case c.Port == 0:
		c.Port = defaultPort
	}

	return nil
}

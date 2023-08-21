package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Cfg struct {
		Port string `env:"PORT"`
		DB   DB
	}
	DB struct {
		Host     string `env:"DB_HOST"`
		Port     string `env:"DB_PORT"`
		Name     string `env:"DB_NAME"`
		Username string `env:"DB_USERNAME"`
		Password string `env:"DB_PASSWORD"`
	}
)

func New() *Cfg {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "dev" {
		return loadConfig("./config/dev.env")
	}

	return nil
}

func loadConfig(path string) *Cfg {
	var cfg Cfg

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}

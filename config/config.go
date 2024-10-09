package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	AppName    string `env:"APP_NAME" envDefault:"Test-Gemini-By-Elloy"`
	Env        string `env:"ENV" envDefault:"development"`
	ApiKey     string `env:"API_KEY" envDefault:"default"`
	Validation string `env:"VALIDATION" envDefault:"ozzo"`
	Port       string `env:"APP_PORT,notEmpty"`
	PrefixURL  string `env:"PREFIX_URL" envDefault:"/api"`
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"tag":   "Config.New.01",
			"error": err.Error(),
		}).Error("loading env file failed")
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		logrus.WithFields(logrus.Fields{
			"tag":   "Config.New.02",
			"error": err.Error(),
		}).Error("parsing env failed")
	}
	fmt.Println("cek en ====0", cfg)
	return &cfg
}

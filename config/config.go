package config

import (
	"os"

	"github.com/golobby/dotenv"
)

type AppConfig struct {
	DBURL       string `env:"DB_URL"`
	DBTimeout   int    `env:"DB_TIMEOUT"`
	HttpAddress string `env:"HTTP_ADDRESS"`
}

func MustLoad() *AppConfig {
	cfg := &AppConfig{}

	if file, err := os.Open(".env"); err == nil {
		defer file.Close()
		if err := dotenv.NewDecoder(file).Decode(cfg); err != nil {
			panic(err)
		}
	}

	return cfg
}

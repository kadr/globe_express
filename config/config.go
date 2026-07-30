package config

import (
	"os"

	"github.com/golobby/dotenv"
)

type AppConfig struct {
	DBURL               string `env:"DB_URL"`
	DBTimeout           int    `env:"DB_TIMEOUT"`
	HttpAddressTravaler string `env:"HTTP_ADDRESS_TRAVALER"`
	HttpAddressOrder    string `env:"HTTP_ADDRESS_ORDER"`

	JWTSecret string `env:"JWT_SECRET"`
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

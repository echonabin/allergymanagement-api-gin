package configs

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type dbEnvVariables struct {
	DbUser          string `env:"DB_USER_NAME"`
	DbHost          string `env:"DB_HOST"`
	DbPassword      string `env:"DB_PASSWORD"`
	DbName          string `env:"DB_NAME"`
	DbPort          string `env:"DB_PORT"`
	JwtAccessToken  string `env:"JWT_ACCESS_TOKEN_SECRET"`
	JwtRefreshToken string `env:"JWT_REFRESH_TOKEN_SECRET"`
}

var DbEnvConfig dbEnvVariables

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := env.Parse(&DbEnvConfig); err != nil {
		log.Fatal("Error parsing .env file")
	}
}

package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
)

func GetEnvVar(name string) string {
	env := os.Getenv("ENV")
	envMap, err := godotenv.Read(".env")
	if env == "test" {
		envMap, err = godotenv.Read("../../.env.test")
	}

	if err != nil {
		log.Debug().Msg("Error reading .env file")
		return ""
	}
	value, ok := envMap[name]
	if !ok {
		log.Debug().Msg("Error reading env var")
		return ""
	}
	return value
}

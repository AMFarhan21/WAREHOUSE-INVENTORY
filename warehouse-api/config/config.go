package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server           string
	ConnectionString string
	JwtSecret        string
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDb       string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Error on loading env")
	}

	cfg := &Config{
		Server:           os.Getenv("SERVER"),
		ConnectionString: os.Getenv("CONNECTION_STRING"),
		JwtSecret:        os.Getenv("JWT_SECRET"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDb:       os.Getenv("POSTGRES_DB"),
	}

	return cfg
}

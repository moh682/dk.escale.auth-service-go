package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Environments Type for environments
type Environments struct {
	PostgresqlUsername string
	PostgresqlPassword string
	PostgresqlDatabase string
	PostgresqlHost     string
	ServerPort         string
}

var env *Environments = nil

func setup() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir = dir + "/.env"

	err = godotenv.Load(dir)
	if err != nil {
		log.Fatalln(err)
	}

	env = &Environments{
		PostgresqlDatabase: os.Getenv("POSTGRES_DB"),
		PostgresqlPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresqlUsername: os.Getenv("POSTGRES_USER"),
		PostgresqlHost:     os.Getenv("POSTGRES_HOST"),
		ServerPort:         os.Getenv("SERVER_PORT"),
	}
}

// GetEnvironments Exports loaded environments
func GetEnvironments() *Environments {
	if env == nil {
		setup()
	}
	return env
}

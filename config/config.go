package config

import (
	"log"
	"os"
	"path"
	"runtime"

	"github.com/joho/godotenv"
)

// Environments Type for environments
type Environments struct {
	PostgresqlUsername string
	PostgresqlPassword string
	PostgresqlDatabase string
	PostgresqlHost     string
	ServerPort         string
	JWTSecretKey       string
	JWTIssuer          string
}

var env *Environments = nil

func init() {
	// environment := os.Getenv("ENV")
	var dir string
	_, filename, _, _ := runtime.Caller(0)
	// if environment == "Test" {
	dir = path.Join(path.Dir(filename), "..")
	// } else {
	// 	dir = path.Join(path.Dir(filename))
	// }
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func setup() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir = dir + "/.env"

	environment := os.Getenv("ENV")

	if environment != "test" && environment != "production" {
		err = godotenv.Load(dir)
		if err != nil {
			log.Fatalln(err)
		}
	}

	env = &Environments{
		PostgresqlDatabase: os.Getenv("POSTGRES_DB"),
		PostgresqlPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresqlUsername: os.Getenv("POSTGRES_USER"),
		PostgresqlHost:     os.Getenv("POSTGRES_HOST"),
		ServerPort:         os.Getenv("SERVER_PORT"),
		JWTSecretKey:       os.Getenv("JWT_SECRET_KEY"),
		JWTIssuer:          os.Getenv("JWT_ISSUER"),
	}
}

// GetEnvironments Exports loaded environments
func GetEnvironments() *Environments {
	if env == nil {
		setup()
	}
	return env
}

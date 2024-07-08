package infrastructure

import (
	"log"
	"os"
)

func fetchEnv(name string, strict bool) string {
	value := os.Getenv(name)
	if value == "" {
		if strict {
			log.Fatalf("Environment variable %s is not set", name)
		}
		log.Printf("WARN: Environment variable %s is not set\n", name)
	}
	return value
}

type DatabaseSettings struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func GetDatabaseSettings() DatabaseSettings {
	return DatabaseSettings{
		User:     fetchEnv("POSTGRES_USER", true),
		Password: fetchEnv("POSTGRES_PASSWORD", true),
		Database: fetchEnv("POSTGRES_DB", true),
		Host:     fetchEnv("DB_HOST", true),
		Port:     fetchEnv("DB_PORT", true),
	}
}

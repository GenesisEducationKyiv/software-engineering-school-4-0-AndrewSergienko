package infrastructure

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

func fetchEnv(name string, strict bool) string { // nolint: all
	value := os.Getenv(name)
	if value == "" {
		if strict {
			log.Fatalf("Environment variable %s is not set", name)
		}
		log.Printf("WARN: Environment variable %s is not set\n", name)
	}

	return value
}

type EmailSettings struct {
	Email    string
	Password string
	Host     string
}

type CurrencyRateServiceAPISettings struct {
	Host           string
	GetCurrencyURL string
}

type ServicesAPISettings struct {
	CurrencyRate *CurrencyRateServiceAPISettings
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

func GetServicesAPISettings(configFilePath string) (*ServicesAPISettings, error) {
	var config ServicesAPISettings

	if _, err := toml.DecodeFile(configFilePath, &config); err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return &config, nil
}

func GetEmailSettings() EmailSettings {
	return EmailSettings{
		Email:    fetchEnv("EMAIL", true),
		Password: fetchEnv("EMAIL_PASSWORD", true),
		Host:     fetchEnv("EMAIL_HOST", true),
	}
}

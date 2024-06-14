package infrastructure

import (
	"log"
	"os"
)

func fetchEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("Environment variable %s is not set", name)
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

type CurrencyAPISettings struct {
	CurrencyCode    string
	CurrencyRateURL string
}

type EmailSettings struct {
	Email    string
	Password string
	Host     string
}

func GetDatabaseSettings() DatabaseSettings {
	return DatabaseSettings{
		User:     fetchEnv("POSTGRES_USER"),
		Password: fetchEnv("POSTGRES_PASSWORD"),
		Database: fetchEnv("POSTGRES_DB"),
		Host:     fetchEnv("DB_HOST"),
		Port:     fetchEnv("DB_PORT"),
	}
}

func GetCurrencyAPISettings() CurrencyAPISettings {
	return CurrencyAPISettings{
		CurrencyCode:    fetchEnv("CURRENCY_CODE"),
		CurrencyRateURL: fetchEnv("CURRENCY_RATE_URL"),
	}
}

func GetEmailSettings() EmailSettings {
	return EmailSettings{
		Email:    fetchEnv("EMAIL"),
		Password: fetchEnv("EMAIL_PASSWORD"),
		Host:     fetchEnv("EMAIL_HOST"),
	}
}

package infrastructure

import (
	"log"
	"os"
)

func fetchEnv(name string, strictArg ...bool) string {
	strict := false

	if len(strictArg) > 0 {
		strict = strictArg[0]
	}

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

type CurrencyAPISettings struct {
	CurrencyAPIURL     string
	FawazaAPIURL       string
	ExchangerateAPIURL string
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
		CurrencyAPIURL:     fetchEnv("CURRENCY_API_URL"),
		FawazaAPIURL:       fetchEnv("FAWAZA_API_URL"),
		ExchangerateAPIURL: fetchEnv("EXCHANGERATE_API_URL"),
	}
}

func GetEmailSettings() EmailSettings {
	return EmailSettings{
		Email:    fetchEnv("EMAIL"),
		Password: fetchEnv("EMAIL_PASSWORD"),
		Host:     fetchEnv("EMAIL_HOST"),
	}
}

package infrastructure

import (
	"log"
	"os"
)

func FetchEnv(name string, strict bool) string {
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

func GetDatabaseSettings() DatabaseSettings {
	return DatabaseSettings{
		User:     FetchEnv("POSTGRES_USER", true),
		Password: FetchEnv("POSTGRES_PASSWORD", true),
		Database: FetchEnv("POSTGRES_DB", true),
		Host:     FetchEnv("DB_HOST", true),
		Port:     FetchEnv("DB_PORT", true),
	}
}

func GetCurrencyAPISettings() CurrencyAPISettings {
	return CurrencyAPISettings{
		CurrencyAPIURL:     FetchEnv("CURRENCY_API_URL", false),
		FawazaAPIURL:       FetchEnv("FAWAZA_API_URL", false),
		ExchangerateAPIURL: FetchEnv("EXCHANGERATE_API_URL", false),
	}
}

package infrastructure

import (
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

func GetCurrencyAPISettings() CurrencyAPISettings {
	return CurrencyAPISettings{
		CurrencyAPIURL:     fetchEnv("CURRENCY_API_URL", false),
		FawazaAPIURL:       fetchEnv("FAWAZA_API_URL", false),
		ExchangerateAPIURL: fetchEnv("EXCHANGERATE_API_URL", false),
	}
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

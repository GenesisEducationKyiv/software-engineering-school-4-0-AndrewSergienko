package infrastructure

import (
	"fmt"
	"log/slog"
	"os"
)

func fetchEnv(name string, strict bool) string { // nolint: all
	value := os.Getenv(name)
	if value == "" {
		if strict {
			slog.Error(fmt.Sprintf("Environment variable %s is not set", name))
			os.Exit(1)
		}
		slog.Warn(fmt.Sprintf("Environment variable %s is not set", name))
	}
	slog.Debug(fmt.Sprintf("Environment variable - %s: %s", name, value))
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

type BrokerSettings struct {
	URL string
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

func GetBrokerSettings() BrokerSettings {
	return BrokerSettings{
		URL: fetchEnv("BROKER_URL", true),
	}
}

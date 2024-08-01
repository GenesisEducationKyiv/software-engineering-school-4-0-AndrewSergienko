package infrastructure

import (
	"fmt"
	"log/slog"
	"os"
)

func FetchEnv(name string, strict bool) string { // nolint: all
	value := os.Getenv(name)
	if value == "" {
		if strict {
			slog.Error(fmt.Sprintf("Environment variable %s is not set", name))
			panic(fmt.Sprintf("Environment variable %s is not set", name))
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

type CacheSettings struct {
	URL string
}

func GetCurrencyAPISettings() CurrencyAPISettings {
	return CurrencyAPISettings{
		CurrencyAPIURL:     FetchEnv("CURRENCY_API_URL", false),
		FawazaAPIURL:       FetchEnv("FAWAZA_API_URL", false),
		ExchangerateAPIURL: FetchEnv("EXCHANGERATE_API_URL", false),
	}
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

func GetBrokerSettings() BrokerSettings {
	return BrokerSettings{
		URL: FetchEnv("BROKER_URL", true),
	}
}

func GetCacheSettings() CacheSettings {
	return CacheSettings{
		URL: FetchEnv("CACHE_URL", true),
	}
}

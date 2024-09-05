package infrastructure

import (
	"fmt"
	"github.com/BurntSushi/toml"
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

type BrokerSettings struct {
	URL string
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

func GetServicesAPISettings(configFilePath string) (*ServicesAPISettings, error) {
	var config ServicesAPISettings

	if _, err := toml.DecodeFile(configFilePath, &config); err != nil {
		slog.Warn(fmt.Sprintf("Read TOML services config error: %v", err))
		return nil, err
	}
	return &config, nil
}

func GetEmailSettings() EmailSettings {
	return EmailSettings{
		Email:    FetchEnv("EMAIL", true),
		Password: FetchEnv("EMAIL_PASSWORD", true),
		Host:     FetchEnv("EMAIL_HOST", true),
	}
}

func GetBrokerSettings() BrokerSettings {
	return BrokerSettings{
		URL: FetchEnv("BROKER_URL", true),
	}
}

package infrastructure

import (
	"fmt"
	"github.com/BurntSushi/toml"
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

type EmailSettings struct {
	Email    string
	Password string
	Host     string
}

type CurrencyRateServiceAPISettings struct {
	Host           string
	GetCurrencyURL string
}

type SubscriberServiceAPISettings struct {
	Host              string
	GetSubscribersURL string
}

type ServicesAPISettings struct {
	CurrencyRate *CurrencyRateServiceAPISettings
	Subscriber   *SubscriberServiceAPISettings
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
		Email:    FetchEnv("EMAIL", true),
		Password: FetchEnv("EMAIL_PASSWORD", true),
		Host:     FetchEnv("EMAIL_HOST", true),
	}
}

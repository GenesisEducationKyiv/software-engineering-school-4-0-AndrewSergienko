package infrastructure

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"go_service/internal/infrastructure"
)

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
		Email:    infrastructure.FetchEnv("EMAIL", true),
		Password: infrastructure.FetchEnv("EMAIL_PASSWORD", true),
		Host:     infrastructure.FetchEnv("EMAIL_HOST", true),
	}
}

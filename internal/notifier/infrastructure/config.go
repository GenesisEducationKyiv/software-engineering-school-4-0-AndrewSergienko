package infrastructure

import "go_service/internal/infrastructure"

type EmailSettings struct {
	Email    string
	Password string
	Host     string
}

type CurrencyServiceAPISettings struct {
	Host           string
	GetCurrencyURL string
}

type SubscriberServiceAPISettings struct {
	Host              string
	GetSubscribersURL string
}

func GetCurrencyServiceAPISettings() CurrencyServiceAPISettings {
	return CurrencyServiceAPISettings{
		Host:           "http://localhost:8080",
		GetCurrencyURL: "/rates/",
	}
}

func GetSubscriberServiceAPISettings() SubscriberServiceAPISettings {
	return SubscriberServiceAPISettings{
		Host:              "http://localhost:8080",
		GetSubscribersURL: "/subscribers/",
	}
}

func GetEmailSettings() EmailSettings {
	return EmailSettings{
		Email:    infrastructure.FetchEnv("EMAIL", true),
		Password: infrastructure.FetchEnv("EMAIL_PASSWORD", true),
		Host:     infrastructure.FetchEnv("EMAIL_HOST", true),
	}
}

package infrastructure

import (
	"log"
	"os"
)

func fetchEnv(name string, strict bool) string {
	value := os.Getenv(name)
	if value == "" {
		if strict {
			log.Fatalf("Environment variable %s is not set", name)
		}
		log.Printf("WARN: Environment variable %s is not set\n", name)
	}

	return value
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

package readers

import (
	"go_service/internal/currency_rate/services"
	"log"
	"strings"
)

type CurrencyAPICurrencyReader struct {
	APIURL string
}

func NewCurrencyAPICurrencyReader(url string) *CurrencyAPICurrencyReader {
	if url == "" {
		log.Printf("WARNING: CurrencyAPICurrencyReader: url is empty")
		return nil
	}
	return &CurrencyAPICurrencyReader{
		APIURL: url,
	}
}

func (cr *CurrencyAPICurrencyReader) GetCurrencyRate(from string, to string) (float32, error) {
	data, err := ReadHTTP(cr.APIURL + "&base_currency=" + strings.ToUpper(from))
	if err != nil {
		return 0, err
	}

	if rates, ok := (*data)["data"].(map[string]interface{}); ok {
		if rate, ok := rates[strings.ToUpper(to)].(map[string]interface{}); ok {
			if value, ok := rate["value"].(float64); ok {
				log.Printf("INFO: CurrencyAPICurrencyReade: rate %.2f", value)
				return float32(value), nil
			}
		}
	}

	return 0, &services.CurrencyNotExistsError{Currency: from, Source: "CurrencyAPI"}
}

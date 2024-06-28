package readers

import (
	"go_service/internal/services"
	"log"
	"strings"
)

type ExchangerateAPICurrencyReader struct {
	APIURL string
}

func NewExchangerateAPICurrencyReader(url string) *ExchangerateAPICurrencyReader {
	if url == "" {
		log.Printf("WARNING: ExchangerateAPICurrencyReader: url is empty")
		return nil
	}
	return &ExchangerateAPICurrencyReader{
		APIURL: url,
	}
}

func (cr *ExchangerateAPICurrencyReader) GetCurrencyRate(from string, to string) (float32, error) {
	data, err := ReadHTTP(cr.APIURL + strings.ToUpper(from))
	if err != nil {
		return 0, err
	}

	if rates, ok := (*data)["conversion_rates"].(map[string]interface{}); ok {
		if rate, ok := rates[strings.ToUpper(to)].(float64); ok {
			log.Printf("INFO: ExchangerateAPICurrencyReader: rate %.2f", rate)
			return float32(rate), nil
		}
	}

	return 0, &services.CurrencyNotExistsError{Currency: from, Source: "ExchangerateAPI"}
}

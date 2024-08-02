package readers

import (
	"fmt"
	"go_service/internal/rateservice/currencyrate/services"
	"go_service/internal/rateservice/infrastructure/metrics"
	"log"
	"log/slog"
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
		metrics.RateSourceTotalRequests.WithLabelValues("CurrencyAPI", "error").Inc()
		return 0, err
	}

	if rates, ok := (*data)["data"].(map[string]interface{}); ok {
		if rate, ok := rates[strings.ToUpper(to)].(map[string]interface{}); ok {
			if value, ok := rate["value"].(float64); ok {
				metrics.RateSourceTotalRequests.WithLabelValues("CurrencyAPI", "success").Inc()
				slog.Info(fmt.Sprintf("CurrencyAPICurrencyReade: rate %.2f", value))
				return float32(value), nil
			}
		}
	}

	metrics.RateSourceTotalRequests.WithLabelValues("CurrencyAPI", "error").Inc()
	return 0, &services.CurrencyNotExistsError{Currency: from, Source: "CurrencyAPI"}
}

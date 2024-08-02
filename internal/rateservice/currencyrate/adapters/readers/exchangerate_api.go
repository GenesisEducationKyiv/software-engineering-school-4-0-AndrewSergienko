package readers

import (
	"fmt"
	"go_service/internal/rateservice/currencyrate/services"
	"go_service/internal/rateservice/infrastructure/metrics"
	"log"
	"log/slog"
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
		metrics.RateSourceTotalRequests.WithLabelValues("ExchangerateAPI", "error").Inc()
		return 0, err
	}

	if rates, ok := (*data)["conversion_rates"].(map[string]interface{}); ok {
		if rate, ok := rates[strings.ToUpper(to)].(float64); ok {
			metrics.RateSourceTotalRequests.WithLabelValues("ExchangerateAPI", "success").Inc()
			slog.Info(fmt.Sprintf("ExchangerateAPICurrencyReader: rate %.2f", rate))
			return float32(rate), nil
		}
	}

	metrics.RateSourceTotalRequests.WithLabelValues("ExchangerateAPI", "error").Inc()
	return 0, &services.CurrencyNotExistsError{Currency: from, Source: "ExchangerateAPI"}
}

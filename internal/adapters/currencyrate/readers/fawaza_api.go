package readers

import (
	"go_service/internal/services"
	"log"
	"strings"
)

type FawazaAPICurrencyReader struct {
	APIURL string
}

func NewFawazaAPICurrencyReader(url string) *FawazaAPICurrencyReader {
	if url == "" {
		log.Printf("WARNING: FawazaAPICurrencyReader: url is empty")
		return nil
	}
	return &FawazaAPICurrencyReader{
		APIURL: url,
	}
}

func (cr *FawazaAPICurrencyReader) GetCurrencyRate(from string, to string) (float32, error) {
	from = strings.ToLower(from)
	to = strings.ToLower(to)

	data, err := ReadHTTP(cr.APIURL + strings.ToLower(from) + ".json")
	if err != nil {
		return 0, err
	}

	if rates, ok := (*data)[from].(map[string]interface{}); ok {
		if rate, ok := rates[to].(float64); ok {
			log.Printf("INFO: FawazaAPICurrencyReader: rate %.2f", rate)
			return float32(rate), nil
		}
	}

	return 0, &services.CurrencyNotExistsError{Currency: from, Source: "FawazaAPI"}
}

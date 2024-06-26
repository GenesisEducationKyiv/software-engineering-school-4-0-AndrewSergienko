package adapters

import (
	"encoding/json"
	"fmt"
	"go_service/internal/infrastructure"
	"io"
	"net/http"
)

type Currency struct {
	R030         int     `json:"r030"`
	Txt          string  `json:"txt"`
	Rate         float32 `json:"rate"`
	CC           string  `json:"cc"`
	ExchangeDate string  `json:"exchangedate"`
}

type APICurrencyReader struct {
	APIURL       string
	CurrencyCode string
}

func NewAPICurrencyReader(settings infrastructure.CurrencyAPISettings) *APICurrencyReader {
	return &APICurrencyReader{
		APIURL:       settings.CurrencyRateURL,
		CurrencyCode: settings.CurrencyCode,
	}
}

func (cr *APICurrencyReader) GetUSDCurrencyRate() (float32, error) {
	resp, err := http.Get(cr.APIURL)
	if err != nil {
		return 0, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var currencies []Currency
	err = json.Unmarshal(body, &currencies)
	if err != nil {
		return 0, err
	}

	for _, currency := range currencies {
		if currency.CC == cr.CurrencyCode {
			return currency.Rate, nil
		}
	}

	return 0, fmt.Errorf("currency not found")
}

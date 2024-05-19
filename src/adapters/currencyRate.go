package adapters

import (
	"encoding/json"
	"fmt"
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
	ApiUrl       string
	CurrencyCode string
}

func (cr *APICurrencyReader) GetUSDCurrencyRate() (float32, error) {
	resp, err := http.Get(cr.ApiUrl)
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

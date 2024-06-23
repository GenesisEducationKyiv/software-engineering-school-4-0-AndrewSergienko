package readers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
	url := cr.APIURL + "&base_currency=" + strings.ToUpper(from)

	resp, err := http.Get(url)
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

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, err
	}

	if rates, ok := data["data"].(map[string]interface{}); ok {
		if rate, ok := rates[strings.ToUpper(to)].(map[string]interface{}); ok {
			if value, ok := rate["value"].(float64); ok {
				log.Printf("INFO: CurrencyAPICurrencyReade: rate %.2f", value)
				return float32(value), nil
			}
		}
	}

	return 0, fmt.Errorf("currency not found")
}

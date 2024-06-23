package readers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
	url := cr.APIURL + strings.ToUpper(from)

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

	if rates, ok := data["conversion_rates"].(map[string]interface{}); ok {
		if rate, ok := rates[strings.ToUpper(to)].(float64); ok {
			log.Printf("INFO: ExchangerateAPICurrencyReader: rate %.2f", rate)
			return float32(rate), nil
		}
	}

	return 0, fmt.Errorf("currency not found")
}

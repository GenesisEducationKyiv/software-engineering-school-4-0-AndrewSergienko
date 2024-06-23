package readers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
	resp, err := http.Get(cr.APIURL + from + ".json")
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

	if rates, ok := data[from].(map[string]interface{}); ok {
		if rate, ok := rates[to].(float64); ok {
			log.Printf("INFO: FawazaAPICurrencyReader: rate %.2f", rate)
			return float32(rate), nil
		}
	}

	return 0, fmt.Errorf("currency not found")
}

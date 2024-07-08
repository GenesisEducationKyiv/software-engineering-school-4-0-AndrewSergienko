package adapters

import (
	"encoding/json"
	"fmt"
	"go_service/internal/notifier/infrastructure"
)

type CurrencyRateResponse struct {
	Rate float32 `json:"rate"`
}

type CurrencyRateAdapter struct {
	currencyServiceSettings *infrastructure.CurrencyRateServiceAPISettings
}

func NewCurrencyRateAdapter(
	currencyServiceSettings *infrastructure.CurrencyRateServiceAPISettings,
) CurrencyRateAdapter {
	return CurrencyRateAdapter{currencyServiceSettings}
}

func (adapter CurrencyRateAdapter) GetCurrencyRate(from string, to string) (float32, error) {
	url := fmt.Sprintf(
		"%s%s?from=%s&to=%s",
		adapter.currencyServiceSettings.Host,
		adapter.currencyServiceSettings.GetCurrencyURL,
		from,
		to,
	)
	body, err := ReadHTTP(url)

	if err != nil {
		return 0, err
	}

	var response CurrencyRateResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return 0, err
	}
	return response.Rate, nil
}

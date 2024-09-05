package currencyrate

import (
	"encoding/json"
	"fmt"
	"go_service/internal/notifier/infrastructure"
)

type Response struct {
	Rate float32 `json:"rate"`
}

type Adapter struct {
	currencyServiceSettings *infrastructure.CurrencyRateServiceAPISettings
}

func NewCurrencyRateAdapter(
	currencyServiceSettings *infrastructure.CurrencyRateServiceAPISettings,
) Adapter {
	return Adapter{currencyServiceSettings}
}

func (adapter Adapter) GetCurrencyRate(from string, to string) (float32, error) {
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

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return 0, err
	}
	return response.Rate, nil
}

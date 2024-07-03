package adapters

import (
	readers2 "go_service/internal/currencyrate/adapters/readers"
	"go_service/internal/infrastructure"
)

type APICurrencyReader interface {
	GetCurrencyRate(from string, to string) (float32, error)
}

func CreateReaders(settings infrastructure.CurrencyAPISettings) []APICurrencyReader {
	var apiReaders []APICurrencyReader

	fawazaAPIReader := readers2.NewFawazaAPICurrencyReader(settings.FawazaAPIURL)
	currencyAPIReader := readers2.NewCurrencyAPICurrencyReader(settings.CurrencyAPIURL)
	exchangerateAPIReader := readers2.NewExchangerateAPICurrencyReader(settings.ExchangerateAPIURL)

	// TODO: eliminate the violation of the Open/Close principle
	if fawazaAPIReader != nil {
		apiReaders = append(apiReaders, fawazaAPIReader)
	}
	if currencyAPIReader != nil {
		apiReaders = append(apiReaders, currencyAPIReader)
	}
	if exchangerateAPIReader != nil {
		apiReaders = append(apiReaders, exchangerateAPIReader)
	}
	return apiReaders
}

package adapters

import (
	"go_service/internal/currencyrate/adapters/readers"
	"go_service/internal/infrastructure"
)

type APICurrencyReader interface {
	GetCurrencyRate(from string, to string) (float32, error)
}

func CreateReaders(settings infrastructure.CurrencyAPISettings) []APICurrencyReader {
	var apiReaders []APICurrencyReader

	fawazaAPIReader := readers.NewFawazaAPICurrencyReader(settings.FawazaAPIURL)
	currencyAPIReader := readers.NewCurrencyAPICurrencyReader(settings.CurrencyAPIURL)
	exchangerateAPIReader := readers.NewExchangerateAPICurrencyReader(settings.ExchangerateAPIURL)

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

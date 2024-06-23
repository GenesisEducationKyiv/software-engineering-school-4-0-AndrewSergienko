package currency_rate

import (
	readers2 "go_service/internal/adapters/currency_rate/readers"
	"go_service/internal/infrastructure"
)

type APICurrencyReader interface {
	GetCurrencyRate(from string, to string) (float32, error)
}

func CreateReaders(settings infrastructure.CurrencyAPISettings) []APICurrencyReader {
	var readers []APICurrencyReader

	fawazaAPIReader := readers2.NewFawazaAPICurrencyReader(settings.FawazaAPIURL)
	currencyAPIReader := readers2.NewCurrencyAPICurrencyReader(settings.CurrencyAPIURL)
	exchangerateAPIReader := readers2.NewExchangerateAPICurrencyReader(settings.ExchangerateAPIURL)

	// TODO: eliminate the violation of the Open/Close principle
	if fawazaAPIReader != nil {
		readers = append(readers, fawazaAPIReader)
	}
	if currencyAPIReader != nil {
		readers = append(readers, currencyAPIReader)
	}
	if exchangerateAPIReader != nil {
		readers = append(readers, exchangerateAPIReader)
	}
	return readers
}

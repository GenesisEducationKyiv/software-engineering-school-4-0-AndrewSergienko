package app

import (
	"go_service/internal/currency_rate/adapters"
	"go_service/internal/currency_rate/presentation"
	"go_service/internal/currency_rate/services"
	"go_service/internal/infrastructure"
)

type IoC struct {
	currencyRateFacade *adapters.APIReaderFacade
}

func NewIoC(currencyAPISettings infrastructure.CurrencyAPISettings) *IoC {
	readers := adapters.CreateReaders(currencyAPISettings)
	return &IoC{
		currencyRateFacade: adapters.NewAPIReaderFacade(readers),
	}
}

func (ioc *IoC) GetCurrencyRate() presentation.Interactor[
	services.GetCurrencyRateInputDTO,
	services.GetCurrencyRateOutputDTO,
] {
	return services.NewGetCurrencyRate(ioc.currencyRateFacade)
}

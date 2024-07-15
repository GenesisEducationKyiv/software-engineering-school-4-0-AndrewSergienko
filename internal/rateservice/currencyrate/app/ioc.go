package app

import (
	"go_service/internal/rateservice/currencyrate/adapters"
	"go_service/internal/rateservice/currencyrate/infrastructure"
	"go_service/internal/rateservice/currencyrate/presentation"
	"go_service/internal/rateservice/currencyrate/services"
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

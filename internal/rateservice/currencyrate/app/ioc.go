package app

import (
	"go_service/internal/rateservice/currencyrate/adapters"
	"go_service/internal/rateservice/currencyrate/presentation"
	"go_service/internal/rateservice/currencyrate/services"
	"go_service/internal/rateservice/infrastructure"
)

type IoC struct {
	currencyRateFacade *adapters.APIReaderFacade
}

func NewIoC(cacheAdapter adapters.CacheRateAdapter, currencyAPISettings infrastructure.CurrencyAPISettings) *IoC {
	readers := adapters.CreateReaders(currencyAPISettings)
	readers = append([]adapters.APICurrencyReader{cacheAdapter}, readers...)
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

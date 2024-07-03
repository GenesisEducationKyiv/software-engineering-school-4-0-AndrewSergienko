package presentation

import (
	"go_service/internal/currencyrate/services"
)

type Interactor[InputDTO, OutputDTO any] interface {
	Handle(data InputDTO) OutputDTO
}

type InteractorFactory interface {
	GetCurrencyRate() Interactor[services.GetCurrencyRateInputDTO, services.GetCurrencyRateOutputDTO]
}

package internal

import (
	services2 "go_service/internal/currency_rate/services"
	"go_service/internal/notifier/services"
	services3 "go_service/internal/subscribers/services"
)

type Interactor[InputDTO, OutputDTO any] interface {
	Handle(data InputDTO) OutputDTO
}

type InteractorFactory interface {
	Subscribe() Interactor[services3.SubscribeInputDTO, services3.SubscribeOutputDTO]
	GetCurrencyRate() Interactor[services2.GetCurrencyRateInputDTO, services2.GetCurrencyRateOutputDTO]
	SendNotification() Interactor[services.SendNotificationInputDTO, services.SendNotificationOutputDTO]
}

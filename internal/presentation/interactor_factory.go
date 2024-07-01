package presentation

import "go_service/internal/services"

type Interactor[InputDTO, OutputDTO any] interface {
	Handle(data InputDTO) OutputDTO
}

type InteractorFactory interface {
	Subscribe() Interactor[services.SubscribeInputDTO, services.SubscribeOutputDTO]
	GetCurrencyRate() Interactor[services.GetCurrencyRateInputDTO, services.GetCurrencyRateOutputDTO]
	SendNotification() Interactor[services.SendNotificationInputDTO, services.SendNotificationOutputDTO]
}

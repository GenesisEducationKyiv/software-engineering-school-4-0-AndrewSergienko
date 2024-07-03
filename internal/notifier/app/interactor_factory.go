package app

import (
	"go_service/internal/notifier/services"
)

type Interactor[InputDTO, OutputDTO any] interface {
	Handle(data InputDTO) OutputDTO
}

type InteractorFactory interface {
	SendNotification() Interactor[services.SendNotificationInputDTO, services.SendNotificationOutputDTO]
}

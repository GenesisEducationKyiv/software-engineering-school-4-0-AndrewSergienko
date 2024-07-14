package app

import (
	"go_service/internal/notifier/services/createsubscriber"
	"go_service/internal/notifier/services/deletesubscriber"
	"go_service/internal/notifier/services/sendnotification"
)

type Interactor[InputDTO, OutputDTO any] interface {
	Handle(data InputDTO) OutputDTO
}

type InteractorFactory interface {
	SendNotification() Interactor[sendnotification.InputData, sendnotification.OutputData]
	CreateSubscriber() Interactor[createsubscriber.InputData, createsubscriber.OutputData]
	DeleteSubscriber() Interactor[deletesubscriber.InputData, deletesubscriber.OutputData]
}

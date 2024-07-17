package presentation

import (
	"go_service/internal/rateservice/customers/services/getall"
	"go_service/internal/rateservice/customers/services/subscribe"
	"go_service/internal/rateservice/customers/services/unsubscribe"
)

type Interactor[InputDTO, OutputDTO any] interface {
	Handle(data InputDTO) OutputDTO
}

type InteractorFactory interface {
	Subscribe() Interactor[subscribe.InputDTO, subscribe.OutputDTO]
	Unsubscribe() Interactor[unsubscribe.InputDTO, unsubscribe.OutputDTO]
	GetAll() Interactor[getall.InputDTO, getall.OutputDTO]
}

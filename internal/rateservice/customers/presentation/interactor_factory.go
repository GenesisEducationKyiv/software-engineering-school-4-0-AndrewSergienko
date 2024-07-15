package presentation

import (
	"go_service/internal/rateservice/customers/services/getall"
	"go_service/internal/rateservice/customers/services/subscribe"
)

type Interactor[InputDTO, OutputDTO any] interface {
	Handle(data InputDTO) OutputDTO
}

type InteractorFactory interface {
	Subscribe() Interactor[subscribe.InputDTO, subscribe.OutputDTO]
	GetAll() Interactor[getall.InputDTO, getall.OutputDTO]
}

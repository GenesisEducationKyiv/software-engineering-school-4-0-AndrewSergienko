package presentation

import (
	"go_service/internal/subscribers/services/get_all"
	"go_service/internal/subscribers/services/subscribe"
)

type Interactor[InputDTO, OutputDTO any] interface {
	Handle(data InputDTO) OutputDTO
}

type InteractorFactory interface {
	Subscribe() Interactor[subscribe.InputDTO, subscribe.OutputDTO]
	GetAll() Interactor[get_all.InputDTO, get_all.OutputDTO]
}

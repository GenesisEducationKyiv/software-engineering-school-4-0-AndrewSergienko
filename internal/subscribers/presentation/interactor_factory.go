package presentation

import (
	services3 "go_service/internal/subscribers/services"
)

type Interactor[InputDTO, OutputDTO any] interface {
	Handle(data InputDTO) OutputDTO
}

type InteractorFactory interface {
	Subscribe() Interactor[services3.SubscribeInputDTO, services3.SubscribeOutputDTO]
}

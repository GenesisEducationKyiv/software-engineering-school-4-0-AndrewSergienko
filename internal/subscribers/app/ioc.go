package app

import (
	"go_service/internal/subscribers/adapters"
	"go_service/internal/subscribers/presentation"
	"go_service/internal/subscribers/services/get_all"
	"go_service/internal/subscribers/services/subscribe"
	"gorm.io/gorm"
)

type IoC struct {
	subscriberAdapter *adapters.SubscriberAdapter
}

func NewIoC(db *gorm.DB) *IoC {
	return &IoC{
		subscriberAdapter: adapters.NewSubscriberAdapter(db),
	}
}

func (ioc *IoC) Subscribe() presentation.Interactor[
	subscribe.InputDTO,
	subscribe.OutputDTO,
] {
	return subscribe.NewSubscribe(ioc.subscriberAdapter)
}

func (ioc *IoC) GetAll() presentation.Interactor[
	get_all.InputDTO,
	get_all.OutputDTO,
] {
	return get_all.NewGetAllHandler(ioc.subscriberAdapter)
}

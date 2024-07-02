package app

import (
	"go_service/internal/subscribers/adapters"
	"go_service/internal/subscribers/presentation"
	"go_service/internal/subscribers/services"
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
	services.SubscribeInputDTO,
	services.SubscribeOutputDTO,
] {
	return services.NewSubscribe(ioc.subscriberAdapter)
}

package app

import (
	"go_service/internal/customers/adapters"
	"go_service/internal/customers/presentation"
	"go_service/internal/customers/services/getall"
	"go_service/internal/customers/services/subscribe"
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
	return subscribe.New(ioc.subscriberAdapter)
}

func (ioc *IoC) GetAll() presentation.Interactor[
	getall.InputDTO,
	getall.OutputDTO,
] {
	return getall.New(ioc.subscriberAdapter)
}

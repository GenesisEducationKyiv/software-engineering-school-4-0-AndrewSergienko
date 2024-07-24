package app

import (
	"github.com/nats-io/nats.go"
	"go_service/internal/rateservice/customers/adapters"
	"go_service/internal/rateservice/customers/presentation"
	"go_service/internal/rateservice/customers/services/createcustomer"
	"go_service/internal/rateservice/customers/services/deletecustomer"
	"go_service/internal/rateservice/customers/services/getall"
	"gorm.io/gorm"
)

type IoC struct {
	subscriberAdapter *adapters.SubscriberAdapter
	natsEventAdapter  adapters.NatsEventEmitter
}

func NewIoC(db *gorm.DB, conn nats.JetStreamContext) *IoC {
	return &IoC{
		subscriberAdapter: adapters.NewSubscriberAdapter(db),
		natsEventAdapter:  adapters.NewNatsEventEmitter(conn),
	}
}

func (ioc *IoC) CreateCustomer() presentation.Interactor[
	createcustomer.InputData,
	createcustomer.OutputData,
] {
	return createcustomer.New(ioc.subscriberAdapter, ioc.natsEventAdapter)
}

func (ioc *IoC) DeleteCustomer() presentation.Interactor[
	deletecustomer.InputData,
	deletecustomer.OutputData,
] {
	return deletecustomer.New(ioc.subscriberAdapter, ioc.natsEventAdapter)
}

func (ioc *IoC) GetAll() presentation.Interactor[
	getall.InputDTO,
	getall.OutputDTO,
] {
	return getall.New(ioc.subscriberAdapter)
}

package app

import (
	"context"
	"github.com/nats-io/nats.go/jetstream"
	"go_service/internal/rateservice/customers/adapters"
	"go_service/internal/rateservice/customers/presentation"
	"go_service/internal/rateservice/customers/services/createcustomer"
	"go_service/internal/rateservice/customers/services/deletecustomer"
	"gorm.io/gorm"
)

type IoC struct {
	subscriberAdapter *adapters.SubscriberAdapter
	natsEventAdapter  adapters.NatsEventEmitter
}

func NewIoC(ctx context.Context, db *gorm.DB, conn jetstream.JetStream) *IoC {
	return &IoC{
		subscriberAdapter: adapters.NewSubscriberAdapter(db),
		natsEventAdapter:  adapters.NewNatsEventEmitter(ctx, conn),
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

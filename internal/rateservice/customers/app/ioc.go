package app

import (
	"github.com/nats-io/nats.go"
	"go_service/internal/rateservice/customers/adapters"
	"go_service/internal/rateservice/customers/presentation"
	"go_service/internal/rateservice/customers/services/getall"
	"go_service/internal/rateservice/customers/services/subscribe"
	"gorm.io/gorm"
)

type IoC struct {
	subscriberAdapter *adapters.SubscriberAdapter
	natsEventAdapter  adapters.NatsEventEmitter
}

func NewIoC(db *gorm.DB, nc *nats.Conn) *IoC {
	return &IoC{
		subscriberAdapter: adapters.NewSubscriberAdapter(db),
		natsEventAdapter:  adapters.NewNatsEventEmitter(nc),
	}
}

func (ioc *IoC) Subscribe() presentation.Interactor[
	subscribe.InputDTO,
	subscribe.OutputDTO,
] {
	return subscribe.New(ioc.subscriberAdapter, ioc.natsEventAdapter)
}

func (ioc *IoC) GetAll() presentation.Interactor[
	getall.InputDTO,
	getall.OutputDTO,
] {
	return getall.New(ioc.subscriberAdapter)
}

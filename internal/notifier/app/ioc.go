package app

import (
	"context"
	"github.com/nats-io/nats.go/jetstream"
	"go_service/internal/notifier/adapters/currencyrate"
	"go_service/internal/notifier/adapters/email"
	"go_service/internal/notifier/adapters/event"
	"go_service/internal/notifier/adapters/subscriber"
	"go_service/internal/notifier/infrastructure"
	"go_service/internal/notifier/services/createsubscriber"
	"go_service/internal/notifier/services/deletesubscriber"
	"go_service/internal/notifier/services/sendnotification"
	"gorm.io/gorm"
)

type IoC struct {
	emailAdapter        email.Adapter
	currencyRateAdapter currencyrate.Adapter
	subscriberAdapter   subscriber.Adapter
	eventEmitter        event.NatsEventEmitter
}

func NewIoC(
	ctx context.Context,
	db *gorm.DB,
	currencyServiceSettings *infrastructure.CurrencyRateServiceAPISettings,
	emailSettings infrastructure.EmailSettings,
	conn jetstream.JetStream,
) *IoC {
	return &IoC{
		currencyRateAdapter: currencyrate.NewCurrencyRateAdapter(currencyServiceSettings),
		subscriberAdapter:   subscriber.NewSubscriberAdapter(db),
		emailAdapter:        email.NewEmailAdapter(emailSettings),
		eventEmitter:        event.NewNatsEventEmitter(ctx, conn),
	}
}

func (ioc *IoC) SendNotification() Interactor[sendnotification.InputData, sendnotification.OutputData] {
	return sendnotification.New(ioc.emailAdapter, &ioc.subscriberAdapter, ioc.currencyRateAdapter)
}

func (ioc *IoC) CreateSubscriber() Interactor[createsubscriber.InputData, createsubscriber.OutputData] {
	return createsubscriber.New(&ioc.subscriberAdapter, &ioc.eventEmitter)
}

func (ioc *IoC) DeleteSubscriber() Interactor[deletesubscriber.InputData, deletesubscriber.OutputData] {
	return deletesubscriber.New(&ioc.subscriberAdapter, &ioc.eventEmitter)
}

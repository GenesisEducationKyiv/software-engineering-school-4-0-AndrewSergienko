package app

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go/jetstream"
	"go_service/internal/rateservice/customers/presentation"
	"go_service/internal/rateservice/customers/services/createcustomer"
	"go_service/internal/rateservice/customers/services/deletecustomer"
	"log/slog"
	"time"
)

type Message struct {
	Title         string                 `json:"title"`
	Type          string                 `json:"type"`
	TransactionID *string                `json:"transaction_id"`
	From          string                 `json:"from"`
	Data          map[string]interface{} `json:"data"`
}

type Consumer struct {
	ctx       context.Context
	js        jetstream.JetStream
	container presentation.InteractorFactory
}

func NewConsumer(ctx context.Context, js jetstream.JetStream, container *IoC) Consumer {
	return Consumer{ctx: ctx, js: js, container: container}
}

func (c Consumer) Run() (jetstream.ConsumeContext, error) {
	ctx, cancel := context.WithTimeout(c.ctx, 5*time.Second)
	defer cancel()

	cons, err := c.js.CreateOrUpdateConsumer(ctx, "events", jetstream.ConsumerConfig{
		Durable:       "customers_consumer",
		AckPolicy:     jetstream.AckExplicitPolicy,
		FilterSubject: "events.*",
	})

	consContext, err := cons.Consume(newMessageHandler(c.container))
	if err != nil {
		return consContext, err
	}
	slog.Info("Consumer started")
	return consContext, err
}

func newMessageHandler(container presentation.InteractorFactory) func(msg jetstream.Msg) {
	return func(msg jetstream.Msg) {
		var event Message
		err := json.Unmarshal(msg.Data(), &event)
		if err != nil {
			slog.Warn(fmt.Sprintf("Error unmarshalling message: %v", err), slog.String("subject", msg.Subject()))
			_ = msg.Nak()
			return
		}

		switch event.Title {
		case "SubscriberCreatedError", "SubscriberCreatedTimeout":
			slog.Info(fmt.Sprintf("Rollback transaction %s", *event.TransactionID))
			interactor := container.DeleteCustomer()
			inputData := deletecustomer.InputData{
				Email:         event.Data["email"].(string),
				TransactionID: event.TransactionID,
				IsRollback:    true,
			}
			interactor.Handle(inputData)
			_ = msg.Ack()
		case "SubscriberDeletedError", "SubscriberDeletedTimeout":
			slog.Info(fmt.Sprintf("Rollback transaction %s", *event.TransactionID))
			interactor := container.CreateCustomer()
			inputData := createcustomer.InputData{
				Email:         event.Data["email"].(string),
				TransactionID: event.TransactionID,
				IsRollback:    true,
			}
			interactor.Handle(inputData)
			_ = msg.Ack()
		default:
			_ = msg.Ack()
		}
	}
}

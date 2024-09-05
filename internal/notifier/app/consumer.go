package app

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go/jetstream"
	"go_service/internal/notifier/infrastructure/broker"
	"go_service/internal/notifier/services/createsubscriber"
	"go_service/internal/notifier/services/deletesubscriber"
	"log"
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
	conn      jetstream.JetStream
	container InteractorFactory
}

func NewConsumer(ctx context.Context, conn jetstream.JetStream, container InteractorFactory) Consumer {
	return Consumer{ctx: ctx, conn: conn, container: container}
}

func (c Consumer) Run() (jetstream.ConsumeContext, error) {
	ctx, cancel := context.WithTimeout(c.ctx, 5*time.Second)
	defer cancel()

	_, _ = broker.NewStream(ctx, c.conn, "events")

	cons, _ := c.conn.CreateOrUpdateConsumer(ctx, "events", jetstream.ConsumerConfig{
		Durable:       "notifier_consumer",
		AckPolicy:     jetstream.AckExplicitPolicy,
		FilterSubject: "events.*",
	})

	consContext, err := cons.Consume(newMessageHandler(c.container))
	if err != nil {
		slog.Error(fmt.Sprintf("Error starting consumer: %v", err))
		return consContext, err
	}
	slog.Info("Consumer started")
	return consContext, err
}

func newMessageHandler(container InteractorFactory) func(msg jetstream.Msg) {
	return func(msg jetstream.Msg) {
		log.Println("Consumed event")
		var event Message
		err := json.Unmarshal(msg.Data(), &event)
		if err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			_ = msg.Nak()
			return
		}

		switch event.Title {
		case "UserCreated":
			interactor := container.CreateSubscriber()
			inputData := createsubscriber.InputData{
				Email:         event.Data["email"].(string),
				TransactionID: event.TransactionID,
			}
			interactor.Handle(inputData)
			_ = msg.Ack()
		case "UserDeleted":
			interactor := container.DeleteSubscriber()
			inputData := deletesubscriber.InputData{
				Email:         event.Data["email"].(string),
				TransactionID: event.TransactionID,
			}
			interactor.Handle(inputData)
			_ = msg.Ack()
		default:
			_ = msg.Ack()
		}
	}
}

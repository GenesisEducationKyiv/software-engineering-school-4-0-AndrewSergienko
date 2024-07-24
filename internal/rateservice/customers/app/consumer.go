package app

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"go_service/internal/notifier/infrastructure/broker"
	"go_service/internal/rateservice/customers/presentation"
	"go_service/internal/rateservice/customers/services/createcustomer"
	"go_service/internal/rateservice/customers/services/deletecustomer"
	"log"
)

type Message struct {
	Title         string                 `json:"title"`
	Type          string                 `json:"type"`
	TransactionID *string                `json:"transaction_id"`
	From          string                 `json:"from"`
	Data          map[string]interface{} `json:"data"`
}

type Consumer struct {
	js        nats.JetStreamContext
	container presentation.InteractorFactory
}

func NewConsumer(js nats.JetStreamContext, container IoC) Consumer {
	return Consumer{js: js, container: &container}
}

func (c Consumer) Run() {
	stream, _ := c.js.StreamInfo("events")
	if stream == nil {
		err := broker.NewStream(c.js, "events")
		if err != nil {
			return
		}
	}

	_, err := c.js.Subscribe("events", newMessageHandler(c.container))
	if err != nil {
		return
	}
	log.Printf("Consumer started")
}

func newMessageHandler(container presentation.InteractorFactory) func(msg *nats.Msg) {
	return func(msg *nats.Msg) {
		var event Message
		err := json.Unmarshal(msg.Data, &event)
		if err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			return
		}

		switch event.Title {
		case "SubscriberCreatedError":
			interactor := container.DeleteCustomer()
			inputData := deletecustomer.InputData{
				Email:         event.Data["email"].(string),
				TransactionID: event.TransactionID,
			}
			interactor.Handle(inputData)
		case "SubscriberDeletedError":
			interactor := container.CreateCustomer()
			inputData := createcustomer.InputData{
				Email:         event.Data["email"].(string),
				TransactionID: event.TransactionID,
			}
			interactor.Handle(inputData)
		}
	}
}

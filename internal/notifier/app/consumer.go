package app

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"go_service/internal/notifier/services/createsubscriber"
	"go_service/internal/notifier/services/deletesubscriber"
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
	conn      nats.JetStreamContext
	container InteractorFactory
}

func NewConsumer(conn nats.JetStreamContext, container InteractorFactory) Consumer {
	return Consumer{conn: conn, container: container}
}

func (c Consumer) Run() {
	_, err := c.conn.Subscribe("events.*", newMessageHandler(c.container))
	if err != nil {
		return
	}
	log.Printf("Consumer started")
}

func newMessageHandler(container InteractorFactory) func(msg *nats.Msg) {
	return func(msg *nats.Msg) {
		log.Println("Consumed event")
		var event Message
		err := json.Unmarshal(msg.Data, &event)
		if err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			return
		}

		switch event.Title {
		case "UserCreated":
			interactor := container.CreateSubscriber()
			inputData := createsubscriber.InputData{
				Email:         event.Data["Email"].(string),
				TransactionID: event.TransactionID,
			}
			interactor.Handle(inputData)
		case "UserDeleted":
			interactor := container.DeleteSubscriber()
			inputData := deletesubscriber.InputData{
				Email:         event.Data["Email"].(string),
				TransactionID: event.TransactionID,
			}
			interactor.Handle(inputData)
		}
	}
}

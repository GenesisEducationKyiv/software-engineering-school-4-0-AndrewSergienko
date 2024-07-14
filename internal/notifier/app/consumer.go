package app

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"go_service/internal/notifier/services/createsubscriber"
	"go_service/internal/notifier/services/deletesubscriber"
	"log"
)

type Message struct {
	Title string                 `json:"title"`
	Type  string                 `json:"type"`
	Data  map[string]interface{} `json:"data"`
}

type Consumer struct {
	nc        *nats.Conn
	container InteractorFactory
}

func NewConsumer(nc *nats.Conn, container InteractorFactory) Consumer {
	return Consumer{nc: nc, container: container}
}

func (c Consumer) Run() {
	_, err := c.nc.Subscribe("events", func(msg *nats.Msg) {
		var event Message
		err := json.Unmarshal(msg.Data, &event)
		if err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			return
		}

		switch event.Title {
		case "UserCreated":
			userCreatedHandle(event, c.container)
		case "UserDeleted":
			userDeletedHandle(event, c.container)
		}
	})
	if err != nil {
		return
	}
	log.Printf("Consumer started")
}

func userCreatedHandle(message Message, container InteractorFactory) *Message {
	interactor := container.CreateSubscriber()
	interactor.Handle(createsubscriber.InputData{Email: message.Data["email"].(string)})
	return nil
}

func userDeletedHandle(message Message, container InteractorFactory) *Message {
	interactor := container.DeleteSubscriber()
	interactor.Handle(deletesubscriber.InputData{Email: message.Data["email"].(string)})
	return nil
}

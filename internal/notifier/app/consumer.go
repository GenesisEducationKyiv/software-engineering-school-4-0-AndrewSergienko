package app

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"go_service/internal/notifier/infrastructure/broker"
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
	js        nats.JetStreamContext
	container InteractorFactory
}

func NewConsumer(js nats.JetStreamContext, container InteractorFactory) Consumer {
	return Consumer{js: js, container: container}
}

func (c Consumer) Run() {
	stream, err := c.js.StreamInfo("events")
	if stream == nil {
		err = broker.NewStream(c.js, "events")
		if err != nil {
			return
		}
	}

	_, err = c.js.Subscribe("events", newMessageHandler(c.container))
	if err != nil {
		return
	}
	log.Printf("Consumer started")
}

func newMessageHandler(container InteractorFactory) func(msg *nats.Msg) {
	return func(msg *nats.Msg) {
		var event Message
		err := json.Unmarshal(msg.Data, &event)
		if err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			return
		}

		switch event.Title {
		case "UserCreated":
			userCreatedHandle(event, container)
		case "UserDeleted":
			userDeletedHandle(event, container)
		}
	}
}

func userCreatedHandle(message Message, container InteractorFactory) {
	interactor := container.CreateSubscriber()
	interactor.Handle(createsubscriber.InputData{Email: message.Data["email"].(string)})
}

func userDeletedHandle(message Message, container InteractorFactory) {
	interactor := container.DeleteSubscriber()
	interactor.Handle(deletesubscriber.InputData{Email: message.Data["email"].(string)})
}

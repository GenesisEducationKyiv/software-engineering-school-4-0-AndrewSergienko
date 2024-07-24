package adapters

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
)

type Message struct {
	Title         string                 `json:"title"`
	Type          string                 `json:"type"`
	TransactionID *string                `json:"transaction_id"`
	From          string                 `json:"from"`
	Data          map[string]interface{} `json:"data"`
}

type NatsEventEmitter struct {
	conn nats.JetStreamContext
}

func NewNatsEventEmitter(conn nats.JetStreamContext) NatsEventEmitter {
	return NatsEventEmitter{conn: conn}
}

func (e NatsEventEmitter) Emit(name string, data map[string]interface{}, transactionID *string) error {
	message := Message{Title: name, Type: "event", From: "notifier", Data: data, TransactionID: transactionID}
	serializedData, err := json.Marshal(message)
	if err != nil {
		return err
	}

	subject := "events"
	if transactionID != nil {
		subject = "events." + *transactionID
	}
	log.Printf("Emitting event: %s to %s\n", name, subject)

	_, err = e.conn.Publish(subject, serializedData)
	return err
}

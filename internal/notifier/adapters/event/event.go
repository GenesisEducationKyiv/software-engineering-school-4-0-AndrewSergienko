package event

import (
	"context"
	"encoding/json"
	"github.com/nats-io/nats.go/jetstream"
	"log"
	"time"
)

type Message struct {
	Title         string                 `json:"title"`
	Type          string                 `json:"type"`
	TransactionID *string                `json:"transaction_id"`
	From          string                 `json:"from"`
	Data          map[string]interface{} `json:"data"`
}

type NatsEventEmitter struct {
	ctx  context.Context
	conn jetstream.JetStream
}

func NewNatsEventEmitter(ctx context.Context, conn jetstream.JetStream) NatsEventEmitter {
	return NatsEventEmitter{ctx: ctx, conn: conn}
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

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	log.Printf("Emitting event: %s to %s\n", name, subject)

	_, err = e.conn.Publish(ctx, subject, serializedData)
	return err
}

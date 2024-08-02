package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go/jetstream"
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

type NatsEventEmitter struct {
	ctx  context.Context
	conn jetstream.JetStream
}

func NewNatsEventEmitter(ctx context.Context, conn jetstream.JetStream) NatsEventEmitter {
	return NatsEventEmitter{ctx: ctx, conn: conn}
}

func (e NatsEventEmitter) Emit(name string, data map[string]interface{}, transactionID *string) error {
	message := Message{Title: name, Type: "event", From: "customers", Data: data, TransactionID: transactionID}
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

	slog.Info(fmt.Sprintf("Emitting event: %s to %s\n", name, subject))
	_, err = e.conn.Publish(ctx, subject, serializedData)
	return err
}

func (e NatsEventEmitter) GetMessages(transactionID string, batchSize int) []Message {
	subject := "events." + transactionID

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	cons, _ := e.conn.CreateOrUpdateConsumer(ctx, "events", jetstream.ConsumerConfig{
		AckPolicy:     jetstream.AckExplicitPolicy,
		FilterSubject: subject,
	})

	fetchOpt := jetstream.FetchMaxWait(2 * time.Second)
	msgBatch, err := cons.Fetch(batchSize, fetchOpt)
	if err != nil {
		log.Println(err)
	}

	var messages []Message
	for msg := range msgBatch.Messages() {
		var message Message
		err = json.Unmarshal(msg.Data(), &message)
		if err == nil {
			messages = append(messages, message)
		}
	}

	return messages
}

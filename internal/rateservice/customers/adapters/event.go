package adapters

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
)

type Message struct {
	Title string                 `json:"title"`
	Type  string                 `json:"type"`
	Data  map[string]interface{} `json:"data"`
}

type NatsEventEmitter struct {
	nc *nats.Conn
}

func NewNatsEventEmitter(nc *nats.Conn) NatsEventEmitter {
	return NatsEventEmitter{nc: nc}
}

func (e NatsEventEmitter) Emit(name string, data map[string]interface{}) error {
	message := Message{Title: name, Type: "event", Data: data}
	serializedData, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return e.nc.Publish("events", serializedData)
}

package adapters

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
)

type NatsEventEmitter struct {
	nc *nats.Conn
}

func NewNatsEventEmitter(nc *nats.Conn) NatsEventEmitter {
	return NatsEventEmitter{nc: nc}
}

func (e NatsEventEmitter) Emit(name string, data map[string]interface{}) error {
	serializedData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return e.nc.Publish(name, serializedData)
}

package broker

import (
	"errors"
	"github.com/nats-io/nats.go"
)

func New() (*nats.Conn, nats.JetStreamContext) {
	conn, _ := nats.Connect("nats://localhost:4222")
	js, _ := conn.JetStream()
	return conn, js
}

func Finalize(conn *nats.Conn) {
	conn.Close()
}

func NewStream(js nats.JetStreamContext, name string) error {
	stream, err := js.StreamInfo(name)
	if err != nil {
		var jsErr *nats.APIError
		if errors.As(err, &jsErr) && jsErr.Code != 404 {
			return err
		}
	}
	if stream == nil {
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     name,
			Subjects: []string{name + ".*"},
		})
		if err != nil {
			return err
		}
	}
	return err
}

package broker

import "github.com/nats-io/nats.go"

func New() (*nats.Conn, nats.JetStreamContext) {
	conn, _ := nats.Connect(nats.DefaultURL)
	js, _ := conn.JetStream()
	return conn, js
}

func Finalize(conn *nats.Conn) {
	conn.Close()
}

func NewStream(js nats.JetStreamContext, name string) error {
	_, err := js.AddStream(&nats.StreamConfig{
		Name:     name,
		Subjects: []string{name + ".*"},
	})
	return err
}

package broker

import (
	"context"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func New() (*nats.Conn, jetstream.JetStream) {
	conn, _ := nats.Connect("nats://localhost:4222")
	js, _ := jetstream.New(conn)
	return conn, js
}

func Finalize(conn *nats.Conn) {
	conn.Close()
}

func NewStream(ctx context.Context, js jetstream.JetStream, name string) (jetstream.Stream, error) {
	return js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     name,
		Subjects: []string{name + ".*"},
	})
}

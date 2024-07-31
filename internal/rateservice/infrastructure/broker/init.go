package broker

import (
	"context"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"log/slog"
)

func New() (*nats.Conn, jetstream.JetStream, error) {
	conn, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		return nil, nil, err
	}
	slog.Debug("Connected to NATS")

	js, err := jetstream.New(conn)
	if err != nil {
		return nil, nil, err
	}
	slog.Debug("Connected to JetStream")

	return conn, js, nil
}

func Finalize(conn *nats.Conn) {
	conn.Close()
	slog.Debug("Connection to NATS closed")
}

func NewStream(ctx context.Context, js jetstream.JetStream, name string) (jetstream.Stream, error) {
	stream, err := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     name,
		Subjects: []string{name, name + ".*"},
	})
	return stream, err
}

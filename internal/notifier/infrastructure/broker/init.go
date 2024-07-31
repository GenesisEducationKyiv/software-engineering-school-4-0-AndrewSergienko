package broker

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"log/slog"
)

func New() (*nats.Conn, jetstream.JetStream) {
	conn, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		slog.Error(fmt.Sprintf("NATS is not available. Error: %s", err))
		return nil, nil
	}
	slog.Debug("Connected to NATS")

	js, err := jetstream.New(conn)
	if err != nil {
		slog.Error(fmt.Sprintf("JetStream is not available. Error: %s", err))
		return nil, nil
	}
	slog.Debug("Connected to JetStream")

	return conn, js
}

func Finalize(conn *nats.Conn) {
	conn.Close()
	slog.Debug("Connection to NATS closed")
}

func NewStream(ctx context.Context, js jetstream.JetStream, name string) (jetstream.Stream, error) {
	return js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     name,
		Subjects: []string{name + ".*"},
	})
}

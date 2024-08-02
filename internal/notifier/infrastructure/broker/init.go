package broker

import (
	"context"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go_service/internal/notifier/infrastructure"
	"log/slog"
)

func New(settings infrastructure.BrokerSettings) (*nats.Conn, jetstream.JetStream, error) {
	conn, err := nats.Connect(settings.URL)
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
	return js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:     name,
		Subjects: []string{name + ".*"},
	})
}

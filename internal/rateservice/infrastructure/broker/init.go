package broker

import "github.com/nats-io/nats.go"

func New() *nats.Conn {
	conn, _ := nats.Connect(nats.DefaultURL)
	return conn
}

func Finalize(conn *nats.Conn) {
	conn.Close()
}

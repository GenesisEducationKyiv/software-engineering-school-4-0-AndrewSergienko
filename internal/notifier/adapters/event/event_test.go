package event

import (
	"context"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/stretchr/testify/suite"
	"go_service/internal/notifier/infrastructure"
	"go_service/internal/notifier/infrastructure/broker"
	"log"
	"time"
)

type NatsEventEmitterTestSuite struct {
	suite.Suite
	conn    *nats.Conn
	js      jetstream.JetStream
	emitter NatsEventEmitter
}

func (suite *NatsEventEmitterTestSuite) SetupTest() {
	ctx := context.Background()
	brokerSettings := infrastructure.GetBrokerSettings()
	conn, js, err := broker.New(brokerSettings)

	suite.NoError(err)

	if conn == nil {
		suite.T().Skip("NATS connection failed")
	}

	_, err = broker.NewStream(ctx, js, "events")

	suite.NoError(err)

	suite.conn = conn
	suite.js = js
	suite.emitter = NewNatsEventEmitter(context.Background(), js)
}

func (suite *NatsEventEmitterTestSuite) TearDownTest() {
	broker.Finalize(suite.conn)
}

func GetMessages(conn jetstream.JetStream, transactionID *string, batchSize int) []Message {
	var subject string

	if transactionID != nil {
		subject = "events." + *transactionID
	} else {
		subject = "events"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	cons, _ := conn.CreateOrUpdateConsumer(ctx, "events", jetstream.ConsumerConfig{
		Durable:       "notifier_consumer_test",
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
		_ = msg.Ack()
	}

	return messages
}

func (suite *NatsEventEmitterTestSuite) TestEmit_Success() {
	data := map[string]interface{}{"key": "value"}

	err := suite.emitter.Emit("TestEvent", data, nil)
	suite.NoError(err)

	message := GetMessages(suite.js, nil, 1)

	suite.True(len(message) == 1)
	suite.Equal("TestEvent", message[0].Title)
	suite.Equal(data, message[0].Data)
}

func (suite *NatsEventEmitterTestSuite) TestEmit_WithTransactionID() {
	transactionID := "12345"
	data := map[string]interface{}{"key": "value"}

	err := suite.emitter.Emit("TestEvent", data, &transactionID)
	suite.NoError(err)

	message := GetMessages(suite.js, &transactionID, 1)

	suite.True(len(message) == 1)
	suite.Equal("TestEvent", message[0].Title)
	suite.Equal(transactionID, *message[0].TransactionID)
	suite.Equal(data, message[0].Data)
}

func (suite *NatsEventEmitterTestSuite) TestEmit_JSONMarshalError() {
	err := suite.emitter.Emit("TestEvent", map[string]interface{}{"key": make(chan int)}, nil)
	suite.Error(err)
}

// func TestNatsEventEmitterSuite(t *testing.T) {
//	suite.Run(t, new(NatsEventEmitterTestSuite))
// }

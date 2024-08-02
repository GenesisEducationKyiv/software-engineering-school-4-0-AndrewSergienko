package app

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/stretchr/testify/suite"
	"go_service/internal/rateservice/customers"
	"go_service/internal/rateservice/customers/adapters"
	"go_service/internal/rateservice/infrastructure"
	"go_service/internal/rateservice/infrastructure/broker"
	"go_service/internal/rateservice/infrastructure/cache"
	"go_service/internal/rateservice/infrastructure/database"
	"gorm.io/gorm"
	"net/http/httptest"
	"testing"
)

type TestMessage struct {
	Title         string                 `json:"title"`
	Type          string                 `json:"type"`
	TransactionID *string                `json:"transaction_id"`
	From          string                 `json:"from"`
	Data          map[string]interface{} `json:"data"`
}

type EventGateway interface {
	Emit(name string, data map[string]interface{}, transactionID *string) error
}

type SubscribersPresentationSuite struct {
	suite.Suite
	db                *gorm.DB
	transaction       *gorm.DB
	webApp            *fiber.App
	conn              *nats.Conn
	js                jetstream.JetStream
	subscriberGateway *adapters.SubscriberAdapter
	eventGateway      EventGateway
	notifierConsumer  jetstream.ConsumeContext
	customersConsumer jetstream.ConsumeContext
}

func (suite *SubscribersPresentationSuite) SetupSuite() {
	ctx := context.Background()

	databaseSettings := infrastructure.GetDatabaseSettings()
	brokerSettings := infrastructure.GetBrokerSettings()

	db, err := database.New(databaseSettings)
	suite.NoError(err)
	suite.db = db

	conn, js, err := broker.New(brokerSettings)
	suite.NoError(err)

	_, err = broker.NewStream(ctx, js, "events")
	suite.NoError(err)

	suite.eventGateway = adapters.NewNatsEventEmitter(ctx, js)

	suite.conn, suite.js = conn, js
}

func (suite *SubscribersPresentationSuite) TearDownSuite() {
	suite.notifierConsumer.Stop()
	broker.Finalize(suite.conn)
}

func runConsumer(js jetstream.JetStream, eventGateway EventGateway, isError bool) jetstream.ConsumeContext {
	ctx := context.Background()
	_, _ = broker.NewStream(ctx, js, "events")

	cons, _ := js.CreateOrUpdateConsumer(ctx, "events", jetstream.ConsumerConfig{
		Durable:       "notifier_consumer_test",
		FilterSubject: "events.*",
	})

	consContext, err := cons.Consume(messageHandler(eventGateway, isError))
	if err != nil {
		return consContext
	}
	return consContext
}

func messageHandler(eventGateway EventGateway, isError bool) func(msg jetstream.Msg) {
	return func(msg jetstream.Msg) {
		var event TestMessage
		err := json.Unmarshal(msg.Data(), &event)
		if err != nil {
			return
		}

		var eventTitle string
		switch event.Title {
		case "UserCreated":
			if !isError {
				eventTitle = "SubscriberCreated"
			} else {
				eventTitle = "SubscriberCreatedError"
			}
			_ = eventGateway.Emit(eventTitle, event.Data, event.TransactionID)
			_ = msg.Ack()
		case "UserDeleted":
			if !isError {
				eventTitle = "SubscriberDeleted"
			} else {
				eventTitle = "SubscriberDeletedError"
			}
			_ = eventGateway.Emit(eventTitle, event.Data, event.TransactionID)
			_ = msg.Ack()
		}
	}
}

func (suite *SubscribersPresentationSuite) SetupTest() {
	ctx := context.Background()
	currencyAPISettings := infrastructure.GetCurrencyAPISettings()
	cacheSettings := infrastructure.GetCacheSettings()

	suite.transaction = suite.db.Begin()

	suite.subscriberGateway = adapters.NewSubscriberAdapter(suite.transaction)
	customersConsumer, err := customers.NewConsumer(ctx, suite.transaction, suite.js).Run()

	cacheClient := cache.New(cacheSettings)

	suite.NoError(err)
	suite.customersConsumer = customersConsumer
	suite.webApp = InitWebApp(ctx, suite.transaction, suite.js, cacheClient, currencyAPISettings)
}

func (suite *SubscribersPresentationSuite) TearDownTest() {
	suite.customersConsumer.Stop()
	suite.transaction.Rollback()
	suite.notifierConsumer.Stop()
}

func (suite *SubscribersPresentationSuite) TestAddSubscriber_Success() {
	suite.notifierConsumer = runConsumer(suite.js, suite.eventGateway, false)

	var jsonStr = []byte(`{"email":"test@gmail.com"}`)

	req := httptest.NewRequest("POST", "/customers/", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp, err := suite.webApp.Test(req)
	suite.Require().NoError(err, "Error executing request")

	suite.Require().Equal("200 OK", resp.Status)

	suite.NotNil(suite.subscriberGateway.GetByEmail("test@gmail.com"))
}

func (suite *SubscribersPresentationSuite) TestAddSubscriber_Error() {
	suite.notifierConsumer = runConsumer(suite.js, suite.eventGateway, true)

	var jsonStr = []byte(`{"email":"test@gmail.com"}`)

	req := httptest.NewRequest("POST", "/customers/", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp, err := suite.webApp.Test(req)
	suite.Require().NoError(err, "Error executing request")

	suite.Require().Equal("500 Internal Server Error", resp.Status)

	suite.Nil(suite.subscriberGateway.GetByEmail("test@gmail.com"))
}

func (suite *SubscribersPresentationSuite) TestGetCurrency() {
	req := httptest.NewRequest("GET", "/rates/?from=USD", nil)

	resp, err := suite.webApp.Test(req)

	suite.Require().NoError(err, "Error executing request")
	suite.Equal("200 OK", resp.Status)
}

func TestSubscriberPresenterTestSuite(t *testing.T) {
	suite.Run(t, new(SubscribersPresentationSuite))
}

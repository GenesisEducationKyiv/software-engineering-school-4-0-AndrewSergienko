package app

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/stretchr/testify/suite"
	"go_service/internal/notifier/adapters/scheduler"
	"go_service/internal/notifier/adapters/subscriber"
	"go_service/internal/notifier/infrastructure"
	"go_service/internal/notifier/infrastructure/broker"
	"go_service/internal/notifier/infrastructure/database"
	"gorm.io/gorm"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"testing"
)

type RateNotifierTestSuite struct {
	suite.Suite
	conn        *nats.Conn
	js          jetstream.JetStream
	db          *gorm.DB
	transaction *gorm.DB
	task        RateNotifier
}

func (suite *RateNotifierTestSuite) SetupSuite() {

	databaseSettings := infrastructure.GetDatabaseSettings()

	brokerSettings := infrastructure.GetBrokerSettings()

	conn, js, err := broker.New(brokerSettings)
	if err != nil {
		slog.Error(fmt.Sprintf("NATS is not available. Error: %s", err))
		return
	}

	suite.conn, suite.js = conn, js

	_, err = broker.NewStream(context.Background(), suite.js, "events")
	suite.NoError(err)

	db, err := database.New(databaseSettings)
	suite.NoError(err)

	suite.db = db

	srv := &http.Server{Addr: ":8081", Handler: http.DefaultServeMux} //nolint: all
	http.HandleFunc("/", handler)
	go func() {
		log.Println("Server is starting on port 8081...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()
}

func (suite *RateNotifierTestSuite) SetupTest() {
	ctx := context.Background()

	emailSettings := infrastructure.GetEmailSettings()
	servicesAPISettings := infrastructure.CurrencyRateServiceAPISettings{
		Host:           "http://127.0.0.1:8081",
		GetCurrencyURL: "/",
	}

	suite.transaction = suite.db.Begin()

	schedulerGateway := scheduler.NewScheduleAdapter(nil)
	container := NewIoC(ctx, suite.transaction, &servicesAPISettings, emailSettings, suite.js)

	suite.task = NewRateNotifier(container, schedulerGateway)
}

func (suite *RateNotifierTestSuite) TearDownSuite() {
	broker.Finalize(suite.conn)
}

func (suite *RateNotifierTestSuite) TestRun() {
	email := os.Getenv("EMAIL")

	subsGateway := subscriber.NewSubscriberAdapter(suite.transaction)
	suite.NoError(subsGateway.Create("test1@gmail.com"))
	suite.NoError(subsGateway.Create("test2@gmail.com"))
	suite.NoError(subsGateway.Create("test3@gmail.com"))

	cron := suite.task.Run()

	suite.True(checkMail(email, "test1@gmail.com"))
	suite.True(checkMail(email, "test2@gmail.com"))
	suite.True(checkMail(email, "test3@gmail.com"))

	cron.Stop()
}

func TestRateNotifierTestSuite(t *testing.T) {
	suite.Run(t, new(RateNotifierTestSuite))
}

func handler(w http.ResponseWriter, _ *http.Request) {
	response := map[string]float64{"rates": 50.0}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

type Mailbox struct {
	Mailbox string `json:"Mailbox"`
	Domain  string `json:"Domain"`
}

type EmailAddress struct {
	Mailbox string `json:"Mailbox"`
	Domain  string `json:"Domain"`
	Params  string `json:"Params"`
}

type ContentHeaders struct {
	From       []string `json:"From"`
	MessageID  []string `json:"Message-ID"`
	Received   []string `json:"Received"`
	ReturnPath []string `json:"Return-Path"`
	Subject    []string `json:"Subject"`
	To         []string `json:"To"`
}

type EmailContent struct {
	Headers ContentHeaders `json:"Headers"`
	Body    string         `json:"Body"`
	Size    int            `json:"Size"`
}

type EmailItem struct {
	ID      string         `json:"ID"`
	From    EmailAddress   `json:"From"`
	To      []EmailAddress `json:"To"`
	Content EmailContent   `json:"Content"`
}

type MailhogResponse struct {
	Total int         `json:"total"`
	Count int         `json:"count"`
	Start int         `json:"start"`
	Items []EmailItem `json:"items"`
}

func checkMail(from string, to string) (bool, error) {
	resp, err := http.Get("http://127.0.0.1:8025/api/v2/messages")
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var mailhogResponse MailhogResponse
	if err = json.Unmarshal(body, &mailhogResponse); err != nil {
		return false, err
	}

	for _, item := range mailhogResponse.Items {
		fromHeader := strings.Join(item.Content.Headers.From, ",")
		if strings.Contains(fromHeader, from) {
			toHeader := strings.Join(item.Content.Headers.To, ",")
			if strings.Contains(toHeader, to) {
				return true, nil
			}
		}
	}
	return false, nil
}

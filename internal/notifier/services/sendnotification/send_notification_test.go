package sendnotification

import (
	"errors"
	"github.com/stretchr/testify/suite"
	"go_service/internal/notifier/infrastructure/database/models"
	"testing"
)

type EmailGatewayMock struct {
	EmailsNotified []string
	RaiseError     bool
}

func (eg *EmailGatewayMock) Send(target string, _ float32) error { // nolint: all
	if eg.RaiseError {
		return errors.New("mock error")
	}
	eg.EmailsNotified = append(eg.EmailsNotified, target)
	return nil
}

type SubscriberGatewayMock struct {
	Subscribers []models.Subscriber
}

func (sg *SubscriberGatewayMock) GetAll() []models.Subscriber {
	return sg.Subscribers
}

type CurrencyRateGatewayMock struct {
	RaiseError bool
}

func (cr *CurrencyRateGatewayMock) GetCurrencyRate(_ string, _ string) (float32, error) { // nolint: all
	if cr.RaiseError {
		return 0, errors.New("mock error")
	}
	return 1.5, nil
}

type SendNotificationTestSuite struct {
	suite.Suite
	emailGateway      *EmailGatewayMock
	subscriberGateway *SubscriberGatewayMock
	currencyGateway   *CurrencyRateGatewayMock
}

func (suite *SendNotificationTestSuite) SetupSuite() {
	suite.emailGateway = &EmailGatewayMock{RaiseError: false}
	suite.subscriberGateway = &SubscriberGatewayMock{Subscribers: []models.Subscriber{
		{Email: "test1@gmail.com"},
		{Email: "test2@gmail.com"},
		{Email: "test3@gmail.com"},
	}}
	suite.currencyGateway = &CurrencyRateGatewayMock{RaiseError: false}
}

func (suite *SendNotificationTestSuite) SetupTest() {
	suite.emailGateway.RaiseError = false
	suite.emailGateway.EmailsNotified = nil
	suite.currencyGateway.RaiseError = false
}

func (suite *SendNotificationTestSuite) TestHandle() {
	service := New(suite.emailGateway, suite.subscriberGateway, suite.currencyGateway)
	output := service.Handle(InputData{From: "USD", To: "EUR"})

	suite.NoError(output.Err)
	suite.Len(suite.emailGateway.EmailsNotified, 3)

	for i, subscriber := range suite.subscriberGateway.Subscribers {
		suite.Equal(subscriber.Email, suite.emailGateway.EmailsNotified[i])
	}
}

func (suite *SendNotificationTestSuite) TestHandle_ErrorEmails() {
	suite.emailGateway.RaiseError = true

	service := New(suite.emailGateway, suite.subscriberGateway, suite.currencyGateway)
	output := service.Handle(InputData{From: "USD", To: "EUR"})

	suite.NoError(output.Err)
	suite.Len(output.ErrEmails, 3)
	suite.Len(suite.emailGateway.EmailsNotified, 0)

	for i, subscriber := range suite.subscriberGateway.Subscribers {
		suite.Equal(subscriber.Email, output.ErrEmails[i])
	}
}

func (suite *SendNotificationTestSuite) TestHandle_ErrorRate() {
	suite.currencyGateway.RaiseError = true

	service := New(suite.emailGateway, suite.subscriberGateway, suite.currencyGateway)
	output := service.Handle(InputData{From: "USD", To: "EUR"})

	suite.Error(output.Err)
	suite.Len(suite.emailGateway.EmailsNotified, 0)
}

func TestSendNotificationSuite(t *testing.T) {
	suite.Run(t, new(SendNotificationTestSuite))
}

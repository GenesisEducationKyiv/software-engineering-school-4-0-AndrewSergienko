package sendnotification

import (
	"github.com/stretchr/testify/suite"
	"go_service/internal/notifier/infrastructure/database/models"
	"testing"
)

type EmailGatewayMock struct {
	EmailsNotified []string
}

func (eg *EmailGatewayMock) Send(target string, rate float32) error { // nolint: all
	eg.EmailsNotified = append(eg.EmailsNotified, target)
	return nil
}

type SubscriberGatewayMock struct {
	Subscribers []models.Subscriber
}

func (sg *SubscriberGatewayMock) GetAll() []models.Subscriber {
	return sg.Subscribers
}

type CurrencyRateGatewayMock struct{}

func (cr *CurrencyRateGatewayMock) GetCurrencyRate(from string, to string) (float32, error) { // nolint: all
	return 1.5, nil
}

type SendNotificationTestSuite struct {
	suite.Suite
	emailGateway      *EmailGatewayMock
	subscriberGateway *SubscriberGatewayMock
	currencyGateway   *CurrencyRateGatewayMock
}

func (suite *SendNotificationTestSuite) SetupSuite() {
	suite.emailGateway = &EmailGatewayMock{}
	suite.subscriberGateway = &SubscriberGatewayMock{Subscribers: []models.Subscriber{
		{Email: "test1@gmail.com"},
		{Email: "test2@gmail.com"},
		{Email: "test3@gmail.com"},
	}}
	suite.currencyGateway = &CurrencyRateGatewayMock{}
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

func TestSendNotificationSuite(t *testing.T) {
	suite.Run(t, new(SendNotificationTestSuite))
}

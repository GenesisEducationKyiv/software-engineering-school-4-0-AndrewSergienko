package createsubscriber

import (
	"github.com/stretchr/testify/suite"
	"go_service/internal/notifier/infrastructure/database/models"
	"testing"
)

type SubscriberGatewayMock struct {
	Subscriber *models.Subscriber
}

func (mock *SubscriberGatewayMock) Create(email string) error {
	mock.Subscriber = &models.Subscriber{Email: email}
	return nil
}

func (mock *SubscriberGatewayMock) GetByEmail(email string) *models.Subscriber {
	if mock.Subscriber != nil && mock.Subscriber.Email == email {
		return mock.Subscriber
	}
	return nil
}

type CreateSubscriberTestSuite struct {
	suite.Suite
	subscriberGateway SubscriberGateway
}

func (suite *CreateSubscriberTestSuite) SetupSuite() {
	suite.subscriberGateway = &SubscriberGatewayMock{}
}

func (suite *CreateSubscriberTestSuite) TestHandle() {
	service := New(suite.subscriberGateway)
	data := InputData{Email: "test@gmail.com"}
	suite.NoError(service.Handle(data).Err)
	suite.NotNil(suite.subscriberGateway.GetByEmail(data.Email))
}

func TestCreateSubscriberSuite(t *testing.T) {
	suite.Run(t, new(CreateSubscriberTestSuite))
}

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

type EventEmitterMock struct {
	EmitCalled bool
}

func (e EventEmitterMock) Emit(_ string, _ map[string]interface{}, _ *string) error {
	e.EmitCalled = true // nolint: all
	return nil
}

type CreateSubscriberTestSuite struct {
	suite.Suite
	subscriberGateway SubscriberGateway
	eventEmitter      EventEmitter
}

func (suite *CreateSubscriberTestSuite) SetupSuite() {
	suite.subscriberGateway = &SubscriberGatewayMock{}
	suite.eventEmitter = &EventEmitterMock{EmitCalled: false}
}

func (suite *CreateSubscriberTestSuite) TestHandle() {
	service := New(suite.subscriberGateway, suite.eventEmitter)
	data := InputData{Email: "test@gmail.com"}
	suite.NoError(service.Handle(data).Err)
	suite.NotNil(suite.subscriberGateway.GetByEmail(data.Email))
	suite.True(suite.eventEmitter.(*EventEmitterMock).EmitCalled)
}

func TestCreateSubscriberSuite(t *testing.T) {
	suite.Run(t, new(CreateSubscriberTestSuite))
}

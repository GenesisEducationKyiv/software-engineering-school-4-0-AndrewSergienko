package createsubscriber

import (
	"errors"
	"github.com/stretchr/testify/suite"
	"go_service/internal/notifier/infrastructure/database/models"
	"testing"
)

type SubscriberGatewayMock struct {
	RaiseError bool
	Subscriber *models.Subscriber
}

func (mock *SubscriberGatewayMock) Create(email string) error {
	if mock.RaiseError {
		return errors.New("mock error")
	}

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
	EmittedEvent string
}

func (e *EventEmitterMock) Emit(event string, _ map[string]interface{}, _ *string) error {
	e.EmittedEvent = event
	return nil
}

type CreateSubscriberTestSuite struct {
	suite.Suite
	subscriberGateway SubscriberGateway
	eventEmitter      EventEmitter
}

func (suite *CreateSubscriberTestSuite) SetupSuite() {
	suite.subscriberGateway = &SubscriberGatewayMock{}
	suite.eventEmitter = &EventEmitterMock{}
}

func (suite *CreateSubscriberTestSuite) TestHandle_Success() {
	suite.subscriberGateway.(*SubscriberGatewayMock).RaiseError = false

	service := New(suite.subscriberGateway, suite.eventEmitter)
	data := InputData{Email: "test@gmail.com"}
	suite.NoError(service.Handle(data).Err)
	suite.NotNil(suite.subscriberGateway.GetByEmail(data.Email))
	suite.Equal("SubscriberCreated", suite.eventEmitter.(*EventEmitterMock).EmittedEvent)
}

func (suite *CreateSubscriberTestSuite) TestHandle_Error() {
	suite.subscriberGateway.(*SubscriberGatewayMock).RaiseError = true

	service := New(suite.subscriberGateway, suite.eventEmitter)
	data := InputData{Email: "test@gmail.com"}
	suite.NoError(service.Handle(data).Err)
	suite.Nil(suite.subscriberGateway.GetByEmail(data.Email))
	suite.Equal("SubscriberCreatedError", suite.eventEmitter.(*EventEmitterMock).EmittedEvent)
}

func TestCreateSubscriberSuite(t *testing.T) {
	suite.Run(t, new(CreateSubscriberTestSuite))
}

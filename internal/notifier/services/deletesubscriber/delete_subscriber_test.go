package deletesubscriber

import (
	"errors"
	"github.com/stretchr/testify/suite"
	"go_service/internal/notifier/infrastructure/database/models"
	"testing"
)

type SubscriberGatewayMock struct {
	subscriber        models.Subscriber
	subscriberDeleted bool
	RaiseError        bool
}

func (mock *SubscriberGatewayMock) Delete(email string) error {
	if mock.RaiseError {
		return errors.New("mock error")
	}

	if mock.subscriber.Email == email {
		mock.subscriberDeleted = true
	}
	return nil
}

type EventEmitterMock struct {
	EmittedEvent string
}

func (e *EventEmitterMock) Emit(event string, _ map[string]interface{}, _ *string) error {
	e.EmittedEvent = event // nolint: all
	return nil
}

type DeleteSubscriberTestSuite struct {
	suite.Suite
	subscriberGateway SubscriberGateway
	eventEmitter      EventEmitter
}

func (suite *DeleteSubscriberTestSuite) SetupSuite() {
	suite.subscriberGateway = &SubscriberGatewayMock{subscriber: models.Subscriber{
		Email: "test@gmail.com",
	}, subscriberDeleted: false, RaiseError: false}
	suite.eventEmitter = &EventEmitterMock{}
}

func (suite *DeleteSubscriberTestSuite) TestHandle_Success() {
	suite.subscriberGateway.(*SubscriberGatewayMock).RaiseError = false

	service := New(suite.subscriberGateway, suite.eventEmitter)
	suite.NoError(service.Handle(InputData{Email: "test@gmail.com"}).Err)
	suite.True(suite.subscriberGateway.(*SubscriberGatewayMock).subscriberDeleted)
	suite.Equal("SubscriberDeleted", suite.eventEmitter.(*EventEmitterMock).EmittedEvent)
}

func (suite *DeleteSubscriberTestSuite) TestHandle_Error() {
	suite.subscriberGateway.(*SubscriberGatewayMock).RaiseError = true

	service := New(suite.subscriberGateway, suite.eventEmitter)
	suite.Error(service.Handle(InputData{Email: "test@gmail.com"}).Err)
	suite.False(suite.subscriberGateway.(*SubscriberGatewayMock).subscriberDeleted)
	suite.Equal("SubscriberDeletedError", suite.eventEmitter.(*EventEmitterMock).EmittedEvent)
}

func TestCreateSubscriberSuite(t *testing.T) {
	suite.Run(t, new(DeleteSubscriberTestSuite))
}

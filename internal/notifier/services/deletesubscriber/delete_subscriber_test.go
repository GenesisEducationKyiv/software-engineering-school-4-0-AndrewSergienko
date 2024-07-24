package deletesubscriber

import (
	"github.com/stretchr/testify/suite"
	"go_service/internal/notifier/infrastructure/database/models"
	"testing"
)

type SubscriberGatewayMock struct {
	subscriber        models.Subscriber
	subscriberDeleted bool
}

func (mock *SubscriberGatewayMock) Delete(email string) error {
	if mock.subscriber.Email == email {
		mock.subscriberDeleted = true
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

type DeleteSubscriberTestSuite struct {
	suite.Suite
	subscriberGateway SubscriberGateway
	eventEmitter      EventEmitter
}

func (suite *DeleteSubscriberTestSuite) SetupSuite() {
	suite.subscriberGateway = &SubscriberGatewayMock{subscriber: models.Subscriber{
		Email: "test@gmail.com",
	}, subscriberDeleted: false}
	suite.eventEmitter = &EventEmitterMock{EmitCalled: false}
}

func (suite *DeleteSubscriberTestSuite) TestHandle() {
	service := New(suite.subscriberGateway, suite.eventEmitter)
	suite.NoError(service.Handle(InputData{Email: "test@gmail.com"}).Err)
	suite.True(suite.subscriberGateway.(*SubscriberGatewayMock).subscriberDeleted)
	suite.True(suite.eventEmitter.(*EventEmitterMock).EmitCalled)
}

func TestCreateSubscriberSuite(t *testing.T) {
	suite.Run(t, new(DeleteSubscriberTestSuite))
}

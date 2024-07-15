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

type DeleteSubscriberTestSuite struct {
	suite.Suite
	subscriberGateway SubscriberGateway
}

func (suite *DeleteSubscriberTestSuite) SetupSuite() {
	suite.subscriberGateway = &SubscriberGatewayMock{subscriber: models.Subscriber{
		Email: "test@gmail.com",
	}, subscriberDeleted: false}
}

func (suite *DeleteSubscriberTestSuite) TestHandle() {
	service := NewDeleteSubscriber(suite.subscriberGateway)
	suite.NoError(service.Handle(InputData{Email: "test@gmail.com"}).Err)
	suite.True(suite.subscriberGateway.(*SubscriberGatewayMock).subscriberDeleted)
}

func TestCreateSubscriberSuite(t *testing.T) {
	suite.Run(t, new(DeleteSubscriberTestSuite))
}

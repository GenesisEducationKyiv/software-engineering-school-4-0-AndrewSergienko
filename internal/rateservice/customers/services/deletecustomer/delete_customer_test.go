package deletecustomer

import (
	"errors"
	"github.com/stretchr/testify/suite"
	"go_service/internal/rateservice/customers/infrastructure/database/models"
	"testing"
)

type CustomerGatewayMock struct {
	RaiseError      bool
	CreatedCustomer *models.Customer
}

func (mock *CustomerGatewayMock) GetByEmail(email string) *models.Customer {
	if mock.CreatedCustomer != nil && mock.CreatedCustomer.Email == email {
		return mock.CreatedCustomer
	}
	return nil
}

func (mock *CustomerGatewayMock) DeleteByEmail(email string) error {
	if mock.RaiseError {
		return errors.New("mock error")
	}

	if mock.CreatedCustomer.Email == email {
		mock.CreatedCustomer = nil
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

type CreateCustomerTestSuite struct {
	suite.Suite
	customerGateway CustomerGateway
	eventEmitter    EventEmitter
}

func (suite *CreateCustomerTestSuite) SetupTest() {
	suite.customerGateway = &CustomerGatewayMock{
		CreatedCustomer: &models.Customer{Email: "test@gmail.com"},
		RaiseError:      false,
	}
	suite.eventEmitter = &EventEmitterMock{}
}

func (suite *CreateCustomerTestSuite) TestHandle_Success() {
	suite.customerGateway.(*CustomerGatewayMock).RaiseError = false

	service := New(suite.customerGateway, suite.eventEmitter)
	data := InputData{Email: "test@gmail.com"}
	suite.NoError(service.Handle(data).Err)
	suite.Nil(suite.customerGateway.GetByEmail(data.Email))
	suite.Equal("UserDeleted", suite.eventEmitter.(*EventEmitterMock).EmittedEvent)
}

func (suite *CreateCustomerTestSuite) TestHandle_Error() {
	suite.customerGateway.(*CustomerGatewayMock).RaiseError = true

	service := New(suite.customerGateway, suite.eventEmitter)
	suite.Error(service.Handle(InputData{Email: "test@gmail.com"}).Err)
	suite.NotNil(suite.customerGateway.GetByEmail("test@gmail.com"))
	suite.Equal("UserDeletedError", suite.eventEmitter.(*EventEmitterMock).EmittedEvent)
}

func TestCreateCustomerSuite(t *testing.T) {
	suite.Run(t, new(CreateCustomerTestSuite))
}

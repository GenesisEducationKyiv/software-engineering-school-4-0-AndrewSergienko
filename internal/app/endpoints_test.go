package app

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
	"go_service/internal/infrastructure"
	"go_service/internal/infrastructure/database"
	"go_service/internal/subscribers/adapters"
	"gorm.io/gorm"
	"net/http/httptest"
	"testing"
)

type SubscribersPresentationSuite struct {
	suite.Suite
	db                *gorm.DB
	transaction       *gorm.DB
	webApp            *fiber.App
	subscriberGateway *adapters.SubscriberAdapter
}

func (suite *SubscribersPresentationSuite) SetupSuite() {
	databaseSettings := infrastructure.GetDatabaseSettings()
	suite.db = database.InitDatabase(databaseSettings)

}

func (suite *SubscribersPresentationSuite) SetupTest() {
	suite.transaction = suite.db.Begin()
	currencyAPISettings := infrastructure.GetCurrencyAPISettings()

	suite.subscriberGateway = adapters.NewSubscriberAdapter(suite.transaction)
	suite.webApp = InitWebApp(suite.transaction, currencyAPISettings)
}

func (suite *SubscribersPresentationSuite) TearDownTest() {
	suite.db.Rollback()
}

func (suite *SubscribersPresentationSuite) TestAddSubscriber() {
	var jsonStr = []byte(`{"email":"test@gmail.com"}`)

	req := httptest.NewRequest("POST", "/subscribers/", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp, err := suite.webApp.Test(req)
	suite.Require().NoError(err, "Error executing request")

	suite.Require().Equal("200 OK", resp.Status)

	suite.NotNil(suite.subscriberGateway.GetByEmail("test@gmail.com"))
}

func (suite *SubscribersPresentationSuite) TestGetCurrency() {
	req := httptest.NewRequest("GET", "/rates/?from=USD", nil)

	resp, err := suite.webApp.Test(req)

	suite.Require().NoError(err, "Error executing request")
	suite.Equal("200 OK", resp.Status)
}

func TestSubscriberPresenterTestSuite(t *testing.T) {
	suite.Run(t, new(SubscribersPresentationSuite))
}

package app

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
	"go_service/internal/adapters"
	"go_service/internal/infrastructure"
	"go_service/internal/infrastructure/database"
	"go_service/internal/presentation"
	"gorm.io/gorm"
	"net/http/httptest"
	"testing"
)

type SubscribersPresentationSuite struct {
	suite.Suite
	database          *gorm.DB
	subscriberGateway presentation.SubscriberGateway
	webApp            *fiber.App
}

func (suite *SubscribersPresentationSuite) SetupSuite() {
	databaseSettings := infrastructure.GetDatabaseSettings()
	currencyAPISettings := infrastructure.GetCurrencyAPISettings()
	suite.database = database.InitDatabase(databaseSettings)

	currencyGateway := adapters.GetAPICurrencyReader(currencyAPISettings)
	subscriberGateway := adapters.GetSubscribersAdapter(suite.database)

	suite.subscriberGateway = subscriberGateway
	suite.webApp = InitWebApp(currencyGateway, subscriberGateway)
}

func (suite *SubscribersPresentationSuite) TearDownTest() {
	infrastructure.ClearDB(suite.database)
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
	req := httptest.NewRequest("GET", "/", nil)

	resp, err := suite.webApp.Test(req)

	suite.Require().NoError(err, "Error executing request")
	suite.Equal("200 OK", resp.Status)
}

func TestSubscriberPresenterTestSuite(t *testing.T) {
	suite.Run(t, new(SubscribersPresentationSuite))
}

package app

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
	"go.uber.org/dig"
	"go_service/internal/adapters"
	"go_service/internal/infrastructure"
	"go_service/internal/infrastructure/database"
	"gorm.io/gorm"
	"net/http/httptest"
	"testing"
)

type SubscribersPresentationSuite struct {
	suite.Suite
	db          *gorm.DB
	transaction *gorm.DB
	container   *dig.Container
	webApp      *fiber.App
}

func (suite *SubscribersPresentationSuite) SetupSuite() {
	databaseSettings := infrastructure.GetDatabaseSettings()
	suite.db = database.InitDatabase(databaseSettings)

}

func (suite *SubscribersPresentationSuite) SetupTest() {
	suite.transaction = suite.db.Begin()
	currencyAPISettings := infrastructure.GetCurrencyAPISettings()
	emailSettings := infrastructure.EmailSettings{}

	container := SetupContainer(suite.transaction, emailSettings, currencyAPISettings)

	suite.webApp = InitWebApp(container)
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

	suite.container.Invoke(func(gateway adapters.SubscriberAdapter) {
		suite.NotNil(gateway.GetByEmail("test@gmail.com"))
	})

}

func (suite *SubscribersPresentationSuite) TestGetCurrency() {
	req := httptest.NewRequest("GET", "/?from=USD", nil)

	resp, err := suite.webApp.Test(req)

	suite.Require().NoError(err, "Error executing request")
	suite.Equal("200 OK", resp.Status)
}

func TestSubscriberPresenterTestSuite(t *testing.T) {
	suite.Run(t, new(SubscribersPresentationSuite))
}

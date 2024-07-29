package adapters

import (
	"github.com/stretchr/testify/suite"
	"go_service/internal/rateservice/customers/infrastructure/database/models"
	"go_service/internal/rateservice/infrastructure"
	"go_service/internal/rateservice/infrastructure/database"
	"gorm.io/gorm"
	"testing"
)

type SubscriberAdapterTestSuite struct {
	suite.Suite
	db          *gorm.DB
	transaction *gorm.DB
	adapter     *SubscriberAdapter
}

func (suite *SubscriberAdapterTestSuite) SetupSuite() {
	settings := infrastructure.GetDatabaseSettings()
	suite.db = database.New(settings)
}

func (suite *SubscriberAdapterTestSuite) SetupTest() {
	suite.transaction = suite.db.Begin()
	suite.adapter = NewSubscriberAdapter(suite.transaction)
}

func (suite *SubscriberAdapterTestSuite) TearDownTest() {
	suite.transaction.Rollback()
}

func (suite *SubscriberAdapterTestSuite) TestGetByEmail() {
	email := "test@gmail.com"
	suite.transaction.Create(&models.Customer{Email: email})

	subscriber := suite.adapter.GetByEmail(email)

	suite.NotNil(subscriber)
}

func (suite *SubscriberAdapterTestSuite) TestCreate() {
	email := "test@gmail.com"

	err := suite.adapter.Create(email)

	suite.Nil(err)

	subscriber := suite.adapter.GetByEmail(email)
	suite.NotNil(subscriber)
}

func (suite *SubscriberAdapterTestSuite) TestCreateDuplicate() {
	email := "test@gmail.com"
	suite.transaction.Create(&models.Customer{Email: email})

	err := suite.adapter.Create(email)
	suite.NotNil(err)
}

func TestSubscriberAdapterSuite(t *testing.T) {
	suite.Run(t, new(SubscriberAdapterTestSuite))
}

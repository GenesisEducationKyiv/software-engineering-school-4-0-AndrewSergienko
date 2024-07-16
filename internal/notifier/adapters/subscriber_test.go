package adapters

import (
	"github.com/stretchr/testify/suite"
	"go_service/internal/notifier/infrastructure"
	"go_service/internal/notifier/infrastructure/database"
	"go_service/internal/notifier/infrastructure/database/models"
	"gorm.io/gorm"
	"testing"
)

type SubscriberAdapterTestSuite struct {
	suite.Suite
	db          *gorm.DB
	transaction *gorm.DB
	adapter     SubscriberAdapter
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

func (suite *SubscriberAdapterTestSuite) TestGetAll() {
	emails := []string{"test1@gmail.com", "test2@gmail.com", "test3@gmail.com"}
	for _, email := range emails {
		suite.transaction.Create(&models.Subscriber{Email: email})
	}

	subscribers := suite.adapter.GetAll()

	suite.Equal(len(emails), len(subscribers))

	for i := 0; i < len(emails); i++ {
		suite.Equal(emails[i], subscribers[i].Email)
	}
}

func TestSubscriberAdapterSuite(t *testing.T) {
	suite.Run(t, new(SubscriberAdapterTestSuite))
}

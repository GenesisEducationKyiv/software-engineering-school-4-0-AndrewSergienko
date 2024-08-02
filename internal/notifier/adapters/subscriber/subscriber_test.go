package subscriber

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
	adapter     Adapter
}

func (suite *SubscriberAdapterTestSuite) SetupSuite() {
	settings := infrastructure.GetDatabaseSettings()
	db, err := database.New(settings)
	suite.NoError(err)

	suite.db = db
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

func (suite *SubscriberAdapterTestSuite) TestGetByEmail() {
	email := "test@gmail.com"
	suite.transaction.Create(&models.Subscriber{Email: email})

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
	suite.transaction.Create(&models.Subscriber{Email: email})

	err := suite.adapter.Create(email)
	suite.NotNil(err)
}

func TestSubscriberAdapterSuite(t *testing.T) {
	suite.Run(t, new(SubscriberAdapterTestSuite))
}

package adapters

import (
	"github.com/stretchr/testify/suite"
	"go_service/internal/infrastructure"
	"go_service/internal/infrastructure/database"
	"go_service/internal/infrastructure/database/models"
	"gorm.io/gorm"
	"testing"
)

type SubscriberAdapterTestSuite struct {
	suite.Suite
	database *gorm.DB
	adapter  *SubscriberAdapter
}

func (suite *SubscriberAdapterTestSuite) SetupSuite() {
	settings := infrastructure.GetDatabaseSettings()
	suite.database = database.InitDatabase(settings)
	suite.adapter = GetSubscribersAdapter(suite.database)
}

func (suite *SubscriberAdapterTestSuite) TearDownTest() {
	infrastructure.ClearDB(suite.database)
}

func (suite *SubscriberAdapterTestSuite) TestGetByEmail() {
	email := "test@gmail.com"
	suite.database.Create(&models.Subscriber{Email: email})

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
	suite.database.Create(&models.Subscriber{Email: email})

	err := suite.adapter.Create(email)
	suite.NotNil(err)
}

func (suite *SubscriberAdapterTestSuite) TestGetAll() {
	emails := []string{"test1@gmail.com", "test2@gmail.com", "test3@gmail.com"}
	for _, email := range emails {
		suite.database.Create(&models.Subscriber{Email: email})
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

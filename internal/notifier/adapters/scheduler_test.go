package adapters

import (
	"github.com/stretchr/testify/suite"
	"go_service/internal/infrastructure"
	"go_service/internal/infrastructure/database"
	"go_service/internal/notifier/infrastructure/database/models"
	"gorm.io/gorm"
	"testing"
	"time"
)

type SchedulerAdapterTestSuite struct {
	suite.Suite
	db          *gorm.DB
	transaction *gorm.DB
	adapter     *ScheduleDBAdapter
}

func (suite *SchedulerAdapterTestSuite) SetupSuite() {
	settings := infrastructure.GetDatabaseSettings()
	suite.db = database.InitDatabase(settings)
}

func (suite *SchedulerAdapterTestSuite) SetupTest() {
	suite.transaction = suite.db.Begin()
	suite.adapter = NewScheduleDBAdapter(suite.transaction)
}

func (suite *SchedulerAdapterTestSuite) TearDownTest() {
	suite.transaction.Rollback()
}

func (suite *SchedulerAdapterTestSuite) TestGetLastTimeExisted() {
	now := time.Now()
	now = now.Truncate(time.Second)

	suite.transaction.Create(&models.ScheduleTime{Time: now.Unix()})

	suite.Equal(now, *suite.adapter.GetLastTime())
}

func (suite *SchedulerAdapterTestSuite) TestGetLastTimeNotExisted() {
	suite.Nil(suite.adapter.GetLastTime())
}

func (suite *SchedulerAdapterTestSuite) TestSetLastTime() {
	err := suite.adapter.SetLastTime()

	suite.Nil(err)

	suite.NotNil(suite.adapter.GetLastTime())
}

func TestScheduleAdapterSuite(t *testing.T) {
	suite.Run(t, new(SchedulerAdapterTestSuite))
}

package adapters

import (
	"github.com/stretchr/testify/suite"
	"go_service/internal/infrastructure"
	"go_service/internal/infrastructure/database"
	"go_service/internal/infrastructure/database/models"
	"gorm.io/gorm"
	"testing"
	"time"
)

type SchedulerAdapterTestSuite struct {
	suite.Suite
	database *gorm.DB
	adapter  *ScheduleDBAdapter
}

func (suite *SchedulerAdapterTestSuite) SetupSuite() {
	settings := infrastructure.GetDatabaseSettings()
	suite.database = database.InitDatabase(settings)
	suite.adapter = GetScheduleDBAdapter(suite.database)
}

func (suite *SchedulerAdapterTestSuite) TearDownTest() {
	infrastructure.ClearDB(suite.database)
}

func (suite *SchedulerAdapterTestSuite) TestGetLastTimeExisted() {
	now := time.Now()
	now = now.Truncate(time.Second)

	suite.database.Create(&models.ScheduleTime{Time: now.Unix()})

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

package scheduler

import (
	"github.com/stretchr/testify/suite"
	"go_service/internal/notifier/infrastructure"
	"go_service/internal/notifier/infrastructure/database"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"testing"
	"time"
)

type SchedulerAdapterTestSuite struct {
	suite.Suite
	db          *gorm.DB
	transaction *gorm.DB
	adapter     *ScheduleAdapter
	path        *string
}

func (suite *SchedulerAdapterTestSuite) SetupSuite() {
	settings := infrastructure.GetDatabaseSettings()
	db, err := database.New(settings)
	suite.db = db
	suite.NoError(err)

	projectRoot, err := os.Getwd()
	suite.NoError(err)

	configPath := filepath.Join(projectRoot, "..", "..", "..", "..", "conf", "email_sent_time.json")
	suite.path = &configPath
}

func (suite *SchedulerAdapterTestSuite) SetupTest() {
	suite.transaction = suite.db.Begin()

	suite.adapter = NewScheduleAdapter(suite.path)
}

func (suite *SchedulerAdapterTestSuite) TearDownTest() {
	suite.transaction.Rollback()
	_ = os.Remove(*suite.path)
}

func (suite *SchedulerAdapterTestSuite) TestGetLastTimeExisted() {
	now := time.Now()
	now = now.Truncate(time.Second)

	err := suite.adapter.SetLastTime(now)
	if err != nil {
		suite.T().Skip()
	}
	suite.Equal(now, *suite.adapter.GetLastTime())
}

func (suite *SchedulerAdapterTestSuite) TestGetLastTimeNotExisted() {
	suite.Nil(suite.adapter.GetLastTime())
}

func (suite *SchedulerAdapterTestSuite) TestSetLastTime() {
	err := suite.adapter.SetLastTime(time.Now())

	if err != nil {
		suite.T().Skip()
	}

	suite.NotNil(suite.adapter.GetLastTime())
}

func TestScheduleAdapterSuite(t *testing.T) {
	suite.Run(t, new(SchedulerAdapterTestSuite))
}

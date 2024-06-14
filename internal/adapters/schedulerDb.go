package adapters

import (
	"go_service/internal/infrastructure/database/models"
	"gorm.io/gorm"
	"time"
)

type SchedulerDbAdapter struct {
	Db *gorm.DB
}

func GetSchedulerDbAdapter(db *gorm.DB) *SchedulerDbAdapter {
	return &SchedulerDbAdapter{Db: db}
}

func (sa SchedulerDbAdapter) GetLastTime() *time.Time {
	var lastTime models.SchedulerTime
	result := sa.Db.Last(&lastTime)
	if result.Error != nil {
		return nil
	}
	lastTimeUnix := time.Unix(lastTime.Time, 0)
	return &lastTimeUnix
}

func (sa SchedulerDbAdapter) SetLastTime() error {
	schedulerTime := models.SchedulerTime{Time: time.Now().Unix()}
	result := sa.Db.Create(&schedulerTime)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

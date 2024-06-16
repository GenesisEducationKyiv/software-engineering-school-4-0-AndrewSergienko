package adapters

import (
	"go_service/internal/infrastructure/database/models"
	"gorm.io/gorm"
	"time"
)

type SchedulerDBAdapter struct {
	db *gorm.DB
}

func GetSchedulerDBAdapter(db *gorm.DB) *SchedulerDBAdapter {
	return &SchedulerDBAdapter{db: db}
}

func (sa SchedulerDBAdapter) GetLastTime() *time.Time {
	var lastTime models.SchedulerTime
	result := sa.db.Last(&lastTime)
	if result.Error != nil {
		return nil
	}
	lastTimeUnix := time.Unix(lastTime.Time, 0)
	return &lastTimeUnix
}

func (sa SchedulerDBAdapter) SetLastTime() error {
	schedulerTime := models.SchedulerTime{Time: time.Now().Unix()}
	result := sa.db.Create(&schedulerTime)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

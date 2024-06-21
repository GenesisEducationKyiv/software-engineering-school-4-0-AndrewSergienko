package adapters

import (
	"go_service/internal/infrastructure/database/models"
	"gorm.io/gorm"
	"time"
)

type ScheduleDBAdapter struct {
	db *gorm.DB
}

func GetScheduleDBAdapter(db *gorm.DB) *ScheduleDBAdapter {
	return &ScheduleDBAdapter{db: db}
}

func (sa ScheduleDBAdapter) GetLastTime() *time.Time {
	var lastTime models.ScheduleTime
	result := sa.db.Last(&lastTime)
	if result.Error != nil {
		return nil
	}
	lastTimeUnix := time.Unix(lastTime.Time, 0)
	return &lastTimeUnix
}

func (sa ScheduleDBAdapter) SetLastTime() error {
	schedulerTime := models.ScheduleTime{Time: time.Now().Unix()}
	result := sa.db.Create(&schedulerTime)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

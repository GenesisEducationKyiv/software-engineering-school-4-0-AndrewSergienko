package scheduler

import (
	"encoding/json"
	"os"
	"time"
)

type TimeData struct {
	LastSentTime string `json:"last_sent_time"`
}

type ScheduleAdapter struct {
	filename string
}

func NewScheduleAdapter(path *string) *ScheduleAdapter {
	if path != nil {
		return &ScheduleAdapter{filename: *path}
	}
	return &ScheduleAdapter{filename: "conf/email_sent_time.json"}
}

func (sa ScheduleAdapter) GetLastTime() *time.Time {
	data, err := readJSON(sa.filename)
	if err != nil {
		return nil
	}

	lastSentTime, err := time.Parse(time.RFC3339, data.LastSentTime)
	if err != nil {
		return nil
	}
	return &lastSentTime
}

func (sa ScheduleAdapter) SetLastTime(lastTime time.Time) error {
	data := TimeData{LastSentTime: lastTime.Format(time.RFC3339)}
	return writeJSON(sa.filename, data)
}

func readJSON(filename string) (TimeData, error) {
	var data TimeData
	file, err := os.ReadFile(filename)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(file, &data)
	return data, err
}

func writeJSON(filename string, data TimeData) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, file, 0600)
	return err
}

package models

type Subscriber struct {
	Email string `gorm:"unique"`
}

func (Subscriber) TableName() string {
	return "subscribers"
}

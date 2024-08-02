package models

type Customer struct {
	Email string `gorm:"unique"`
}

func (Customer) TableName() string {
	return "subscribers"
}

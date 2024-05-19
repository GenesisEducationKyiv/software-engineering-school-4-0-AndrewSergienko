package src

type Subscriber struct {
	Email string `gorm:"unique"`
}

package src

type CurrencyReader interface {
	GetUSDCurrencyRate() (float32, error)
}

type SubscriberReader interface {
	GetByEmail(email string) *Subscriber
}

type SubscriberWriter interface {
	Create(email string) error
}

type SubscriberDeleter interface {
	Delete(id int)
}

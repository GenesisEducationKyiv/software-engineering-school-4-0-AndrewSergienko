package presentation

type Interactor[T, U any] interface {
	Handle(data T) (U, error)
}

type InteractorFactory interface {
	GetCurrencyRate() Interactor[string, float32]
	Subscribe() Interactor[string, error]
}

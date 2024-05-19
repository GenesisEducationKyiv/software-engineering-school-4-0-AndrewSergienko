package src

type CurrencyReader interface {
	getUSDCurrencyRate() (float32, error)
}

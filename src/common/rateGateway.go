package common

type CurrencyReader interface {
	GetUSDCurrencyRate() (float32, error)
}

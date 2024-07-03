package services

import "fmt"

type CurrencyNotExistsError struct {
	Currency string
	Source   string
}

func (e *CurrencyNotExistsError) Error() string {
	return fmt.Sprintf("Currency %s not exists in %s", e.Currency, e.Source)
}

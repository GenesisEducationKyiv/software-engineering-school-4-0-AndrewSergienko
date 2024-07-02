package internal

import "fmt"

type EmailConflictError struct {
	Email string
}

func (e *EmailConflictError) Error() string {
	return fmt.Sprintf("Email %s already exists", e.Email)
}

type CurrencyNotExistsError struct {
	Currency string
	Source   string
}

func (e *CurrencyNotExistsError) Error() string {
	return fmt.Sprintf("Currency %s not exists in %s", e.Currency, e.Source)
}

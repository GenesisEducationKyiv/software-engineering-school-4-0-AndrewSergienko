package domain

import "fmt"

type EmailConflictError struct {
	Email string
}

func (e *EmailConflictError) Error() string {
	return fmt.Sprintf("Email %s already exists", e.Email)
}

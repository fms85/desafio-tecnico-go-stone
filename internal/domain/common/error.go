package common

import "fmt"

const (
	NOT_FOUND_ERROR = "not found"
	FOUND_ERROR     = "already exists"
)

type ValidationError struct {
	Msg string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf(e.Msg)
}

package gottbot

import (
	"fmt"
)

// Error Server returns this if there was an exception to your request
type Error struct {
	// Code Error code
	Code string `json:"code"`

	// Error Error
	ErrorString *string `json:"error,omitempty"`

	// Message Human-readable description
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("An error occured with the code '%s' due to '%s'", e.Code, e.Message)
}

func (e *Error) Unwrap() error {
	if e.ErrorString != nil {
		return fmt.Errorf(*e.ErrorString)
	}
	return nil
}

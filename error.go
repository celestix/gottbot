package gottbot

import (
	"fmt"
)

var (
	AttachmentNotReadyError = &Error{
		Code:    "attachment.not.ready",
		Message: "Key: errors.process.attachment.file.not.processed",
	}
	InvalidPhotoPayloadError = &Error{
		Code:    "proto.payload",
		Message: "No `photos`, `url` or `token` provided. Check payload.",
	}
)

// Error Server returns this if there was an exception to your request
type Error struct {
	// Code Error code
	Code string `json:"code"`

	// Message Human-readable description
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("An error occured with the code '%s' due to '%s'", e.Code, e.Message)
}

func EqErrors(err error, target *Error) bool {
	unwrappedErr, ok := err.(*Error)
	if !ok {
		return false
	}
	if target.Message == "" {
		return unwrappedErr.Code == target.Code
	}
	return unwrappedErr.Code == target.Code && unwrappedErr.Message == target.Message
}

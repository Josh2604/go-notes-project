package apierrors

import (
	"net/http"
	"strings"
)

const (
	MethodNotAllowed    = "Method not allowed"
	InternalServerError = "Internal Server Error"
	BadRequest          = "Bad request"
)

func NewErrorMessage(code int, message string, err string) *ErrorMessage {
	return &ErrorMessage{
		Code:    code,
		Message: message,
		Error:   err,
	}
}

func NewBadRequest(messages ...string) *ErrorMessage {
	message := BadRequest
	if len(message) > 0 {
		message += strings.Join(messages, " - ")
	}
	return NewErrorMessage(http.StatusBadRequest, message, "BAD_REQUEST")
}

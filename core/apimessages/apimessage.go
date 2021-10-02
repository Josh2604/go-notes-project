package apimessages

import (
	"net/http"
	"strings"
)

const (
	Created = "Resource created successfully"
	Updated = "Resource updated successfully"
	Deleted = "Resource deleted successfully"
)

func CreateSuccessMessage(messages ...string) *SuccessMessage {
	message := Created
	if len(message) > 0 {
		message += strings.Join(messages, " - ")
	}
	return NewSuccessMessage(http.StatusCreated, message)
}

func NewSuccessMessage(code int, message string) *SuccessMessage {
	return &SuccessMessage{
		Code:    code,
		Message: message,
	}
}

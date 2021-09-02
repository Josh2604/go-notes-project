package apierrors

import "errors"

var (
	ErrUserNotFound       = errors.New("User not found")
	ErrSignUpUser         = errors.New("Invalid username or password")
	ErrInvalidAccessToken = errors.New("Invalid access_token")
)

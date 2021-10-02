package apierrors

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrSignUpUser         = errors.New("invalid username or password")
	ErrInvalidAccessToken = errors.New("invalid access_token")
	ErrParsingParam       = errors.New("error parsing param")
)

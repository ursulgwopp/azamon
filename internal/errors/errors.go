package errors

import "errors"

var (
	ErrInvalidUsername = errors.New("invalid username")
	ErrUsernameExists  = errors.New("username exists")

	ErrInvalidEmail = errors.New("invalid email")
	ErrEmailExists  = errors.New("email exists")

	ErrInvalidPassword = errors.New("invalid password")

	ErrInvalidUsernameOrPassword = errors.New("invalid username or password")
)

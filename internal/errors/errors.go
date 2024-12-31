package errors

import "errors"

var (
	//auth
	ErrInvalidUsername           = errors.New("invalid username")
	ErrUsernameExists            = errors.New("username exists")
	ErrInvalidEmail              = errors.New("invalid email")
	ErrEmailExists               = errors.New("email exists")
	ErrInvalidPassword           = errors.New("invalid password")
	ErrInvalidUsernameOrPassword = errors.New("invalid username or password")

	//jwt
	ErrInvalidToken         = errors.New("invalid token")
	ErrInvalidSigningMethod = errors.New("invalid signing method")
	ErrInvalidTokenClaims   = errors.New("invalid token claims")

	//items
	ErrItemIdNotFound     = errors.New("item id not found")
	ErrInvalidName        = errors.New("invalid name")
	ErrInvalidDescription = errors.New("invalid description")
	ErrInvalidQuantity    = errors.New("invalid quantity")
	ErrInvalidPrice       = errors.New("invalid price")
	ErrAccessToItemDenied = errors.New("access to item denied")
)

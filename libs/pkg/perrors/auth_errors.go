package perrors

import "errors"

var (
	ErrAlreadyLoggedOut = errors.New("user already logged out")
)

package pkgerr

import "errors"

var ErrAlreadyLoggedOut = errors.New("user already logged out")

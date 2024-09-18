package svcerr

import "errors"

var (
	ErrMaxOtpAttempt = errors.New("max otp attempts reached")
)

package otp

import (
	"math/rand/v2"
	"time"
)

const (
	otpNums = "0123456789"
	expTime = 5 * time.Minute
)

type Otp struct {
	Value     string
	ExpiredAt time.Time
}

// Generate new Otp.
func Generate(length int) Otp {
	return Otp{
		Value:     newValue(length),
		ExpiredAt: newExpiredAt(),
	}
}

// newValue generates new string OTP.
func newValue(length int) string {
	var otp string
	for i := 0; i < length; i++ {
		otp += string(otpNums[rand.IntN(length)])
	}
	return otp
}

// newExpiredAt create new expired time for OTP.
// The expired time is 5 minute.
func newExpiredAt() time.Time {
	now := time.Now()
	return now.Add(expTime)
}

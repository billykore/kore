package otp

import (
	"time"

	"gorm.io/gorm"
)

// OTP is OTP entity.
type OTP struct {
	gorm.Model
	Email     string
	Otp       string
	IsActive  bool
	ExpiresAt time.Time
}

func (OTP) TableName() string {
	return "otp"
}

// IsExpired determine if ExpiresAt is before the current time.
func (otp OTP) IsExpired() bool {
	return otp.ExpiresAt.Before(time.Now())
}

// EmailData defines data for email to send.
type EmailData struct {
	Recipient string
	Subject   string
	OTP       string
}

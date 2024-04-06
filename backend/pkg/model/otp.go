package model

import (
	"time"

	"gorm.io/gorm"
)

type Otp struct {
	gorm.Model
	Email     string
	Otp       string
	IsActive  bool
	ExpiresAt time.Time
}

func (Otp) TableName() string {
	return "otp"
}

// IsExpired determine if ExpiresAt is before the current time.
func (otp Otp) IsExpired() bool {
	return otp.ExpiresAt.Before(time.Now())
}

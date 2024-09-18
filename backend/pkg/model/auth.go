package model

import (
	"time"

	"gorm.io/gorm"
)

type AuthActivities struct {
	gorm.Model
	UUID             string
	Username         string
	LoginTime        time.Time
	LogoutTime       time.Time
	IsLoggedOut      bool
	IsLoginSucceed   bool
	LastLoginAttempt time.Time
}

// IsLoginLastAttemptExpired tells if last failed login activity is pass 24 hours.
func (a *AuthActivities) IsLoginLastAttemptExpired() bool {
	return time.Now().After(a.LastLoginAttempt.Add(24 * time.Hour))
}

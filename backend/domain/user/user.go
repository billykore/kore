package user

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// User entity.
type User struct {
	gorm.Model
	Username string
	Password string
}

// AuthActivities entity.
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

func (a *AuthActivities) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}

func (a *AuthActivities) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, a)
}

package model

import (
	"time"

	"gorm.io/gorm"
)

type AuthActivities struct {
	gorm.Model
	UUID        string
	Username    string
	LoginTime   time.Time
	LogoutTime  time.Time
	IsLoggedOut bool
}

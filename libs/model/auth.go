package model

import "time"

type Auth struct {
	Id          string
	Username    string
	Token       string
	LoginTime   time.Time
	LogoutTime  time.Time
	IsLoggedOut bool
}

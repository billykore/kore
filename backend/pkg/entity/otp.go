package entity

import "time"

type SendOtpRequest struct {
	Email string `json:"email"`
}

type VerifyOtpRequest struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}

type OtpResponse struct {
	Otp       string    `json:"otp"`
	ExpiresAt time.Time `json:"expiresAt"`
}

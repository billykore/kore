package entity

type SendOtpRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type VerifyOtpRequest struct {
	Email string `json:"email" validate:"required,email"`
	Otp   string `json:"otp" validate:"required,number,len=6"`
}

type OtpResponse struct {
	Otp       string `json:"otp"`
	ExpiresAt string `json:"expiresAt"`
}

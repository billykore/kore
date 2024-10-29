package otp

type SendOtpRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type VerifyOtpRequest struct {
	Email string `json:"email" validate:"required,email"`
	Otp   string `json:"otp" validate:"required,number,len=6"`
}

type Response struct {
	Otp       string `json:"otp"`
	ExpiresAt string `json:"expiresAt"`
}

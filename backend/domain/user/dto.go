package user

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=8"`
}

type RegisterResponse struct {
	Username string `json:"username"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token *Token `json:"token"`
}

type Token struct {
	LoginId     string `json:"loginId"`
	AccessToken string `json:"accessToken"`
	ExpiredTime int64  `json:"expiredTime"`
}

type LogoutRequest struct {
	LoginId string `json:"loginId"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}

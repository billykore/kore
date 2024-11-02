package user

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
	LoginId     string `json:"loginId"`
	AccessToken string `json:"accessToken"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}

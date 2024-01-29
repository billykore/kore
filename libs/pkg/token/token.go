package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const tokenExpiredTime = 15 * time.Minute

type Token struct {
	AccessToken string
	ExpiredTime int64
}

func New(username string) (*Token, error) {
	return generateToken(username)
}

func generateToken(username string) (*Token, error) {
	exp := time.Now().Add(tokenExpiredTime).Unix()

	claims := make(jwt.MapClaims)
	claims["username"] = username
	claims["authorize"] = true
	claims["exp"] = exp

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["kid"] = "drPB_Wlr8gCYSaNp4GxJi6w61b8N1oosZQ8sxD9R1Is"

	t, err := token.SignedString([]byte("token-secret"))
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken: t,
		ExpiredTime: exp,
	}, nil
}

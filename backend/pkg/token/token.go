package token

import (
	"errors"
	"time"

	"github.com/billykore/kore/backend/pkg/config"
	"github.com/golang-jwt/jwt"
)

const tokenExpiredTime = 15 * time.Minute

type Token struct {
	AccessToken string
	ExpiredTime int64
}

// New return new generated token.
func New(username string) (*Token, error) {
	return generateToken(username)
}

func generateToken(username string) (*Token, error) {
	exp := time.Now().Add(tokenExpiredTime).Unix()

	claims := make(jwt.MapClaims)
	claims["username"] = username
	claims["authorize"] = true
	claims["exp"] = exp

	cfg := config.Get()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["kid"] = cfg.Token.HeaderKid

	t, err := token.SignedString([]byte(cfg.Token.Secret))
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken: t,
		ExpiredTime: exp,
	}, nil
}

// Verify if token is valid or not.
func Verify(token string) (string, error) {
	return verifyToken(token)
}

func verifyToken(token string) (string, error) {
	cfg := config.Get()
	t, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return []byte(cfg.Token.Secret), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := t.Claims.(jwt.MapClaims); ok || t.Valid {
		return claims["username"].(string), nil
	}
	return "", errors.New("invalid token")
}

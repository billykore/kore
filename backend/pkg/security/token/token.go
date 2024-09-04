package token

import (
	"errors"
	"time"

	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/entity"
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

// Verify if token is valid or not and return extracted user data from token.
func Verify(token string) (entity.User, error) {
	return verifyToken(token)
}

func verifyToken(token string) (entity.User, error) {
	cfg := config.Get()
	t, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return []byte(cfg.Token.Secret), nil
	})
	if err != nil {
		return entity.User{}, err
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok && !t.Valid {
		return entity.User{}, errors.New("invalid token")
	}
	return entity.User{
		Username: claims["username"].(string),
	}, nil
}

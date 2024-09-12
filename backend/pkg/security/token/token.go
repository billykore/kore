package token

import (
	"errors"
	"time"

	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/uuid"
	"github.com/golang-jwt/jwt"
)

const tokenExpiredTime = 15 * time.Minute

type Token struct {
	AccessToken string `json:"accessToken"`
	ExpiredTime int64  `json:"expiredTime"`
}

// New return new generated token.
func New(username string) (Token, error) {
	return generateToken(username)
}

func generateToken(username string) (Token, error) {
	id, err := uuid.New()
	if err != nil {
		return Token{}, err
	}
	now := time.Now()
	exp := now.Add(tokenExpiredTime)

	claims := jwt.StandardClaims{
		Id:        id,
		Issuer:    "https://gateway.kore.co.id",
		Subject:   username,
		IssuedAt:  now.Unix(),
		ExpiresAt: exp.Unix(),
		NotBefore: exp.Unix(),
	}

	cfg := config.Get()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["kid"] = cfg.Token.HeaderKid

	t, err := token.SignedString([]byte(cfg.Token.Secret))
	if err != nil {
		return Token{}, err
	}

	return Token{
		AccessToken: t,
		ExpiredTime: exp.Unix(),
	}, nil
}

// Verify if token is valid or not and return extracted user data from token.
func Verify(token string) (entity.User, error) {
	return verifyToken(token)
}

func verifyToken(token string) (entity.User, error) {
	cfg := config.Get()
	claims := jwt.StandardClaims{}
	t, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (any, error) {
		return []byte(cfg.Token.Secret), nil
	})
	if err != nil {
		return entity.User{}, err
	}
	if !t.Valid {
		return entity.User{}, errors.New("invalid token")
	}
	return entity.User{
		LoginId:  claims.Id,
		Username: claims.Subject,
	}, nil
}

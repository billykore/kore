package token

import (
	"time"

	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/uuid"
	"github.com/golang-jwt/jwt/v5"
)

const tokenExpiredTime = 15 * time.Minute

type Token struct {
	AccessToken string `json:"accessToken"`
	ExpiredTime int64  `json:"expiredTime"`
}

// New return new generated token for the given username.
func New(username string) (Token, error) {
	cfg := config.Get()
	return generateToken(cfg, username)
}

func generateToken(cfg *config.Config, username string) (Token, error) {
	id, err := uuid.New()
	if err != nil {
		return Token{}, err
	}
	now := time.Now()
	exp := now.Add(tokenExpiredTime)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"jti": id,
		"sub": username,
		"iss": "todo-app",
		"aud": username,
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})
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

// UserFromToken return user information from JWT token.
func UserFromToken(token *jwt.Token) entity.User {
	claims := token.Claims.(jwt.MapClaims)
	return entity.User{
		LoginId:  claims["jti"].(string),
		Username: claims["sub"].(string),
	}
}

package usecase

import (
	"context"
	"errors"

	"github.com/billykore/todolist/internal/pkg/token"
	v1 "github.com/billykore/todolist/internal/proto/v1"
)

type AuthUsecase struct{}

func NewAuthUsecase() *AuthUsecase {
	return &AuthUsecase{}
}

func (uc *AuthUsecase) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	if req.GetUsername() != "kore" {
		return nil, errors.New("invalid username")
	}
	if req.GetPassword() != "passwd" {
		return nil, errors.New("invalid password")
	}

	t, err := token.New(req.GetUsername())
	if err != nil {
		return nil, err
	}
	return &v1.LoginReply{
		AccessToken: t.AccessToken,
		ExpiredTime: t.ExpiredTime,
	}, nil
}

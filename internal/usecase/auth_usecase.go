package usecase

import (
	"context"
	"errors"

	"github.com/billykore/todolist/internal/pkg/log"
	"github.com/billykore/todolist/internal/pkg/token"
	v1 "github.com/billykore/todolist/internal/proto/v1"
)

type AuthUsecase struct {
	log *log.Logger
}

func NewAuthUsecase(log *log.Logger) *AuthUsecase {
	return &AuthUsecase{log: log}
}

func (uc *AuthUsecase) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	if req.GetUsername() != "kore" {
		uc.log.Usecase("Login").Errorf("invalid username %s", req.GetUsername())
		return nil, errors.New("invalid username")
	}
	if req.GetPassword() != "passwd" {
		uc.log.Usecase("Login").Errorf("invalid password %s", req.GetUsername())
		return nil, errors.New("invalid password")
	}

	t, err := token.New(req.GetUsername())
	if err != nil {
		uc.log.Usecase("Login").Errorf("failed to create token: %v", err)
		return nil, err
	}

	return &v1.LoginReply{
		AccessToken: t.AccessToken,
		ExpiredTime: t.ExpiredTime,
	}, nil
}

package usecase

import (
	"context"

	"github.com/billykore/todolist/internal/errors"
	"github.com/billykore/todolist/internal/pkg/log"
	"github.com/billykore/todolist/internal/pkg/password"
	"github.com/billykore/todolist/internal/pkg/token"
	v1 "github.com/billykore/todolist/internal/proto/v1"
	"github.com/billykore/todolist/internal/repository"
)

type AuthUsecase struct {
	log  *log.Logger
	repo *repository.UserRepository
}

func NewAuthUsecase(log *log.Logger, repo *repository.UserRepository) *AuthUsecase {
	return &AuthUsecase{
		log:  log,
		repo: repo,
	}
}

func (uc *AuthUsecase) Login(ctx context.Context, req *v1.LoginRequest) (*v1.Token, error) {
	user, err := uc.repo.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		uc.log.Usecase("Login").Errorf("failed to get user by username %s: %v", req.GetUsername(), err)
		return nil, errors.Error{
			Type:    errors.TypeUnauthorized,
			Message: "Sorry, your username or password was incorrect",
		}
	}
	err = password.Verify(user.Password, req.GetPassword())
	if err != nil {
		uc.log.Usecase("Login").Errorf("failed to verify user %s password: %v", user.Username, err)
		return nil, errors.Error{
			Type:    errors.TypeUnauthorized,
			Message: "Sorry, your username or password was incorrect",
		}
	}
	t, err := token.New(user.Username)
	if err != nil {
		uc.log.Usecase("Login").Errorf("failed to create token: %v", err)
		return nil, errors.Error{
			Type:    errors.TypeUnauthorized,
			Message: "Sorry, your username or password was incorrect",
		}
	}
	return &v1.Token{
		AccessToken: t.AccessToken,
		ExpiredTime: t.ExpiredTime,
	}, nil
}

package usecase

import (
	"context"

	"github.com/billykore/kore/libs/pkg/log"
	"github.com/billykore/kore/libs/pkg/messages"
	"github.com/billykore/kore/libs/pkg/password"
	"github.com/billykore/kore/libs/pkg/token"
	"github.com/billykore/kore/libs/proto/v1"
	"github.com/billykore/kore/libs/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthUsecase struct {
	log  *log.Logger
	repo repository.User
}

func NewAuthUsecase(log *log.Logger, repo repository.User) *AuthUsecase {
	return &AuthUsecase{
		log:  log,
		repo: repo,
	}
}

func (uc *AuthUsecase) Login(ctx context.Context, req *v1.LoginRequest) (*v1.Token, error) {
	user, err := uc.repo.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		uc.log.Usecase("Login").Errorf("failed to get user by username %s: %v", req.GetUsername(), err)
		return nil, status.Error(codes.Unauthenticated, messages.InvalidUsernameOrPassword)
	}
	err = password.Verify(user.Password, req.GetPassword())
	if err != nil {
		uc.log.Usecase("Login").Errorf("failed to verify user %s password: %v", user.Username, err)
		return nil, status.Error(codes.Unauthenticated, messages.InvalidUsernameOrPassword)
	}
	t, err := token.New(user.Username)
	if err != nil {
		uc.log.Usecase("Login").Errorf("failed to create token: %v", err)
		return nil, status.Error(codes.Unauthenticated, messages.InvalidUsernameOrPassword)
	}
	return &v1.Token{
		AccessToken: t.AccessToken,
		ExpiredTime: t.ExpiredTime,
	}, nil
}

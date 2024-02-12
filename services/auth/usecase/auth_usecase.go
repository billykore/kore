package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/billykore/kore/libs/model"
	"github.com/billykore/kore/libs/pkg/log"
	"github.com/billykore/kore/libs/pkg/messages"
	"github.com/billykore/kore/libs/pkg/password"
	"github.com/billykore/kore/libs/pkg/perrors"
	"github.com/billykore/kore/libs/pkg/token"
	"github.com/billykore/kore/libs/proto/v1"
	"github.com/billykore/kore/libs/repo"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthUsecase struct {
	log      *log.Logger
	userRepo repo.UserRepository
	authRepo repo.AuthRepository
}

func NewAuthUsecase(log *log.Logger, userRepo repo.UserRepository, authRepo repo.AuthRepository) *AuthUsecase {
	return &AuthUsecase{
		log:      log,
		userRepo: userRepo,
		authRepo: authRepo,
	}
}

func (uc *AuthUsecase) Login(ctx context.Context, req *v1.LoginRequest) (*v1.Token, error) {
	user, err := uc.userRepo.GetByUsername(ctx, req.GetUsername())
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

	id, err := uuid.NewUUID()
	if err != nil {
		uc.log.Usecase("Login").Error(err)
		return nil, status.Error(codes.Internal, messages.LoginFailed)
	}

	err = uc.authRepo.Login(ctx, &model.Auth{
		Id:        id.String(),
		Username:  req.GetUsername(),
		Token:     t.AccessToken,
		LoginTime: time.Now(),
	})
	if err != nil {
		uc.log.Usecase("Login").Errorf("failed to save login activity: %v", err)
		return nil, status.Error(codes.Internal, messages.LoginFailed)
	}

	return &v1.Token{
		LoginId:     id.String(),
		AccessToken: t.AccessToken,
		ExpiredTime: t.ExpiredTime,
	}, nil
}

func (uc *AuthUsecase) Logout(ctx context.Context, req *v1.LogoutRequest) error {
	username, err := token.Verify(req.GetAccessToken())
	if err != nil {
		uc.log.Usecase("Logout").Errorf("failed to verify token: %v", err)
		return status.Error(codes.Unauthenticated, messages.LogoutFailed)
	}

	err = uc.authRepo.Logout(ctx, &model.Auth{
		Id:       req.GetLoginId(),
		Username: username,
	})
	if err != nil && errors.Is(err, perrors.ErrAlreadyLoggedOut) {
		uc.log.Usecase("Logout").Errorf("failed to save logout activity: %v", err)
		return status.Error(codes.Unauthenticated, messages.UserAlreadyLoggedOut)
	}
	if err != nil {
		uc.log.Usecase("Logout").Errorf("failed to save logout activity: %v", err)
		return status.Error(codes.Unauthenticated, messages.LogoutFailed)
	}

	return nil
}

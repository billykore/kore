package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/billykore/kore/backend/pkg/codes"
	"github.com/billykore/kore/backend/pkg/ctxt"
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/pkg/messages"
	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/repo"
	"github.com/billykore/kore/backend/pkg/security/password"
	"github.com/billykore/kore/backend/pkg/security/token"
	"github.com/billykore/kore/backend/pkg/status"
	"github.com/billykore/kore/backend/pkg/svcerr"
	"github.com/billykore/kore/backend/pkg/uuid"
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

func (uc *AuthUsecase) Login(ctx context.Context, req entity.LoginRequest) (*entity.Token, error) {
	user, err := uc.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		uc.log.Usecase("Login").Errorf("failed to get user by username (%s): %v", req.Username, err)
		return nil, status.Error(codes.BadRequest, messages.InvalidUsernameOrPassword)
	}

	err = password.Verify(user.Password, req.Password)
	if err != nil {
		uc.log.Usecase("Login").Errorf("failed to verify user (%s) password: %v", user.Username, err)
		return nil, status.Error(codes.BadRequest, messages.InvalidUsernameOrPassword)
	}

	t, err := token.New(user.Username)
	if err != nil {
		uc.log.Usecase("Login").Errorf("failed to create token: %v", err)
		return nil, status.Error(codes.BadRequest, messages.InvalidUsernameOrPassword)
	}

	id, err := uuid.New()
	if err != nil {
		uc.log.Usecase("Login").Error(err)
		return nil, status.Error(codes.Internal, messages.LoginFailed)
	}

	err = uc.authRepo.Login(ctx, &model.AuthActivities{
		UUID:      id,
		Username:  req.Username,
		LoginTime: time.Now(),
	})
	if err != nil {
		uc.log.Usecase("Login").Errorf("failed to save login activity: %v", err)
		return nil, status.Error(codes.Internal, messages.LoginFailed)
	}

	return &entity.Token{
		LoginId:     id,
		AccessToken: t.AccessToken,
		ExpiredTime: t.ExpiredTime,
	}, nil
}

func (uc *AuthUsecase) Logout(ctx context.Context, req entity.LogoutRequest) (*entity.LogoutResponse, error) {
	user, ok := ctxt.UserFromContext(ctx)
	if !ok {
		uc.log.Usecase("Logout").Error(ctxt.ErrGetUserFromContext)
		return nil, status.Error(codes.Unauthenticated, messages.LogoutFailed)
	}

	err := uc.authRepo.Logout(ctx, &model.AuthActivities{
		UUID:     req.LoginId,
		Username: user.Username,
	})
	if err != nil && errors.Is(err, svcerr.ErrAlreadyLoggedOut) {
		uc.log.Usecase("Logout").Errorf("failed to save logout activity: %v", err)
		return nil, status.Error(codes.Unauthenticated, messages.UserAlreadyLoggedOut)
	}
	if err != nil {
		uc.log.Usecase("Logout").Errorf("failed to save logout activity: %v", err)
		return nil, status.Error(codes.Unauthenticated, messages.LogoutFailed)
	}

	return &entity.LogoutResponse{Message: messages.LogoutSucceed}, nil
}

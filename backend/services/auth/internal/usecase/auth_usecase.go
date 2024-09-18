package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/billykore/kore/backend/pkg/codes"
	"github.com/billykore/kore/backend/pkg/ctxt"
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/pkg/messages"
	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/security/password"
	"github.com/billykore/kore/backend/pkg/security/token"
	"github.com/billykore/kore/backend/pkg/status"
	"github.com/billykore/kore/backend/pkg/svcerr"
	"github.com/billykore/kore/backend/pkg/uuid"
	"github.com/billykore/kore/backend/services/auth/internal/repo"
)

type AuthUsecase struct {
	log      *log.Logger
	userRepo *repo.UserRepository
	authRepo *repo.AuthRepository
}

func NewAuthUsecase(log *log.Logger, userRepo *repo.UserRepository, authRepo *repo.AuthRepository) *AuthUsecase {
	return &AuthUsecase{
		log:      log,
		userRepo: userRepo,
		authRepo: authRepo,
	}
}

// stores username login attempts.
var userLoginAttempt = make(map[string]int)

// max user login attempts.
const maxUserLoginAttempts = 5

func (uc *AuthUsecase) Login(ctx context.Context, req entity.LoginRequest) (*entity.Token, error) {
	if userLoginAttempt[req.Username] == 0 {
		err := uc.checkFailedLoginActivity(ctx, req.Username)
		if err != nil {
			uc.log.Usecase("Login").Errorf("failed to check failed login activity by username (%s): %v", req.Username, err)
			return nil, status.Error(codes.Internal, messages.FailedLoginAttemptNotExpired)
		}
	}

	id, err := uuid.New()
	if err != nil {
		uc.log.Usecase("Login").Error(err)
		return nil, status.Error(codes.Internal, messages.LoginFailed)
	}

	if userLoginAttempt[req.Username] >= maxUserLoginAttempts {
		err = uc.authRepo.SaveLoginActivity(ctx, model.AuthActivities{
			UUID:             id,
			Username:         req.Username,
			IsLoggedOut:      false,
			IsLoginSucceed:   false,
			LastLoginAttempt: time.Now(),
		})
		if err != nil {
			uc.log.Usecase("Login").Errorf("failed to save login activity: %v", err)
			return nil, status.Error(codes.Internal, messages.LoginFailed)
		}

		// max attempts is reached, reset attempts and return error.
		userLoginAttempt[req.Username] = 0

		uc.log.Usecase("Login").Error(fmt.Errorf("user (%s) already login 5 times", req.Username))
		return nil, status.Error(codes.BadRequest, messages.MaxLoginAttemptReached)
	}

	user, err := uc.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		// increment user login attempts.
		userLoginAttempt[req.Username]++

		uc.log.Usecase("Login").Errorf("failed to get user by username (%s): %v", req.Username, err)
		return nil, status.Error(codes.BadRequest, messages.InvalidUsernameOrPassword)
	}

	err = password.Verify(user.Password, req.Password)
	if err != nil {
		// increment user login attempts.
		userLoginAttempt[req.Username]++

		uc.log.Usecase("Login").Errorf("failed to verify user (%s) password: %v", user.Username, err)
		return nil, status.Error(codes.BadRequest, messages.InvalidUsernameOrPassword)
	}

	t, err := token.New(user.Username)
	if err != nil {
		uc.log.Usecase("Login").Errorf("failed to create token: %v", err)
		return nil, status.Error(codes.BadRequest, messages.InvalidUsernameOrPassword)
	}

	err = uc.authRepo.Login(ctx, model.AuthActivities{
		UUID:           id,
		Username:       req.Username,
		LoginTime:      time.Now(),
		IsLoginSucceed: true,
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

// checkFailedLoginActivity checks any user's failed login activities.
// If there is any, then checks if the failed login time is expired,
// so the user can try new login.
func (uc *AuthUsecase) checkFailedLoginActivity(ctx context.Context, username string) error {
	activity, err := uc.authRepo.FindFailedLoginByUsername(ctx, username)
	if err != nil {
		return fmt.Errorf("failed to get login activity by username (%s): %v", username, err)
	}
	if activity == nil {
		return nil
	}
	if !activity.IsLoginSucceed && !activity.IsLoginLastAttemptExpired() {
		return errors.New("last failed login attempt is not expired yet")
	}
	return nil
}

func (uc *AuthUsecase) Logout(ctx context.Context, req entity.LogoutRequest) (*entity.LogoutResponse, error) {
	user, ok := ctxt.UserFromContext(ctx)
	if !ok {
		uc.log.Usecase("Logout").Error(ctxt.ErrGetUserFromContext)
		return nil, status.Error(codes.Unauthenticated, messages.LogoutFailed)
	}

	err := uc.authRepo.Logout(ctx, model.AuthActivities{
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

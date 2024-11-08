package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/billykore/kore/backend/pkg/codes"
	"github.com/billykore/kore/backend/pkg/ctxt"
	"github.com/billykore/kore/backend/pkg/logger"
	"github.com/billykore/kore/backend/pkg/messages"
	"github.com/billykore/kore/backend/pkg/pkgerr"
	"github.com/billykore/kore/backend/pkg/security/password"
	"github.com/billykore/kore/backend/pkg/security/token"
	"github.com/billykore/kore/backend/pkg/status"
	"github.com/billykore/kore/backend/pkg/uuid"
)

type Repository interface {
	GetByUsername(ctx context.Context, username string) (*User, error)
	Login(ctx context.Context, auth AuthActivities) error
	SaveLoginActivity(ctx context.Context, auth AuthActivities) error
	Logout(ctx context.Context, auth AuthActivities) error
	FindFailedLoginByUsername(ctx context.Context, username string) (*AuthActivities, error)
}

type Service struct {
	logger *logger.Logger
	repo   Repository
}

func NewService(logger *logger.Logger, repo Repository) *Service {
	return &Service{
		logger: logger,
		repo:   repo,
	}
}

// stores username login attempts.
var userLoginAttempt = make(map[string]int)

// max user login attempts.
const maxUserLoginAttempts = 5

func (s *Service) Login(ctx context.Context, req LoginRequest) (*Token, error) {
	if userLoginAttempt[req.Username] == 0 {
		err := s.checkFailedLoginActivity(ctx, req.Username)
		if err != nil {
			s.logger.Usecase("Login").Errorf("failed to check failed loggerin activity by username (%s): %v", req.Username, err)
			return nil, status.Error(codes.Internal, messages.FailedLoginAttemptNotExpired)
		}
	}

	id, err := uuid.New()
	if err != nil {
		s.logger.Usecase("Login").Error(err)
		return nil, status.Error(codes.Internal, messages.LoginFailed)
	}

	if userLoginAttempt[req.Username] >= maxUserLoginAttempts {
		err = s.repo.SaveLoginActivity(ctx, AuthActivities{
			UUID:             id,
			Username:         req.Username,
			IsLoggedOut:      false,
			IsLoginSucceed:   false,
			LastLoginAttempt: time.Now(),
		})
		if err != nil {
			s.logger.Usecase("Login").Errorf("failed to save loggerin activity: %v", err)
			return nil, status.Error(codes.Internal, messages.LoginFailed)
		}

		// max attempts is reached, reset attempts and return error.
		userLoginAttempt[req.Username] = 0

		s.logger.Usecase("Login").Error(fmt.Errorf("user (%s) already loggerin %d times",
			req.Username, maxUserLoginAttempts))
		return nil, status.Error(codes.BadRequest, messages.MaxLoginAttemptReached)
	}

	userLogin, err := s.repo.GetByUsername(ctx, req.Username)
	if err != nil {
		// increment user login attempts.
		userLoginAttempt[req.Username]++

		s.logger.Usecase("Login").Errorf("failed to get user by username (%s): %v", req.Username, err)
		return nil, status.Error(codes.BadRequest, messages.InvalidUsernameOrPassword)
	}

	err = password.Verify(userLogin.Password, req.Password)
	if err != nil {
		// increment user login attempts.
		userLoginAttempt[req.Username]++

		s.logger.Usecase("Login").Errorf("failed to verify user (%s) password: %v", userLogin.Username, err)
		return nil, status.Error(codes.BadRequest, messages.InvalidUsernameOrPassword)
	}

	t, err := token.New(userLogin.Username)
	if err != nil {
		s.logger.Usecase("Login").Errorf("failed to create token: %v", err)
		return nil, status.Error(codes.BadRequest, messages.InvalidUsernameOrPassword)
	}

	err = s.repo.Login(ctx, AuthActivities{
		UUID:           id,
		Username:       req.Username,
		LoginTime:      time.Now(),
		IsLoginSucceed: true,
	})
	if err != nil {
		s.logger.Usecase("Login").Errorf("failed to save loggerin activity: %v", err)
		return nil, status.Error(codes.Internal, messages.LoginFailed)
	}

	return &Token{
		LoginId:     id,
		AccessToken: t.AccessToken,
		ExpiredTime: t.ExpiredTime,
	}, nil
}

// checkFailedLoginActivity checks any user's failed login activities.
// If there is any, then checks if the failed login time is expired,
// so the user can try new login.
func (s *Service) checkFailedLoginActivity(ctx context.Context, username string) error {
	activity, err := s.repo.FindFailedLoginByUsername(ctx, username)
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

func (s *Service) Logout(ctx context.Context, req LogoutRequest) (*LogoutResponse, error) {
	userLogin, ok := ctxt.UserFromContext(ctx)
	if !ok {
		s.logger.Usecase("Logout").Error(ctxt.ErrGetUserFromContext)
		return nil, status.Error(codes.Unauthenticated, messages.LogoutFailed)
	}

	err := s.repo.Logout(ctx, AuthActivities{
		UUID:     req.LoginId,
		Username: userLogin.Username,
	})
	if err != nil && errors.Is(err, pkgerr.ErrAlreadyLoggedOut) {
		s.logger.Usecase("Logout").Errorf("failed to save loggerout activity: %v", err)
		return nil, status.Error(codes.Unauthenticated, messages.UserAlreadyLoggedOut)
	}
	if err != nil {
		s.logger.Usecase("Logout").Errorf("failed to save loggerout activity: %v", err)
		return nil, status.Error(codes.Unauthenticated, messages.LogoutFailed)
	}

	return &LogoutResponse{Message: messages.LogoutSucceed}, nil
}

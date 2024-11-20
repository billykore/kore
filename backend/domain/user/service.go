package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/billykore/kore/backend/pkg/codes"
	"github.com/billykore/kore/backend/pkg/ctxt"
	"github.com/billykore/kore/backend/pkg/logger"
	"github.com/billykore/kore/backend/pkg/security/password"
	"github.com/billykore/kore/backend/pkg/security/token"
	"github.com/billykore/kore/backend/pkg/status"
	"github.com/billykore/kore/backend/pkg/uuid"
)

// maxLoginAttempts is max user login attempts.
const maxLoginAttempts = 5

// userLoginAttempt stores username login attempts.
var userLoginAttempt = make(map[string]int)

// ErrAlreadyExists indicates an attempt to create a user failed
// because the user already exists in the system.
var ErrAlreadyExists = errors.New("user already exists")

// ErrAlreadyLoggedOut indicates an operation was attempted on a user who has already logged out.
var ErrAlreadyLoggedOut = errors.New("user already logged out")

var errFailedCacheLoginActivity = errors.New("failed cache login activity")

// Repository defines the methods to interacting with persistence storage used by user domain.
type Repository interface {
	// Create creates new user.
	Create(ctx context.Context, user User) error

	// GetByUsername gets specific user by username.
	GetByUsername(ctx context.Context, username string) (*User, error)

	// Login saves user login.
	Login(ctx context.Context, auth AuthActivities) error

	// SaveLoginActivity saves user login activities.
	SaveLoginActivity(ctx context.Context, auth AuthActivities) error

	// Logout set user login activity to logged out.
	Logout(ctx context.Context, auth AuthActivities) error
}

// Cache defines the methods to interacting with cache storage used by user domain.
type Cache interface {
	// SaveFailedLoginAttempt saves failed login attempt of user by the given username.
	SaveFailedLoginAttempt(ctx context.Context, authActivity AuthActivities) error

	// GetFailedLoginAttempt gets failed login attempt by username.
	GetFailedLoginAttempt(ctx context.Context, username string) (*AuthActivities, error)
}

type Service struct {
	log   *logger.Logger
	repo  Repository
	cache Cache
}

func NewService(log *logger.Logger, repo Repository, cache Cache) *Service {
	return &Service{
		log:   log,
		repo:  repo,
		cache: cache,
	}
}

func (s *Service) Create(ctx context.Context, req RegisterRequest) (*RegisterResponse, error) {
	hashedPwd, err := password.Hash(req.Password)
	if err != nil {
		s.log.Usecase("Create").Errorf("failed to hash password: %v", err)
		return nil, status.Error(codes.Internal, messageRegisterFailed)
	}
	err = s.repo.Create(ctx, User{
		Username: req.Username,
		Password: hashedPwd,
	})
	if err != nil && errors.Is(err, ErrAlreadyExists) {
		s.log.Usecase("Create").Errorf("failed to create new user: %v", err)
		return nil, status.Errorf(codes.Conflict, "%s: %s", messageRegisterFailed, err.Error())
	}
	if err != nil {
		s.log.Usecase("Create").Errorf("failed to create new user: %v", err)
		return nil, status.Error(codes.Internal, messageRegisterFailed)
	}
	return &RegisterResponse{
		Username: req.Username,
	}, nil
}

func (s *Service) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
	if userLoginAttempt[req.Username] == 0 {
		err := s.checkFailedLoginActivity(ctx, req.Username)
		if err != nil {
			s.log.Usecase("Login").Errorf("failed to check failed login activity by username (%s): %v", req.Username, err)
			return nil, status.Error(codes.Forbidden, messageFailedLoginAttemptNotExpired)
		}
	}

	id, err := uuid.New()
	if err != nil {
		s.log.Usecase("Login").Error(err)
		return nil, status.Error(codes.Internal, messageLoginFailed)
	}

	if userLoginAttempt[req.Username] >= maxLoginAttempts {
		err = s.saveLoginActivity(ctx, AuthActivities{
			UUID:             id,
			Username:         req.Username,
			IsLoggedOut:      false,
			IsLoginSucceed:   false,
			LastLoginAttempt: time.Now(),
		})
		if err != nil {
			s.log.Usecase("Login").Errorf("failed to save login activity: %v", err)
			return nil, status.Error(codes.Internal, messageLoginFailed)
		}

		// max attempts is reached, delete from map and return error.
		delete(userLoginAttempt, req.Username)

		s.log.Usecase("Login").Errorf("user (%s) already login %d times",
			req.Username, maxLoginAttempts)
		return nil, status.Error(codes.BadRequest, messageMaxLoginAttemptReached)
	}

	userLogin, err := s.repo.GetByUsername(ctx, req.Username)
	if err != nil {
		// increment user login attempts.
		userLoginAttempt[req.Username]++

		s.log.Usecase("Login").Errorf("failed to get user by username (%s): %v", req.Username, err)
		return nil, status.Error(codes.BadRequest, messageInvalidUsernameOrPassword)
	}

	err = password.Verify(userLogin.Password, req.Password)
	if err != nil {
		// increment user login attempts.
		userLoginAttempt[req.Username]++

		s.log.Usecase("Login").Errorf("failed to verify user (%s) password: %v", userLogin.Username, err)
		return nil, status.Error(codes.BadRequest, messageInvalidUsernameOrPassword)
	}

	t, err := token.New(userLogin.Username)
	if err != nil {
		s.log.Usecase("Login").Errorf("failed to create token: %v", err)
		return nil, status.Error(codes.BadRequest, messageInvalidUsernameOrPassword)
	}

	err = s.repo.Login(ctx, AuthActivities{
		UUID:           id,
		Username:       req.Username,
		LoginTime:      time.Now(),
		IsLoginSucceed: true,
	})
	if err != nil {
		s.log.Usecase("Login").Errorf("failed to save login activity: %v", err)
		return nil, status.Error(codes.Internal, messageLoginFailed)
	}

	// reset user login attempt if login is successful.
	delete(userLoginAttempt, req.Username)

	return &LoginResponse{
		LoginId:     id,
		AccessToken: t.AccessToken,
		ExpiredTime: t.ExpiredTime,
	}, nil
}

// checkFailedLoginActivity checks any user's failed login activities.
// If there is any, then checks if the failed login time is expired,
// so the user can try new login.
func (s *Service) checkFailedLoginActivity(ctx context.Context, username string) error {
	activity, err := s.cache.GetFailedLoginAttempt(ctx, username)
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

// saveLoginActivity saves user login activity, either it was succeeded or failed.
func (s *Service) saveLoginActivity(ctx context.Context, activity AuthActivities) error {
	err := s.cache.SaveFailedLoginAttempt(ctx, activity)
	if err != nil {
		s.log.Usecase("saveLoginActivity").Error(err)
		return errFailedCacheLoginActivity
	}
	return s.repo.SaveLoginActivity(ctx, activity)
}

func (s *Service) Logout(ctx context.Context, req LogoutRequest) (*LogoutResponse, error) {
	userLogin, ok := ctxt.UserFromContext(ctx)
	if !ok {
		s.log.Usecase("Logout").Error(ctxt.ErrGetUserFromContext)
		return nil, status.Error(codes.Unauthenticated, messageLogoutFailed)
	}
	err := s.repo.Logout(ctx, AuthActivities{
		UUID:     req.LoginId,
		Username: userLogin.Username,
	})
	if err != nil && errors.Is(err, ErrAlreadyLoggedOut) {
		s.log.Usecase("Logout").Errorf("failed to save logout activity: %v", err)
		return nil, status.Errorf(codes.Unauthenticated, "%s: %s", messageLogoutFailed, err.Error())
	}
	if err != nil {
		s.log.Usecase("Logout").Errorf("failed to save logout activity: %v", err)
		return nil, status.Error(codes.Unauthenticated, messageLogoutFailed)
	}
	return &LogoutResponse{Message: messageLogoutSucceed}, nil
}

package otp

import (
	"context"
	"fmt"

	"github.com/billykore/kore/backend/pkg/codes"
	"github.com/billykore/kore/backend/pkg/datetime"
	"github.com/billykore/kore/backend/pkg/logger"
	otpPkg "github.com/billykore/kore/backend/pkg/security/otp"
	"github.com/billykore/kore/backend/pkg/status"
)

// Repository defines the methods to interacting with persistence storage used by OTP domain.
type Repository interface {
	// Get gets the OTP.
	Get(ctx context.Context, otp OTP) (*OTP, error)

	// Save saves new OTP.
	Save(ctx context.Context, otp OTP) error

	// Update updates active OTP.
	Update(ctx context.Context, otp OTP) error
}

// Email is interface for email service used by OTP domain.
type Email interface {
	// SendOTP sends OTP email.
	SendOTP(EmailData) error
}

type Service struct {
	log   *logger.Logger
	repo  Repository
	email Email
}

func NewService(log *logger.Logger, repo Repository, email Email) *Service {
	return &Service{
		log:   log,
		repo:  repo,
		email: email,
	}
}

func (s *Service) SendOtp(ctx context.Context, req SendOtpRequest) (*Response, error) {
	newOtp, err := s.checkAndGenerateOtp(ctx, req.Email)
	if err != nil {
		s.log.Usecase("SendOtp").Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = s.repo.Save(ctx, OTP{
		Email:     req.Email,
		Otp:       newOtp.Value,
		ExpiresAt: newOtp.ExpiredAt,
		IsActive:  true,
	})
	if err != nil {
		s.log.Usecase("SendOtp").Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = s.email.SendOTP(EmailData{
		Recipient: req.Email,
		Subject:   "Login OTP",
		OTP:       newOtp.Value,
	})
	if err != nil {
		s.log.Usecase("SendOtp").Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &Response{
		Otp:       newOtp.Value,
		ExpiresAt: newOtp.ExpiredAt.Format(datetime.DefaultTimeLayout),
	}, nil
}

func (s *Service) checkAndGenerateOtp(ctx context.Context, email string) (otpPkg.Otp, error) {
	newOtp := otpPkg.Generate(6)
	existingOtp, err := s.repo.Get(ctx, OTP{
		Email: email,
		Otp:   newOtp.Value,
	})
	if err != nil {
		return otpPkg.Otp{}, err
	}
	if existingOtp == nil {
		return newOtp, nil
	}
	if existingOtp.IsActive {
		newOtp, err = s.checkAndGenerateOtp(ctx, email)
		if err != nil {
			return otpPkg.Otp{}, err
		}
	}
	return newOtp, nil
}

func (s *Service) VerifyOtp(ctx context.Context, req VerifyOtpRequest) error {
	currentOtp, err := s.repo.Get(ctx, OTP{
		Email: req.Email,
		Otp:   req.Otp,
	})
	if err != nil || currentOtp == nil {
		s.log.Usecase("VerifyOtp").Error(err)
		return status.Error(codes.NotFound, messageInvalidOTP)
	}
	if !currentOtp.IsActive {
		s.log.Usecase("VerifyOtp").Error(
			fmt.Errorf("otp (%s) is not active", currentOtp.Otp))
		return status.Error(codes.BadRequest, messageInvalidOTP)
	}
	if currentOtp.IsExpired() {
		s.log.Usecase("VerifyOtp").Error(
			fmt.Errorf("otp (%s) is expired", currentOtp.Otp))
		return status.Error(codes.BadRequest, messageExpiredOTP)
	}
	err = s.repo.Update(ctx, OTP{Otp: req.Otp})
	if err != nil {
		s.log.Usecase("VerifyOtp").Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

package otp

import (
	"context"
	"fmt"

	"github.com/billykore/kore/backend/internal/domain/otp"
	"github.com/billykore/kore/backend/internal/infra/mail"
	"github.com/billykore/kore/backend/internal/infra/mail/templates"
	"github.com/billykore/kore/backend/pkg/codes"
	"github.com/billykore/kore/backend/pkg/datetime"
	"github.com/billykore/kore/backend/pkg/logger"
	"github.com/billykore/kore/backend/pkg/messages"
	otpPkg "github.com/billykore/kore/backend/pkg/security/otp"
	"github.com/billykore/kore/backend/pkg/status"
)

type Service struct {
	log    *logger.Logger
	repo   otp.Repository
	mailer *mail.Mailer
}

func NewService(log *logger.Logger, repo otp.Repository, mailer *mail.Mailer) *Service {
	return &Service{
		log:    log,
		repo:   repo,
		mailer: mailer,
	}
}

func (s *Service) SendOtp(ctx context.Context, req SendOtpRequest) (*Response, error) {
	newOtp, err := s.checkAndGenerateOtp(ctx, req.Email)
	if err != nil {
		s.log.Usecase("SendOtp").Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = s.repo.Save(ctx, otp.Otp{
		Email:     req.Email,
		Otp:       newOtp.Value,
		ExpiresAt: newOtp.ExpiredAt,
		IsActive:  true,
	})
	if err != nil {
		s.log.Usecase("SendOtp").Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	tmpl, err := templates.OtpTemplate(templates.OtpData{Otp: newOtp.Value})
	if err != nil {
		s.log.Usecase("SendOtp").Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = s.mailer.Send(mail.Data{
		Recipient: req.Email,
		Subject:   "Login OTP",
		Body:      tmpl,
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
	existingOtp, err := s.repo.Get(ctx, otp.Otp{
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
	currentOtp, err := s.repo.Get(ctx, otp.Otp{
		Email: req.Email,
		Otp:   req.Otp,
	})
	if err != nil || currentOtp == nil {
		s.log.Usecase("VerifyOtp").Error(err)
		return status.Error(codes.NotFound, messages.InvalidOTP)
	}
	if !currentOtp.IsActive {
		s.log.Usecase("VerifyOtp").Error(
			fmt.Errorf("otp (%s) is not active", currentOtp.Otp))
		return status.Error(codes.BadRequest, messages.InvalidOTP)
	}
	if currentOtp.IsExpired() {
		s.log.Usecase("VerifyOtp").Error(
			fmt.Errorf("otp (%s) is expired", currentOtp.Otp))
		return status.Error(codes.BadRequest, messages.ExpiredOTP)
	}
	err = s.repo.Update(ctx, otp.Otp{Otp: req.Otp})
	if err != nil {
		s.log.Usecase("VerifyOtp").Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

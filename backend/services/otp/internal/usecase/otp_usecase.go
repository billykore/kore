package usecase

import (
	"context"
	"fmt"

	"github.com/billykore/kore/backend/pkg/codes"
	"github.com/billykore/kore/backend/pkg/datetime"
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/pkg/mail"
	"github.com/billykore/kore/backend/pkg/mail/templates"
	"github.com/billykore/kore/backend/pkg/messages"
	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/security/otp"
	"github.com/billykore/kore/backend/pkg/status"
	"github.com/billykore/kore/backend/services/otp/internal/repo"
)

type OtpUsecase struct {
	log     *log.Logger
	otpRepo *repo.OtpRepository
	mailer  *mail.Mailer
}

func NewOtpUsecase(log *log.Logger, otpRepo *repo.OtpRepository, mailer *mail.Mailer) *OtpUsecase {
	return &OtpUsecase{
		log:     log,
		otpRepo: otpRepo,
		mailer:  mailer,
	}
}

func (uc *OtpUsecase) SendOtp(ctx context.Context, req entity.SendOtpRequest) (*entity.OtpResponse, error) {
	newOtp, err := uc.checkAndGenerateOtp(ctx, req.Email)
	if err != nil {
		uc.log.Usecase("SendOtp").Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = uc.otpRepo.Save(ctx, model.Otp{
		Email:     req.Email,
		Otp:       newOtp.Value,
		ExpiresAt: newOtp.ExpiredAt,
		IsActive:  true,
	})
	if err != nil {
		uc.log.Usecase("SendOtp").Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	tmpl, err := templates.OtpTemplate(templates.OtpData{Otp: newOtp.Value})
	if err != nil {
		uc.log.Usecase("SendOtp").Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = uc.mailer.Send(mail.Data{
		Recipient: req.Email,
		Subject:   "Login OTP",
		Body:      tmpl,
	})
	if err != nil {
		uc.log.Usecase("SendOtp").Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &entity.OtpResponse{
		Otp:       newOtp.Value,
		ExpiresAt: newOtp.ExpiredAt.Format(datetime.DefaultTimeLayout),
	}, nil
}

func (uc *OtpUsecase) checkAndGenerateOtp(ctx context.Context, email string) (otp.Otp, error) {
	newOtp := otp.Generate(6)
	existingOtp, err := uc.otpRepo.Get(ctx, model.Otp{
		Email: email,
		Otp:   newOtp.Value,
	})
	if err != nil {
		return otp.Otp{}, err
	}
	if existingOtp == nil {
		return newOtp, nil
	}
	if existingOtp.IsActive {
		newOtp, err = uc.checkAndGenerateOtp(ctx, email)
		if err != nil {
			return otp.Otp{}, err
		}
	}
	return newOtp, nil
}

func (uc *OtpUsecase) VerifyOtp(ctx context.Context, req entity.VerifyOtpRequest) error {
	currentOtp, err := uc.otpRepo.Get(ctx, model.Otp{
		Email: req.Email,
		Otp:   req.Otp,
	})
	if err != nil || currentOtp == nil {
		uc.log.Usecase("VerifyOtp").Error(err)
		return status.Error(codes.NotFound, messages.InvalidOTP)
	}
	if !currentOtp.IsActive {
		uc.log.Usecase("VerifyOtp").Error(
			fmt.Errorf("otp (%s) is not active", currentOtp.Otp))
		return status.Error(codes.BadRequest, messages.InvalidOTP)
	}
	if currentOtp.IsExpired() {
		uc.log.Usecase("VerifyOtp").Error(
			fmt.Errorf("otp (%s) is expired", currentOtp.Otp))
		return status.Error(codes.BadRequest, messages.ExpiredOTP)
	}
	err = uc.otpRepo.Update(ctx, model.Otp{Otp: req.Otp})
	if err != nil {
		uc.log.Usecase("VerifyOtp").Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

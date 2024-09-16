package handler

import (
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/validation"
	"github.com/billykore/kore/backend/services/otp/internal/usecase"
	"github.com/labstack/echo/v4"
)

type OtpHandler struct {
	va *validation.Validator
	uc *usecase.OtpUsecase
}

func NewOtpHandler(uc *usecase.OtpUsecase, validator *validation.Validator) *OtpHandler {
	return &OtpHandler{
		va: validator,
		uc: uc,
	}
}

// SendOtp swaggo annotation.
//
//	@Summary		Send OTP
//	@Description	Send OTP to customer email
//	@Tags			otp-service
//	@Accept			json
//	@Produce		json
//	@Param			SendOtpRequest	body		entity.SendOtpRequest	true	"Send OTP request"
//	@Success		200				{object}	entity.Response
//	@Failure		400				{object}	entity.Response
//	@Failure		401				{object}	entity.Response
//	@Failure		404				{object}	entity.Response
//	@Failure		500				{object}	entity.Response
//	@Router			/otp/send [post]
func (s *OtpHandler) SendOtp(ctx echo.Context) error {
	var req entity.SendOtpRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	err = s.va.Validate(req)
	if err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	otp, err := s.uc.SendOtp(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(otp))
}

// VerifyOtp swaggo annotation.
//
//	@Summary		Verify OTP
//	@Description	Verify customer OTP
//	@Tags			otp-service
//	@Accept			json
//	@Produce		json
//	@Param			VerifyOtpRequest	body		entity.VerifyOtpRequest	true	"Verify OTP request"
//	@Success		200					{object}	entity.Response
//	@Failure		400					{object}	entity.Response
//	@Failure		401					{object}	entity.Response
//	@Failure		404					{object}	entity.Response
//	@Failure		500					{object}	entity.Response
//	@Router			/otp/verify [post]
func (s *OtpHandler) VerifyOtp(ctx echo.Context) error {
	var req entity.VerifyOtpRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	err = s.va.Validate(req)
	if err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	err = s.uc.VerifyOtp(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccessNilData())
}

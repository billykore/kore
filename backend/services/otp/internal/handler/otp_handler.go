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

package handler

import (
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/services/otp/internal/usecase"
	"github.com/labstack/echo/v4"
)

type OtpHandler struct {
	uc *usecase.OtpUsecase
}

func NewOtpHandler(uc *usecase.OtpUsecase) *OtpHandler {
	return &OtpHandler{uc: uc}
}

func (s *OtpHandler) SendOtp(ctx echo.Context) error {
	var req entity.SendOtpRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	greet, err := s.uc.SendOtp(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(greet))
}

func (s *OtpHandler) VerifyOtp(ctx echo.Context) error {
	var req entity.VerifyOtpRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	err = s.uc.VerifyOtp(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccessNilData())
}

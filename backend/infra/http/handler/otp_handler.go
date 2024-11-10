package handler

import (
	"github.com/billykore/kore/backend/domain/otp"
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/validation"
	"github.com/labstack/echo/v4"
)

type OtpHandler struct {
	va  *validation.Validator
	svc *otp.Service
}

func NewOtpHandler(svc *otp.Service, validator *validation.Validator) *OtpHandler {
	return &OtpHandler{
		va:  validator,
		svc: svc,
	}
}

// SendOtp swaggo annotation.
//
//	@Summary		Send OTP
//	@Description	Send OTP to customer email
//	@Tags			otp
//	@Accept			json
//	@Produce		json
//	@Param			SendOtpRequest	body		otp.SendOtpRequest	true	"Send OTP request"
//	@Success		200				{object}	entity.Response
//	@Failure		400				{object}	entity.Response
//	@Failure		401				{object}	entity.Response
//	@Failure		404				{object}	entity.Response
//	@Failure		500				{object}	entity.Response
//	@Router			/otp/send [post]
func (h *OtpHandler) SendOtp(ctx echo.Context) error {
	var req otp.SendOtpRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	err = h.va.Validate(req)
	if err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	res, err := h.svc.SendOtp(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(res))
}

// VerifyOtp swaggo annotation.
//
//	@Summary		Verify OTP
//	@Description	Verify customer OTP
//	@Tags			otp
//	@Accept			json
//	@Produce		json
//	@Param			VerifyOtpRequest	body		otp.VerifyOtpRequest	true	"Verify OTP request"
//	@Success		200					{object}	entity.Response
//	@Failure		400					{object}	entity.Response
//	@Failure		401					{object}	entity.Response
//	@Failure		404					{object}	entity.Response
//	@Failure		500					{object}	entity.Response
//	@Router			/otp/verify [post]
func (h *OtpHandler) VerifyOtp(ctx echo.Context) error {
	var req otp.VerifyOtpRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	err = h.va.Validate(req)
	if err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	err = h.svc.VerifyOtp(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccessNilData())
}

package handler

import (
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/services/auth/internal/usecase"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	uc *usecase.AuthUsecase
}

func NewAuthHandler(uc *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{uc: uc}
}

func (s *AuthHandler) Login(ctx echo.Context) error {
	var req entity.LoginRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	token, err := s.uc.Login(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(token))
}

func (s *AuthHandler) Logout(ctx echo.Context) error {
	var req entity.LogoutRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	logout, err := s.uc.Logout(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(logout))
}

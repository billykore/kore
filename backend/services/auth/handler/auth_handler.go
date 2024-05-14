package handler

import (
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/services/auth/usecase"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	uc *usecase.AuthUsecase
}

func NewAuthHandler(uc *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{uc: uc}
}

func (s *AuthHandler) Login(ctx echo.Context) error {
	in := new(entity.LoginRequest)
	err := ctx.Bind(in)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	token, err := s.uc.Login(ctx.Request().Context(), in)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("token", token))
}

func (s *AuthHandler) Logout(ctx echo.Context) error {
	in := new(entity.LogoutRequest)
	err := ctx.Bind(in)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	logout, err := s.uc.Logout(ctx.Request().Context(), in)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("logout", logout))
}

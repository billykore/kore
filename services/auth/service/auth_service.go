package service

import (
	"github.com/billykore/kore/pkg/entity"
	"github.com/billykore/kore/services/auth/usecase"
	"github.com/labstack/echo/v4"
)

type AuthService struct {
	uc *usecase.AuthUsecase
}

func NewAuthService(uc *usecase.AuthUsecase) *AuthService {
	return &AuthService{uc: uc}
}

func (s *AuthService) Login(ctx echo.Context) error {
	in := new(entity.LoginRequest)
	err := ctx.Bind(in)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	token, err := s.uc.Login(ctx.Request().Context(), in)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(token))
}

func (s *AuthService) Logout(ctx echo.Context) error {
	in := new(entity.LogoutRequest)
	err := ctx.Bind(in)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	logout, err := s.uc.Logout(ctx.Request().Context(), in)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(logout))
}

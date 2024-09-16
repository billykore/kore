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

// Login swaggo annotation.
//
//	@Summary		User login
//	@Description	User login by username and password
//	@Tags			auth-service
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Authorization header contains Bearer token"
//	@Param			LoginRequest	body		entity.LoginRequest	true	"Login Request"
//	@Success		200				{object}	entity.Response
//	@Failure		400				{object}	entity.Response
//	@Failure		404				{object}	entity.Response
//	@Failure		500				{object}	entity.Response
//	@Router			/login [post]
func (s *AuthHandler) Login(ctx echo.Context) error {
	var req entity.LoginRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	token, err := s.uc.Login(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(token))
}

// Logout swaggo annotation.
//
//	@Summary		User logout
//	@Description	User logout by access token
//	@Tags			auth-service
//	@Accept			json
//	@Produce		json
//	@Param			LogoutRequest	body		entity.LogoutRequest	true	"Logout Request"
//	@Success		200				{object}	entity.Response
//	@Failure		400				{object}	entity.Response
//	@Failure		401				{object}	entity.Response
//	@Failure		404				{object}	entity.Response
//	@Failure		500				{object}	entity.Response
//	@Router			/logout [post]
func (s *AuthHandler) Logout(ctx echo.Context) error {
	var req entity.LogoutRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	logout, err := s.uc.Logout(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(logout))
}

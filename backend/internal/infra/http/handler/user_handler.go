package handler

import (
	"github.com/billykore/kore/backend/internal/domain/user"
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	svc *user.Service
}

func NewUserHandler(svc *user.Service) *UserHandler {
	return &UserHandler{svc: svc}
}

// Login swaggo annotation.
//
//	@Summary		User login
//	@Description	User login by username and password
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Authorization header contains Bearer token"
//	@Param			LoginRequest	body		user.LoginRequest	true	"Login Request"
//	@Success		200				{object}	entity.Response
//	@Failure		400				{object}	entity.Response
//	@Failure		404				{object}	entity.Response
//	@Failure		500				{object}	entity.Response
//	@Router			/login [post]
func (h *UserHandler) Login(ctx echo.Context) error {
	var req user.LoginRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	token, err := h.svc.Login(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(token))
}

// Logout swaggo annotation.
//
//	@Summary		User logout
//	@Description	User logout by access token
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			LogoutRequest	body		user.LogoutRequest	true	"Logout Request"
//	@Success		200				{object}	entity.Response
//	@Failure		400				{object}	entity.Response
//	@Failure		401				{object}	entity.Response
//	@Failure		404				{object}	entity.Response
//	@Failure		500				{object}	entity.Response
//	@Router			/logout [post]
func (h *UserHandler) Logout(ctx echo.Context) error {
	var req user.LogoutRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	logout, err := h.svc.Logout(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(logout))
}

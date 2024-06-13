package handler

import (
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/services/product/internal/usecase"
	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	uc *usecase.CartUsecase
}

func NewCartHandler(uc *usecase.CartUsecase) *CartHandler {
	return &CartHandler{
		uc: uc,
	}
}

func (s *CartHandler) GetCartItemList(ctx echo.Context) error {
	var req entity.CartRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	items, err := s.uc.GetCartItemList(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("carts", items))
}

func (s *CartHandler) AddCartItem(ctx echo.Context) error {
	var req entity.AddCartItemRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	err := s.uc.AddCartItem(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccessNilData())
}

func (s *CartHandler) UpdateCartItemQuantity(ctx echo.Context) error {
	var req entity.UpdateCartItemRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	err := s.uc.UpdateCartItemQuantity(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccessNilData())
}

func (s *CartHandler) DeleteCartItem(ctx echo.Context) error {
	var req entity.DeleteCartItemRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	err := s.uc.DeleteCartItem(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccessNilData())
}

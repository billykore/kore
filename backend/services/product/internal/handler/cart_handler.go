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

func (h *CartHandler) GetCartItemList(ctx echo.Context) error {
	var req entity.CartRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	items, err := h.uc.GetCartItemList(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(items))
}

func (h *CartHandler) AddCartItem(ctx echo.Context) error {
	var req entity.AddCartItemRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	err := h.uc.AddCartItem(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccessNilData())
}

func (h *CartHandler) UpdateCartItemQuantity(ctx echo.Context) error {
	var req entity.UpdateCartItemRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	err := h.uc.UpdateCartItemQuantity(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccessNilData())
}

func (h *CartHandler) DeleteCartItem(ctx echo.Context) error {
	var req entity.DeleteCartItemRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	err := h.uc.DeleteCartItem(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccessNilData())
}

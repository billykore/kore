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

// GetCartItemList swaggo annotation.
//
//	@Summary		Cart item list
//	@Description	Get list of cart items
//	@Tags			product-service
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		integer	true	"Number of cart to display"
//	@Param			startId	query		integer	true	"ID of cart for begin to display"
//	@Success		200		{object}	entity.Response
//	@Failure		400		{object}	entity.Response
//	@Failure		401		{object}	entity.Response
//	@Failure		404		{object}	entity.Response
//	@Failure		500		{object}	entity.Response
//	@Router			/carts [get]
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

// AddCartItem swaggo annotation.
//
//	@Summary		Add cart item
//	@Description	Add new item to cart
//	@Tags			product-service
//	@Accept			json
//	@Produce		json
//	@Param			AddCartItemRequest	body		entity.AddCartItemRequest	true	"Add cart item request"
//	@Success		200					{object}	entity.Response
//	@Failure		400					{object}	entity.Response
//	@Failure		401					{object}	entity.Response
//	@Failure		404					{object}	entity.Response
//	@Failure		500					{object}	entity.Response
//	@Router			/carts [post]
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

// UpdateCartItemQuantity swaggo annotation.
//
//	@Summary		Update cart item
//	@Description	Update existing cart item
//	@Tags			product-service
//	@Accept			json
//	@Produce		json
//	@Param			cartId					path		integer							true	"Cart ID"
//	@Param			UpdateCartItemRequest	body		entity.UpdateCartItemRequest	true	"Update cart item request"
//	@Success		200						{object}	entity.Response
//	@Failure		400						{object}	entity.Response
//	@Failure		401						{object}	entity.Response
//	@Failure		404						{object}	entity.Response
//	@Failure		500						{object}	entity.Response
//	@Router			/carts/{cartId} [put]
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

// DeleteCartItem swaggo annotation.
//
//	@Summary		Delete specific cart
//	@Description	Delete a cart by ID
//	@Tags			product-service
//	@Accept			json
//	@Produce		json
//	@Param			cartId	path		integer	true	"Cart ID"
//	@Success		200		{object}	entity.Response
//	@Failure		400		{object}	entity.Response
//	@Failure		401		{object}	entity.Response
//	@Failure		404		{object}	entity.Response
//	@Failure		500		{object}	entity.Response
//	@Router			/carts/{id} [delete]
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

package handler

import (
	"github.com/billykore/kore/backend/internal/app/product"
	"github.com/billykore/kore/backend/pkg/entity"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	svc *product.Service
}

func NewProductHandler(svc *product.Service) *ProductHandler {
	return &ProductHandler{
		svc: svc,
	}
}

// GetProductList swaggo annotation.
//
//	@Summary		List of products
//	@Description	Get list of products
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		integer	true	"Number of products to display"
//	@Param			startId	query		integer	true	"ID of products for begin to display"
//	@Success		200		{object}	entity.Response
//	@Failure		400		{object}	entity.Response
//	@Failure		401		{object}	entity.Response
//	@Failure		404		{object}	entity.Response
//	@Failure		500		{object}	entity.Response
//	@Router			/products [get]
func (h *ProductHandler) GetProductList(ctx echo.Context) error {
	var req product.GetRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	products, err := h.svc.GetProductList(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(products))
}

// GetProductById swaggo annotation.
//
//	@Summary		Get specific product
//	@Description	Get product by ID
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			productId	path		integer	true	"Product ID"
//	@Success		200			{object}	entity.Response
//	@Failure		400			{object}	entity.Response
//	@Failure		401			{object}	entity.Response
//	@Failure		404			{object}	entity.Response
//	@Failure		500			{object}	entity.Response
//	@Router			/products/{productId} [get]
func (h *ProductHandler) GetProductById(ctx echo.Context) error {
	var req product.GetRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	res, err := h.svc.GetProductById(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(res))
}

// GetCategoryList swaggo annotation.
//
//	@Summary		List of product categories
//	@Description	Get list of product categories
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	entity.Response
//	@Failure		400	{object}	entity.Response
//	@Failure		401	{object}	entity.Response
//	@Failure		404	{object}	entity.Response
//	@Failure		500	{object}	entity.Response
//	@Router			/categories [get]
func (h *ProductHandler) GetCategoryList(ctx echo.Context) error {
	categories, err := h.svc.GetCategoryList(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(categories))
}

// GetDiscountList swaggo annotation.
//
//	@Summary		List of discounts
//	@Description	Get list of discounts
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		integer	true	"Number of discounts to display"
//	@Param			startId	query		integer	true	"ID of discount for begin to display"
//	@Success		200		{object}	entity.Response
//	@Failure		400		{object}	entity.Response
//	@Failure		401		{object}	entity.Response
//	@Failure		404		{object}	entity.Response
//	@Failure		500		{object}	entity.Response
//	@Router			/discounts [get]
func (h *ProductHandler) GetDiscountList(ctx echo.Context) error {
	var req product.DiscountRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	discounts, err := h.svc.GetDiscountList(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(discounts))
}

// GetCartItemList swaggo annotation.
//
//	@Summary		Cart item list
//	@Description	Get list of cart items
//	@Tags			product
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
func (h *ProductHandler) GetCartItemList(ctx echo.Context) error {
	var req product.CartRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	items, err := h.svc.GetCartItemList(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(items))
}

// AddCartItem swaggo annotation.
//
//	@Summary		Add cart item
//	@Description	Add new item to cart
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			AddCartItemRequest	body		product.AddCartItemRequest	true	"Add cart item request"
//	@Success		200					{object}	entity.Response
//	@Failure		400					{object}	entity.Response
//	@Failure		401					{object}	entity.Response
//	@Failure		404					{object}	entity.Response
//	@Failure		500					{object}	entity.Response
//	@Router			/carts [post]
func (h *ProductHandler) AddCartItem(ctx echo.Context) error {
	var req product.AddCartItemRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	err := h.svc.AddCartItem(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccessNilData())
}

// UpdateCartItemQuantity swaggo annotation.
//
//	@Summary		Update cart item
//	@Description	Update existing cart item
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			cartId					path		integer							true	"Cart ID"
//	@Param			UpdateCartItemRequest	body		product.UpdateCartItemRequest	true	"Update cart item request"
//	@Success		200						{object}	entity.Response
//	@Failure		400						{object}	entity.Response
//	@Failure		401						{object}	entity.Response
//	@Failure		404						{object}	entity.Response
//	@Failure		500						{object}	entity.Response
//	@Router			/carts/{cartId} [put]
func (h *ProductHandler) UpdateCartItemQuantity(ctx echo.Context) error {
	var req product.UpdateCartItemRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	err := h.svc.UpdateCartItemQuantity(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccessNilData())
}

// DeleteCartItem swaggo annotation.
//
//	@Summary		Delete specific cart
//	@Description	Delete a cart by ID
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			cartId	path		integer	true	"Cart ID"
//	@Success		200		{object}	entity.Response
//	@Failure		400		{object}	entity.Response
//	@Failure		401		{object}	entity.Response
//	@Failure		404		{object}	entity.Response
//	@Failure		500		{object}	entity.Response
//	@Router			/carts/{id} [delete]
func (h *ProductHandler) DeleteCartItem(ctx echo.Context) error {
	var req product.DeleteCartItemRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	err := h.svc.DeleteCartItem(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccessNilData())
}

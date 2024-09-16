package handler

import (
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/services/product/internal/usecase"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	uc *usecase.ProductUsecase
}

func NewProductHandler(uc *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{uc: uc}
}

// GetProductList swaggo annotation.
//
//	@Summary		List of products
//	@Description	Get list of products
//	@Tags			product-service
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
	var req entity.ProductRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	products, err := h.uc.GetProductList(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(products))
}

// GetProductById swaggo annotation.
//
//	@Summary		Get specific product
//	@Description	Get product by ID
//	@Tags			product-service
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
	var req entity.ProductRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	product, err := h.uc.GetProductById(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(product))
}

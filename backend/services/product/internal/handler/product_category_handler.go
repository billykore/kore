package handler

import (
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/services/product/internal/usecase"
	"github.com/labstack/echo/v4"
)

type ProductCategoryHandler struct {
	uc *usecase.ProductCategoryUsecase
}

func NewProductCategoryHandler(uc *usecase.ProductCategoryUsecase) *ProductCategoryHandler {
	return &ProductCategoryHandler{
		uc: uc,
	}
}

// GetCategoryList swaggo annotation.
//
//	@Summary		List of product categories
//	@Description	Get list of product categories
//	@Tags			product-service
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	entity.Response
//	@Failure		400	{object}	entity.Response
//	@Failure		401	{object}	entity.Response
//	@Failure		404	{object}	entity.Response
//	@Failure		500	{object}	entity.Response
//	@Router			/categories [get]
func (h *ProductCategoryHandler) GetCategoryList(ctx echo.Context) error {
	categories, err := h.uc.GetCategoryList(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(categories))
}

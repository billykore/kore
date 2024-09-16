package handler

import (
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/services/product/internal/usecase"
	"github.com/labstack/echo/v4"
)

type DiscountHandler struct {
	uc *usecase.DiscountUsecase
}

func NewDiscountHandler(uc *usecase.DiscountUsecase) *DiscountHandler {
	return &DiscountHandler{
		uc: uc,
	}
}

// GetDiscountList swaggo annotation.
//
//	@Summary		List of discounts
//	@Description	Get list of discounts
//	@Tags			product-service
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
func (h *DiscountHandler) GetDiscountList(ctx echo.Context) error {
	var req entity.DiscountRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	discounts, err := h.uc.GetDiscountList(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(discounts))
}

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

func (s *DiscountHandler) GetDiscountList(ctx echo.Context) error {
	var req entity.DiscountRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	discounts, err := s.uc.GetDiscountList(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(discounts))
}

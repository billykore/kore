package service

import (
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/services/product/usecase"
	"github.com/labstack/echo/v4"
)

type DiscountService struct {
	uc *usecase.DiscountUsecase
}

func NewDiscountService(uc *usecase.DiscountUsecase) *DiscountService {
	return &DiscountService{
		uc: uc,
	}
}

func (s *DiscountService) GetDiscountList(ctx echo.Context) error {
	var req entity.DiscountRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	discounts, err := s.uc.GetDiscountList(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("discounts", discounts))
}

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

func (s *ProductCategoryHandler) GetCategoryList(ctx echo.Context) error {
	categories, err := s.uc.GetCategoryList(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("categories", categories))
}

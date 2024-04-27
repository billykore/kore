package service

import (
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/services/product/usecase"
	"github.com/labstack/echo/v4"
)

type ProductCategoryService struct {
	uc *usecase.ProductCategoryUsecase
}

func NewProductCategoryService(uc *usecase.ProductCategoryUsecase) *ProductCategoryService {
	return &ProductCategoryService{
		uc: uc,
	}
}

func (s *ProductCategoryService) GetCategoryList(ctx echo.Context) error {
	categories, err := s.uc.GetCategoryList(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("categories", categories))
}

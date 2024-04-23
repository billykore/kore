package service

import (
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/services/product/usecase"
	"github.com/labstack/echo/v4"
)

type ProductService struct {
	uc *usecase.ProductUsecase
}

func NewProductService(uc *usecase.ProductUsecase) *ProductService {
	return &ProductService{uc: uc}
}

func (s *ProductService) ProductList(ctx echo.Context) error {
	products, err := s.uc.ProductList(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(products))
}

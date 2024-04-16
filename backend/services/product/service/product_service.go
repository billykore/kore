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

func (s *ProductService) Greet(ctx echo.Context) error {
	in := new(entity.ProductRequest)
	err := ctx.Bind(in)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	greet, err := s.uc.Greet(ctx.Request().Context(), in)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(greet))
}
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

func (s *ProductService) GetProductList(ctx echo.Context) error {
	var req entity.ProductRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	products, err := s.uc.GetProductList(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("products", products))
}

func (s *ProductService) GetProductById(ctx echo.Context) error {
	var req entity.ProductRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	product, err := s.uc.GetProductById(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("product", product))
}

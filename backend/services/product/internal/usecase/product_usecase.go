package usecase

import (
	"context"

	"github.com/billykore/kore/backend/pkg/codes"
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/pkg/status"
	"github.com/billykore/kore/backend/services/product/internal/repo"
)

type ProductUsecase struct {
	log         *log.Logger
	productRepo *repo.ProductRepository
}

func NewProductUsecase(log *log.Logger, productRepo *repo.ProductRepository) *ProductUsecase {
	return &ProductUsecase{
		log:         log,
		productRepo: productRepo,
	}
}

func (uc *ProductUsecase) GetProductList(ctx context.Context, req entity.ProductRequest) ([]*entity.ProductResponse, error) {
	products, err := uc.productRepo.List(ctx, req.CategoryId, req.Limit, req.StartId)
	if err != nil {
		uc.log.Usecase("ProductList").Error(err)
		return nil, status.Error(codes.NotFound, err.Error())
	}
	resp := make([]*entity.ProductResponse, 0)
	for _, p := range products {
		resp = append(resp, entity.MakeProductResponse(p))
	}
	return resp, nil
}

func (uc *ProductUsecase) GetProductById(ctx context.Context, req entity.ProductRequest) (*entity.ProductResponse, error) {
	product, err := uc.productRepo.GetById(ctx, req.ProductId)
	if err != nil {
		uc.log.Usecase("GetProductById").Error(err)
		return nil, status.Error(codes.NotFound, err.Error())
	}
	resp := entity.MakeProductResponse(product)
	return resp, nil
}

package usecase

import (
	"context"

	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/services/product/internal/repo"
)

type ProductCategoryUsecase struct {
	log                 *log.Logger
	productCategoryRepo *repo.ProductCategoryRepository
}

func NewProductCategoryUsecase(log *log.Logger, productCategoryRepo *repo.ProductCategoryRepository) *ProductCategoryUsecase {
	return &ProductCategoryUsecase{
		log:                 log,
		productCategoryRepo: productCategoryRepo,
	}
}

func (uc *ProductCategoryUsecase) GetCategoryList(ctx context.Context) ([]*entity.ProductCategoryResponse, error) {
	categories, err := uc.productCategoryRepo.List(ctx)
	if err != nil {
		uc.log.Usecase("CategoryList").Error(err)
		return nil, err
	}
	resp := make([]*entity.ProductCategoryResponse, 0)
	for _, c := range categories {
		resp = append(resp, entity.MakeProductCategoryResponse(c))
	}
	return resp, nil
}

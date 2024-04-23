package usecase

import (
	"context"

	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/pkg/repo"
)

type ProductUsecase struct {
	log  *log.Logger
	repo repo.ProductRepository
}

func NewProductUsecase(log *log.Logger, repo repo.ProductRepository) *ProductUsecase {
	return &ProductUsecase{
		log:  log,
		repo: repo,
	}
}

func (uc *ProductUsecase) ProductList(ctx context.Context) ([]entity.ProductResponse, error) {
	products, err := uc.repo.List(ctx)
	if err != nil {
		uc.log.Usecase("ProductList").Error(err)
		return nil, err
	}
	var resp []entity.ProductResponse
	for _, product := range products {
		resp = append(resp, entity.MakeProductResponse(product))
	}
	return resp, nil
}

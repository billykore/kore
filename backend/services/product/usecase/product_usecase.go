package usecase

import (
	"context"

	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/pkg/repo"
)

type ProductUsecase struct {
	log  *log.Logger
	repo repo.GreeterRepository
}

func NewProductUsecase(log *log.Logger, repo repo.GreeterRepository) *ProductUsecase {
	return &ProductUsecase{
		log:  log,
		repo: repo,
	}
}

func (uc *ProductUsecase) Greet(ctx context.Context, req *entity.ProductRequest) (*entity.ProductResponse, error) {
	uc.log.Usecase("Greet").Infof("Greet %s", req.Name)
	return &entity.ProductResponse{
		Message: "Hello " + req.Name,
	}, nil
}

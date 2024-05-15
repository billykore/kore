package usecase

import (
	"context"

	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/pkg/repo"
)

type ShippingUsecase struct {
	log  *log.Logger
	repo repo.GreeterRepository
}

func NewShippingUsecase(log *log.Logger, repo repo.GreeterRepository) *ShippingUsecase {
	return &ShippingUsecase{
		log:  log,
		repo: repo,
	}
}

func (uc *ShippingUsecase) Greet(ctx context.Context, req *entity.ShippingRequest) (*entity.ShippingResponse, error) {
	uc.log.Usecase("Greet").Infof("Greet %s", req.Name)
	return &entity.ShippingResponse{
		Message: "Hello " + req.Name,
	}, nil
}

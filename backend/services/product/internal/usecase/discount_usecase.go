package usecase

import (
	"context"

	"github.com/billykore/kore/backend/pkg/codes"
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/pkg/status"
	"github.com/billykore/kore/backend/services/product/internal/repo"
)

type DiscountUsecase struct {
	log          *log.Logger
	discountRepo *repo.DiscountRepository
}

func NewDiscountUsecase(discountRepo *repo.DiscountRepository) *DiscountUsecase {
	return &DiscountUsecase{
		discountRepo: discountRepo,
	}
}

func (uc *DiscountUsecase) GetDiscountList(ctx context.Context, req entity.DiscountRequest) ([]*entity.DiscountResponse, error) {
	discounts, err := uc.discountRepo.List(ctx, req.Limit, req.StartId)
	if err != nil {
		uc.log.Usecase("GetDiscountList").Error(err)
		return nil, status.Error(codes.BadRequest, err.Error())
	}
	resp := make([]*entity.DiscountResponse, 0)
	for _, d := range discounts {
		resp = append(resp, entity.MakeDiscountResponse(d))
	}
	return resp, nil
}

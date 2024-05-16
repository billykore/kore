package usecase

import (
	"context"

	"github.com/billykore/kore/backend/pkg/codes"
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/repo"
	"github.com/billykore/kore/backend/pkg/status"
)

type ShippingUsecase struct {
	log  *log.Logger
	repo repo.ShippingRepository
}

func NewShippingUsecase(log *log.Logger, repo repo.ShippingRepository) *ShippingUsecase {
	return &ShippingUsecase{
		log:  log,
		repo: repo,
	}
}

func (uc *ShippingUsecase) CreateShipping(ctx context.Context, req entity.CreateShippingRequest) (*entity.CreateShippingResponse, error) {
	fee := entity.GetShippingFee(req.ShippingType)
	id, err := uc.repo.Save(ctx, model.Shipping{
		ShipperName:     req.ShipperName,
		ShippingType:    req.ShippingType,
		CustomerAddress: req.Address,
		CustomerName:    req.CustomerName,
		SenderName:      req.SenderName,
		Status:          model.ShippingStatusCreated,
		Fee:             fee,
	})
	if err != nil {
		uc.log.Usecase("CreateShipping").Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &entity.CreateShippingResponse{
		Id:          id,
		Fee:         fee,
		Status:      model.ShippingStatusCreated,
		ShipperName: req.ShipperName,
	}, nil
}

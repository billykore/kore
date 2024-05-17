package usecase

import (
	"context"

	"github.com/billykore/kore/backend/pkg/broker/rabbit"
	"github.com/billykore/kore/backend/pkg/codes"
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/repo"
	"github.com/billykore/kore/backend/pkg/status"
)

type ShippingUsecase struct {
	log    *log.Logger
	rabbit *rabbit.Rabbit
	repo   repo.ShippingRepository
}

func NewShippingUsecase(log *log.Logger, rabbit *rabbit.Rabbit, repo repo.ShippingRepository) *ShippingUsecase {
	return &ShippingUsecase{
		log:    log,
		rabbit: rabbit,
		repo:   repo,
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
		Status:          model.ShippingStatusCreated.String(),
		Fee:             fee,
	})
	if err != nil {
		uc.log.Usecase("CreateShipping").Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &entity.CreateShippingResponse{
		Id:          id,
		Fee:         fee,
		Status:      model.ShippingStatusCreated.String(),
		ShipperName: req.ShipperName,
	}, nil
}

func (uc *ShippingUsecase) UpdateShippingStatus(ctx context.Context, req entity.UpdateShippingStatusRequest) error {
	err := uc.repo.UpdateStatus(ctx, req.Id, model.ShippingStatus(req.NewStatus), model.ShippingStatus(req.CurrentStatus))
	if err != nil {
		uc.log.Usecase("UpdateShippingStatus").Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	data := &entity.UpdateShippingPublishData{
		ShippingId: req.Id,
		Status:     req.NewStatus,
	}
	payload := rabbit.NewPayload("shipping-service", data)
	bytePayload, err := payload.MarshalBinary()
	if err != nil {
		uc.log.Usecase("UpdateShippingPublishData").Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	err = uc.rabbit.Publish(ctx, bytePayload)
	if err != nil {
		uc.log.Usecase("UpdateShippingStatus").Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

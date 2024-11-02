package shipping

import (
	"context"

	"github.com/billykore/kore/backend/internal/infra/messaging/rabbitmq"
	"github.com/billykore/kore/backend/pkg/codes"
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/logger"
	"github.com/billykore/kore/backend/pkg/status"
)

type Repository interface {
	GetById(ctx context.Context, id uint) (*Shipping, error)
	Save(ctx context.Context, shipping Shipping) (uint, error)
	UpdateStatus(ctx context.Context, id uint, newStatus, currentStatus Status) error
}

type Service struct {
	log        *logger.Logger
	rabbitConn *rabbitmq.Connection
	repo       Repository
}

func NewService(log *logger.Logger, rabbitConn *rabbitmq.Connection, repo Repository) *Service {
	return &Service{
		log:        log,
		rabbitConn: rabbitConn,
		repo:       repo,
	}
}

func (uc *Service) CreateShipping(ctx context.Context, req CreateShippingRequest) (*CreateShippingResponse, error) {
	fee := GetFee(req.ShippingType)
	id, err := uc.repo.Save(ctx, Shipping{
		ShipperName:     req.ShipperName,
		ShippingType:    req.ShippingType,
		CustomerAddress: req.Address,
		CustomerName:    req.CustomerName,
		SenderName:      req.SenderName,
		Status:          StatusCreated.String(),
		Fee:             fee,
	})
	if err != nil {
		uc.log.Usecase("CreateShipping").Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &CreateShippingResponse{
		Id:          id,
		Fee:         fee,
		Status:      StatusCreated.String(),
		ShipperName: req.ShipperName,
	}, nil
}

func (uc *Service) UpdateShippingStatus(ctx context.Context, req UpdateShippingStatusRequest) (*entity.MessagePayload[*rabbitmq.UpdateShippingRabbitData], error) {
	s, err := uc.repo.GetById(ctx, req.Id)
	if err != nil {
		uc.log.Usecase("UpdateShippingStatus").Error(err)
		return nil, status.Error(codes.NotFound, err.Error())
	}
	err = uc.repo.UpdateStatus(ctx, s.ID, Status(req.NewStatus), Status(req.CurrentStatus))
	if err != nil {
		uc.log.Usecase("UpdateShippingStatus").Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	data := &rabbitmq.UpdateShippingRabbitData{
		ShippingId: req.Id,
		Status:     req.NewStatus,
	}
	msgPayload := &entity.MessagePayload[*rabbitmq.UpdateShippingRabbitData]{
		Origin: "shipping-service",
		Data:   data,
	}
	return msgPayload, nil
}

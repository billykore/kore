package shipping

import (
	"context"

	"github.com/billykore/kore/backend/pkg/codes"
	"github.com/billykore/kore/backend/pkg/logger"
	"github.com/billykore/kore/backend/pkg/status"
)

type Messaging interface {
	ProduceStatusChange(context.Context, StatusChangeData) error
}

type Repository interface {
	GetById(ctx context.Context, id uint) (*Shipping, error)
	Save(ctx context.Context, shipping Shipping) (uint, error)
	UpdateStatus(ctx context.Context, id uint, newStatus, currentStatus Status) error
}

type Service struct {
	log  *logger.Logger
	repo Repository
	msg  Messaging
}

func NewService(log *logger.Logger, repo Repository) *Service {
	return &Service{
		log:  log,
		repo: repo,
	}
}

func (s *Service) CreateShipping(ctx context.Context, req CreateShippingRequest) (*CreateShippingResponse, error) {
	fee := GetFee(req.ShippingType)
	id, err := s.repo.Save(ctx, Shipping{
		ShipperName:     req.ShipperName,
		ShippingType:    req.ShippingType,
		CustomerAddress: req.Address,
		CustomerName:    req.CustomerName,
		SenderName:      req.SenderName,
		Status:          StatusCreated.String(),
		Fee:             fee,
	})
	if err != nil {
		s.log.Usecase("CreateShipping").Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &CreateShippingResponse{
		Id:          id,
		Fee:         fee,
		Status:      StatusCreated.String(),
		ShipperName: req.ShipperName,
	}, nil
}

func (s *Service) UpdateShippingStatus(ctx context.Context, req UpdateShippingStatusRequest) error {
	shipping, err := s.repo.GetById(ctx, req.Id)
	if err != nil {
		s.log.Usecase("UpdateShippingStatus").Error(err)
		return status.Error(codes.NotFound, err.Error())
	}
	err = s.repo.UpdateStatus(ctx, shipping.ID, Status(req.NewStatus), Status(req.CurrentStatus))
	if err != nil {
		s.log.Usecase("UpdateShippingStatus").Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	err = s.msg.ProduceStatusChange(ctx, StatusChangeData{
		ShippingId: shipping.ID,
		Status:     req.NewStatus,
	})
	if err != nil {
		s.log.Usecase("UpdateShippingStatus").Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

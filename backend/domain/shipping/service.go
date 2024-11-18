package shipping

import (
	"context"

	"github.com/billykore/kore/backend/pkg/codes"
	"github.com/billykore/kore/backend/pkg/logger"
	"github.com/billykore/kore/backend/pkg/status"
)

// Messaging is interface for messaging service used by shipping domain.
type Messaging interface {
	// PublishStatusChange publish shipping status change.
	PublishStatusChange(context.Context, StatusChangeData) error
}

// Repository defines the methods to interacting with persistence storage used by shipping domain.
type Repository interface {
	// GetById gets specific shipping by ID.
	GetById(ctx context.Context, id uint) (*Shipping, error)

	// Save saves new shipping.
	Save(ctx context.Context, shipping Shipping) (uint, error)

	// UpdateStatus updates existing shipping status.
	UpdateStatus(ctx context.Context, id uint, newStatus, currentStatus Status) error
}

type Service struct {
	log  *logger.Logger
	repo Repository
	msg  Messaging
}

func NewService(log *logger.Logger, repo Repository, msg Messaging) *Service {
	return &Service{
		log:  log,
		repo: repo,
		msg:  msg,
	}
}

func (s *Service) CreateShipping(ctx context.Context, req CreateShippingRequest) (*CreateShippingResponse, error) {
	fee := getFee(req.ShippingType)
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
	err = s.msg.PublishStatusChange(ctx, StatusChangeData{
		ShippingId: shipping.ID,
		Status:     req.NewStatus,
	})
	if err != nil {
		s.log.Usecase("UpdateShippingStatus").Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

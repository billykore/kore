package order

import (
	"context"
	"errors"

	"github.com/billykore/kore/backend/pkg/codes"
	"github.com/billykore/kore/backend/pkg/ctxt"
	"github.com/billykore/kore/backend/pkg/logger"
	"github.com/billykore/kore/backend/pkg/status"
	"github.com/billykore/kore/backend/pkg/types"
)

// PaymentService is interface for payment service used by order domain.
type PaymentService interface {
	// Pay pays the order.
	Pay(srcName, srcAccount string, amount types.Money) (*PaymentResponse, error)
}

// ShippingService is interface for shipping service used by order domain.
type ShippingService interface {
	// Ship ships the order.
	Ship(ShippingData) (*ShippingResponse, error)
}

// Repository defines the methods to interacting with persistence storage used by order domain.
type Repository interface {
	// GetById gets specific order by ID.
	GetById(ctx context.Context, id uint) (*Order, error)

	// GetByIdAndStatus gets specific order by ID and status.
	GetByIdAndStatus(ctx context.Context, id uint, status Status) (*Order, error)

	// GetByShippingId gets specific order by shipping ID.
	GetByShippingId(ctx context.Context, shippingId uint) (*Order, error)

	// Save saves new order.
	Save(ctx context.Context, order Order) error

	// UpdateStatus updates order status.
	UpdateStatus(ctx context.Context, id uint, newStatus Status, currentStatus ...Status) error

	// UpdateShipping updates order's shipping status.
	UpdateShipping(ctx context.Context, id uint, shippingId int) error
}

type Service struct {
	log         *logger.Logger
	repo        Repository
	paymentSvc  PaymentService
	shippingSvc ShippingService
}

func NewService(log *logger.Logger, repo Repository, paymentSvc PaymentService, shippingSvc ShippingService) *Service {
	return &Service{
		log:         log,
		repo:        repo,
		paymentSvc:  paymentSvc,
		shippingSvc: shippingSvc,
	}
}

func (s *Service) Checkout(ctx context.Context, req CheckoutRequest) error {
	user, ok := ctxt.UserFromContext(ctx)
	if !ok {
		s.log.Usecase("Checkout").Error(errors.New("failed to get user from context"))
		return status.Error(codes.Internal, messageCheckoutFailed)
	}

	newOrder := Order{
		Username:      user.Username,
		PaymentMethod: req.PaymentMethod,
		Status:        StatusCreated,
	}
	newOrder.SetCartIds(req.cartIds())

	err := s.repo.Save(ctx, newOrder)
	if err != nil {
		s.log.Usecase("Checkout").Errorf("failed to save order: %v", err)
		return status.Error(codes.Internal, messageCheckoutFailed)
	}

	return nil
}

func (s *Service) GetOrderById(ctx context.Context, req GetRequest) (*Response, error) {
	order, err := s.repo.GetById(ctx, req.Id)
	if err != nil {
		s.log.Usecase("GetOrderById").Errorf("failed to get order by id (%d): %v", req.Id, err)
		return nil, status.Error(codes.NotFound, messageOrderNotFound)
	}
	return &Response{
		Id:            order.ID,
		Username:      order.Username,
		CartIds:       order.IntCartIds(),
		PaymentMethod: order.PaymentMethod,
		Status:        order.Status.String(),
		ShippingId:    order.ShippingId,
	}, nil
}

func (s *Service) PayOrder(ctx context.Context, req PaymentRequest) (*PaymentResponse, error) {
	order, err := s.repo.GetByIdAndStatus(ctx, req.Id, StatusCreated)
	if err != nil {
		s.log.Usecase("OrderPayment").
			Errorf("failed to get order by id (%d) and status (%s): %v", req.Id, StatusCreated.String(), err)
		return nil, status.Error(codes.NotFound, messagePayOrderFailed)
	}
	paymentResp, err := s.paymentSvc.Pay(req.AccountName, req.AccountNumber, order.TotalPrice())
	if err != nil {
		s.log.Usecase("OrderPayment").Errorf("failed to pay order: %v", err)
		return nil, status.Error(codes.Internal, messagePayOrderFailed)
	}
	err = s.repo.UpdateStatus(ctx, req.Id, StatusPaymentSucceed, StatusCreated)
	if err != nil {
		s.log.Usecase("UpdateOrderStatus").Errorf("failed to update order status: %v", err)
		return nil, status.Error(codes.Internal, messagePayOrderFailed)
	}
	return paymentResp, nil
}

func (s *Service) ShipOrder(ctx context.Context, req ShippingRequest) (*ShippingResponse, error) {
	order, err := s.repo.GetByIdAndStatus(ctx, req.OrderId, StatusPaymentSucceed)
	if err != nil {
		s.log.Usecase("ShipOrder").
			Errorf("failed to get order by id (%d) and status (%s): %v", req.OrderId, StatusWaitingForPayment.String(), err)
		return nil, status.Error(codes.NotFound, messageShipOrderFailed)
	}
	shippingResp, err := s.shippingSvc.Ship(ShippingData{
		Address:      req.Address,
		CustomerName: req.CustomerName,
		ShippingType: req.ShippingType,
	})
	if err != nil {
		s.log.Usecase("ShipOrder").Errorf("failed to ship order: %v", err)
		return nil, status.Error(codes.Internal, messageShipOrderFailed)
	}
	err = s.repo.UpdateShipping(ctx, order.ID, shippingResp.Id)
	if err != nil {
		s.log.Usecase("ShipOrder").Errorf("failed to update shipping status: %v", err)
		return nil, status.Error(codes.Internal, messageShipOrderFailed)
	}
	return shippingResp, nil
}

func (s *Service) CancelOrder(ctx context.Context, req CancelOrderRequest) error {
	err := s.repo.UpdateStatus(ctx, req.OrderId, StatusCancelled, StatusCanCancel...)
	if err != nil {
		s.log.Usecase("CancelOrder").Errorf("failed to cancel order: %v", err)
		return status.Error(codes.Internal, messageCancelOrderFailed)
	}
	return nil
}

func (s *Service) ConsumeOrderStatusChanges(ctx context.Context, data StatusChangeData) error {
	order, err := s.repo.GetByShippingId(ctx, data.ShippingId)
	if err != nil {
		s.log.Usecase("ListenOrderStatusChanges").
			Errorf("failed to get order by shipping id (%d): %v", data.ShippingId, err)
		return err
	}
	err = s.repo.UpdateStatus(ctx, order.ID, Status(data.Status), StatusWaitingForShipment)
	if err != nil {
		s.log.Usecase("ListenOrderStatusChanges").
			Errorf("failed to update order status: %v", err)
		return err
	}
	return nil
}

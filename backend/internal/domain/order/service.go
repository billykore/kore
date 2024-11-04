package order

import (
	"context"
	"errors"

	"github.com/billykore/kore/backend/internal/infra/shipping"
	"github.com/billykore/kore/backend/internal/infra/transaction"
	"github.com/billykore/kore/backend/pkg/codes"
	"github.com/billykore/kore/backend/pkg/ctxt"
	"github.com/billykore/kore/backend/pkg/logger"
	"github.com/billykore/kore/backend/pkg/status"
)

type Repository interface {
	GetById(ctx context.Context, id uint) (*Order, error)
	GetByIdAndStatus(ctx context.Context, id uint, status Status) (*Order, error)
	GetByShippingId(ctx context.Context, shippingId uint) (*Order, error)
	Save(ctx context.Context, order Order) error
	UpdateStatus(ctx context.Context, id uint, newStatus Status, currentStatus ...Status) error
	UpdateShipping(ctx context.Context, id uint, shippingId int) error
}

type Service struct {
	log  *logger.Logger
	repo Repository
}

func NewService(log *logger.Logger, repo Repository) *Service {
	return &Service{
		log:  log,
		repo: repo,
	}
}

func (s *Service) Checkout(ctx context.Context, req CheckoutRequest) error {
	user, ok := ctxt.UserFromContext(ctx)
	if !ok {
		s.log.Usecase("Checkout").Error(errors.New("failed to get user from context"))
		return status.Error(codes.Internal, "Failed checkout order")
	}

	newOrder := Order{
		Username:      user.Username,
		PaymentMethod: req.PaymentMethod,
		Status:        StatusCreated,
	}
	newOrder.SetCartIds(req.CartIds())

	err := s.repo.Save(ctx, newOrder)
	if err != nil {
		s.log.Usecase("Checkout").Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func (s *Service) GetOrderById(ctx context.Context, req GetRequest) (*Response, error) {
	o, err := s.repo.GetById(ctx, req.Id)
	if err != nil {
		s.log.Usecase("GetOrderById").Error(err)
		return nil, status.Error(codes.NotFound, err.Error())
	}
	resp := MakeResponse(o)
	return resp, nil
}

func (s *Service) PayOrder(ctx context.Context, req PaymentRequest) (*PaymentResponse, error) {
	o, err := s.repo.GetByIdAndStatus(ctx, req.Id, StatusWaitingForPayment)
	if err != nil {
		s.log.Usecase("OrderPayment").Error(err)
		return nil, status.Error(codes.NotFound, err.Error())
	}
	paymentMethod := transaction.NewPayment(req.Method, req.AccountName, req.AccountNumber)
	paymentResp, err := paymentMethod.Pay(o.TotalPrice())
	if err != nil {
		s.log.Usecase("OrderPayment").Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = s.repo.UpdateStatus(ctx, req.Id, StatusPaymentSucceed, StatusWaitingForPayment)
	if err != nil {
		s.log.Usecase("UpdateOrderStatus").Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	resp := MakePaymentResponse(paymentResp)
	return resp, nil
}

func (s *Service) ShipOrder(ctx context.Context, req ShippingRequest) (*ShippingResponse, error) {
	o, err := s.repo.GetByIdAndStatus(ctx, req.OrderId, StatusPaymentSucceed)
	if err != nil {
		s.log.Usecase("ShipOrder").Error(err)
		return nil, status.Error(codes.NotFound, err.Error())
	}
	shipper := shipping.New(req.ShipperName)
	shippingResp, err := shipper.Create(&shipping.Data{
		Address:      req.Address,
		CustomerName: req.CustomerName,
		ShippingType: req.ShippingType,
	})
	if err != nil {
		s.log.Usecase("ShipOrder").Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = s.repo.UpdateShipping(ctx, o.ID, shippingResp.Id)
	if err != nil {
		s.log.Usecase("ShipOrder").Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	resp := MakeShippingResponse(shippingResp)
	return resp, nil
}

func (s *Service) CancelOrder(ctx context.Context, req CancelOrderRequest) error {
	err := s.repo.UpdateStatus(ctx, req.OrderId, StatusCancelled, StatusCanCancel...)
	if err != nil {
		s.log.Usecase("CancelOrder").Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (s *Service) ConsumeOrderStatusChanges(ctx context.Context, data StatusChangeData) error {
	order, err := s.repo.GetByShippingId(ctx, data.ShippingId)
	if err != nil {
		s.log.Usecase("ListenOrderStatusChanges").Error(err)
		return err
	}
	err = s.repo.UpdateStatus(ctx, order.ID, Status(data.Status), StatusWaitingForShipment)
	if err != nil {
		s.log.Usecase("ListenOrderStatusChanges").Error(err)
		return err
	}
	return nil
}

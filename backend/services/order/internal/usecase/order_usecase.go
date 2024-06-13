package usecase

import (
	"context"

	"github.com/billykore/kore/backend/pkg/codes"
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/net/rabbit"
	"github.com/billykore/kore/backend/pkg/repo"
	"github.com/billykore/kore/backend/pkg/shipping"
	"github.com/billykore/kore/backend/pkg/status"
	"github.com/billykore/kore/backend/pkg/transaction"
	amqp "github.com/rabbitmq/amqp091-go"
)

type OrderUsecase struct {
	log  *log.Logger
	repo repo.OrderRepository
}

func NewOrderUsecase(log *log.Logger, repo repo.OrderRepository) *OrderUsecase {
	return &OrderUsecase{
		log:  log,
		repo: repo,
	}
}

func (uc *OrderUsecase) Checkout(ctx context.Context, req entity.CheckoutRequest) error {
	newOrder := model.Order{
		UserId:        req.UserId,
		PaymentMethod: req.PaymentMethod,
		Status:        model.OrderStatusCreated,
	}
	err := uc.repo.Save(ctx, newOrder)
	if err != nil {
		uc.log.Usecase("Checkout").Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (uc *OrderUsecase) GetOrderById(ctx context.Context, req entity.OrderRequest) (*entity.OrderResponse, error) {
	order, err := uc.repo.GetById(ctx, req.Id)
	if err != nil {
		uc.log.Usecase("GetOrderById").Error(err)
		return nil, status.Error(codes.NotFound, err.Error())
	}
	resp := entity.MakeOrderResponse(order)
	return resp, nil
}

func (uc *OrderUsecase) PayOrder(ctx context.Context, req entity.OrderPaymentRequest) (*entity.OrderPaymentResponse, error) {
	order, err := uc.repo.GetByIdAndStatus(ctx, req.Id, model.OrderStatusWaitingForPayment)
	if err != nil {
		uc.log.Usecase("OrderPayment").Error(err)
		return nil, status.Error(codes.NotFound, err.Error())
	}
	paymentMethod := transaction.NewPayment(req.Method, req.AccountName, req.AccountNumber)
	paymentResp, err := paymentMethod.Pay(order.TotalPrice())
	if err != nil {
		uc.log.Usecase("OrderPayment").Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = uc.repo.UpdateStatus(ctx, req.Id, model.OrderStatusPaymentSucceed, model.OrderStatusWaitingForPayment)
	if err != nil {
		uc.log.Usecase("UpdateOrderStatus").Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	resp := entity.MakeOrderPaymentResponse(paymentResp)
	return resp, nil
}

func (uc *OrderUsecase) ShipOrder(ctx context.Context, req entity.ShippingRequest) (*entity.ShippingResponse, error) {
	order, err := uc.repo.GetByIdAndStatus(ctx, req.OrderId, model.OrderStatusPaymentSucceed)
	if err != nil {
		uc.log.Usecase("ShipOrder").Error(err)
		return nil, status.Error(codes.NotFound, err.Error())
	}
	shipper := shipping.New(req.ShipperName)
	shippingResp, err := shipper.Create(&shipping.Data{
		Address:      req.Address,
		CustomerName: req.CustomerName,
		ShippingType: req.ShippingType,
	})
	if err != nil {
		uc.log.Usecase("ShipOrder").Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = uc.repo.UpdateShipping(ctx, order.ID, shippingResp.Id)
	if err != nil {
		uc.log.Usecase("ShipOrder").Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	resp := entity.MakeShippingResponse(shippingResp)
	return resp, nil
}

func (uc *OrderUsecase) CancelOrder(ctx context.Context, req entity.CancelOrderRequest) error {
	err := uc.repo.UpdateStatus(ctx, req.OrderId, model.OrderStatusCancelled, model.OrderStatusCanCancel...)
	if err != nil {
		uc.log.Usecase("CancelOrder").Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (uc *OrderUsecase) ListenOrderStatusChanges(ctx context.Context, delivery amqp.Delivery) error {
	payload := new(rabbit.Payload[*entity.UpdateShippingRabbitData])
	err := payload.UnmarshalBinary(delivery.Body)
	if err != nil {
		uc.log.Usecase("ListenOrderStatusChanges").Error(err)
		return err
	}
	order, err := uc.repo.GetByShippingId(ctx, payload.Data.ShippingId)
	if err != nil {
		uc.log.Usecase("ListenOrderStatusChanges").Error(err)
		return err
	}
	err = uc.repo.UpdateStatus(ctx, order.ID, model.OrderStatus(payload.Data.Status), model.OrderStatusWaitingForShipment)
	if err != nil {
		uc.log.Usecase("ListenOrderStatusChanges").Error(err)
		return err
	}
	return nil
}

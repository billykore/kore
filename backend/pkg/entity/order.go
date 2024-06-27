package entity

import (
	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/shipping"
	"github.com/billykore/kore/backend/pkg/transaction"
	"github.com/billykore/kore/backend/pkg/types"
)

type OrderRequest struct {
	Id uint `param:"orderId"`
}

type OrderResponse struct {
	Id            uint   `json:"id"`
	UserId        int    `json:"userId"`
	CartIds       []int  `json:"cartIds"`
	PaymentMethod string `json:"paymentMethod"`
	Status        string `json:"status"`
	ShippingId    int    `json:"shippingId"`
}

// MakeOrderResponse makes OrderResponse from model.Order model.
func MakeOrderResponse(m *model.Order) *OrderResponse {
	return &OrderResponse{
		Id:            m.ID,
		UserId:        m.UserId,
		CartIds:       m.IntCartIds(),
		PaymentMethod: m.PaymentMethod,
		Status:        m.Status.String(),
		ShippingId:    m.ShippingId,
	}
}

type CheckoutRequest struct {
	UserId        int        `json:"userId"`
	PaymentMethod string     `json:"paymentMethod"`
	AccountNumber string     `json:"accountNumber"`
	AccountName   string     `json:"accountName"`
	Items         []CartItem `json:"items"`
}

type CartItem struct {
	Id uint `json:"id"`
}

type UpdateOrderRequest struct {
	Id     uint   `param:"orderId"`
	Status string `param:"status"`
}

type OrderPaymentRequest struct {
	Id            uint   `param:"orderId"`
	Method        string `json:"method"`
	AccountNumber string `json:"accountNumber"`
	AccountName   string `json:"accountName"`
}

type OrderPaymentResponse struct {
	Method    string      `json:"method"`
	Amount    types.Money `json:"amount"`
	PaymentId int         `json:"paymentId"`
}

func MakeOrderPaymentResponse(paymentResp *transaction.PaymentResponse) *OrderPaymentResponse {
	return &OrderPaymentResponse{
		Method:    paymentResp.Method,
		Amount:    paymentResp.Amount,
		PaymentId: paymentResp.PaymentId,
	}
}

type ShippingRequest struct {
	OrderId      uint   `param:"orderId"`
	ShipperName  string `json:"shipperName"`
	ShippingType string `json:"shippingType"`
	Address      string `json:"address"`
	CustomerName string `json:"customerName"`
}

type ShippingResponse struct {
	Id          int         `json:"id"`
	Fee         types.Money `json:"fee"`
	Status      string      `json:"status"`
	ShipperName string      `json:"shipperName"`
}

func MakeShippingResponse(shippingResp *shipping.Response) *ShippingResponse {
	return &ShippingResponse{
		Id:          shippingResp.Id,
		Fee:         shippingResp.Fee,
		Status:      shippingResp.Status,
		ShipperName: shippingResp.ShipperName,
	}
}

type CancelOrderRequest struct {
	OrderId uint `param:"orderId"`
}

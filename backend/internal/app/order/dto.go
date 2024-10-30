package order

import (
	"github.com/billykore/kore/backend/internal/domain/order"
	"github.com/billykore/kore/backend/internal/infra/shipping"
	"github.com/billykore/kore/backend/internal/infra/transaction"
	"github.com/billykore/kore/backend/pkg/types"
)

type GetRequest struct {
	Id uint `param:"orderId"`
}

type Response struct {
	Id            uint   `json:"id"`
	Username      string `json:"userId"`
	CartIds       []int  `json:"cartIds"`
	PaymentMethod string `json:"paymentMethod"`
	Status        string `json:"status"`
	ShippingId    int    `json:"shippingId"`
}

// MakeResponse makes OtpResponse from Order entity.
func MakeResponse(m *order.Order) *Response {
	return &Response{
		Id:            m.ID,
		Username:      m.Username,
		CartIds:       m.IntCartIds(),
		PaymentMethod: m.PaymentMethod,
		Status:        m.Status.String(),
		ShippingId:    m.ShippingId,
	}
}

type CheckoutRequest struct {
	PaymentMethod string     `json:"paymentMethod"`
	AccountNumber string     `json:"accountNumber"`
	AccountName   string     `json:"accountName"`
	Items         []CartItem `json:"items"`
}

// CartIds gets id from all item in the order.
func (r *CheckoutRequest) CartIds() []uint {
	var ids []uint
	for _, item := range r.Items {
		ids = append(ids, item.Id)
	}
	return ids
}

type CartItem struct {
	Id uint `json:"id"`
}

type UpdateOrderRequest struct {
	Id     uint   `param:"orderId"`
	Status string `param:"status"`
}

type PaymentRequest struct {
	Id            uint   `param:"orderId" swaggerignore:"true"`
	Method        string `json:"method"`
	AccountNumber string `json:"accountNumber"`
	AccountName   string `json:"accountName"`
}

type PaymentResponse struct {
	Method    string      `json:"method"`
	Amount    types.Money `json:"amount"`
	PaymentId int         `json:"paymentId"`
}

func MakePaymentResponse(paymentResp *transaction.PaymentResponse) *PaymentResponse {
	return &PaymentResponse{
		Method:    paymentResp.Method,
		Amount:    paymentResp.Amount,
		PaymentId: paymentResp.PaymentId,
	}
}

type ShippingRequest struct {
	OrderId      uint   `param:"orderId" swaggerignore:"true"`
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

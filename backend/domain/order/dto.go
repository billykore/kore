package order

import (
	"github.com/billykore/kore/backend/pkg/types"
)

type GetRequest struct {
	Id uint `param:"orderId"`
}

type Response struct {
	Id            uint   `json:"id"`
	Username      string `json:"userId"`
	CartIds       []int  `json:"cartIds"`
	PaymentMethod string `json:"payment"`
	Status        string `json:"status"`
	ShippingId    int    `json:"shippingId"`
}

// MakeResponse makes OtpResponse from Order entity.
func MakeResponse(m *Order) *Response {
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
	PaymentMethod string     `json:"payment"`
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
	AccountName   string `json:"accountName"`
	AccountNumber string `json:"accountNumber"`
}

type PaymentResponse struct {
	Method        string      `json:"method"`
	Amount        types.Money `json:"amount"`
	PaymentId     int         `json:"paymentId"`
	AccountName   string      `json:"accountName"`
	AccountNumber string      `json:"accountNumber"`
}

type ShippingRequest struct {
	OrderId      uint   `param:"orderId" swaggerignore:"true"`
	ShipperName  string `json:"shipperName"`
	ShippingType string `json:"shippingType"`
	Address      string `json:"address"`
	CustomerName string `json:"customerName"`
}

type ShippingData struct {
	Address      string
	CustomerName string
	ShippingType string
}

type ShippingResponse struct {
	Id           int         `json:"id"`
	Fee          types.Money `json:"fee"`
	Status       string      `json:"status"`
	ShipperName  string      `json:"shipperName"`
	Address      string      `json:"address"`
	CustomerName string      `json:"customerName"`
}

type CancelOrderRequest struct {
	OrderId uint `param:"orderId"`
}

type StatusChangeData struct {
	ShippingId uint   `json:"shippingId"`
	Status     string `json:"status"`
}

package entity

import (
	"github.com/billykore/kore/backend/pkg/types"
)

var shippingFee = map[string]types.Money{
	"regular": 20_000,
	"express": 30_000,
	"sameDay": 50_000,
}

// GetShippingFee return fee base on the shipping type.
func GetShippingFee(shippingType string) types.Money {
	if v, ok := shippingFee[shippingType]; ok {
		return v
	}
	return shippingFee["regular"]
}

type CreateShippingRequest struct {
	OrderId      int    `json:"orderId"`
	ShipperName  string `json:"shipperName"`
	ShippingType string `json:"shippingType"`
	Address      string `json:"address"`
	CustomerName string `json:"customerName"`
	SenderName   string `json:"senderName"`
}

type CreateShippingResponse struct {
	Id          uint        `json:"id"`
	Fee         types.Money `json:"fee"`
	Status      string      `json:"status"`
	ShipperName string      `json:"shipperName"`
}

type UpdateShippingStatusRequest struct {
	Id            uint   `param:"shippingId" swaggerignore:"true"`
	NewStatus     string `json:"newStatus"`
	CurrentStatus string `json:"currentStatus"`
}

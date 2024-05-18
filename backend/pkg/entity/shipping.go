package entity

import (
	"github.com/billykore/kore/backend/pkg/types"
)

func GetShippingFee(shippingType string) types.Money {
	var shippingFee = map[string]types.Money{
		"regular": 20_000,
		"express": 30_000,
		"sameDay": 50_000,
	}
	if v, ok := shippingFee[shippingType]; ok {
		return v
	}
	return shippingFee["regular"]
}

type CreateShippingRequest struct {
	OrderId      int
	ShipperName  string
	ShippingType string
	Address      string
	CustomerName string
	SenderName   string
}

type CreateShippingResponse struct {
	Id          uint        `json:"id"`
	Fee         types.Money `json:"fee"`
	Status      string      `json:"status"`
	ShipperName string      `json:"shipperName"`
}

type UpdateShippingStatusRequest struct {
	Id            uint `param:"shippingId"`
	NewStatus     string
	CurrentStatus string
}

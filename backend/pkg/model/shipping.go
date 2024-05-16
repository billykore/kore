package model

import (
	"github.com/billykore/kore/backend/pkg/types"
	"gorm.io/gorm"
)

const (
	ShippingStatusCreated = "created"
)

type Shipping struct {
	gorm.Model
	ShipperName     string
	ShippingType    string
	CustomerAddress string
	CustomerName    string
	SenderName      string
	Status          string
	Fee             types.Money
}

func (Shipping) TableName() string {
	return "shipping"
}

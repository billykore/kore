package model

import (
	"github.com/billykore/kore/backend/pkg/types"
	"gorm.io/gorm"
)

type ShippingStatus string

const (
	ShippingStatusCreated ShippingStatus = "created"
)

func (s ShippingStatus) String() string {
	return string(s)
}

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

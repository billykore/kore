package shipping

import (
	"github.com/billykore/kore/backend/pkg/types"
	"gorm.io/gorm"
)

type Status string

const (
	StatusCreated Status = "created"
)

func (s Status) String() string {
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

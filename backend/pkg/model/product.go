package model

import (
	"github.com/billykore/kore/backend/pkg/types"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name             string
	Description      string
	SKU              string
	Price            types.Money
	CategoryId       int
	InventoryId      int
	DiscountId       int
	ProductCategory  ProductCategory  `gorm:"foreignKey:CategoryId"`
	ProductInventory ProductInventory `gorm:"foreignKey:InventoryId"`
	Discount         Discount         `gorm:"foreignKey:DiscountId"`
}

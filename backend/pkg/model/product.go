package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name             string
	Description      string
	SKU              string
	Price            uint64
	CategoryId       int
	InventoryId      int
	DiscountId       int
	ProductCategory  ProductCategory  `gorm:"foreignKey:CategoryId"`
	ProductInventory ProductInventory `gorm:"foreignKey:InventoryId"`
	Discount         Discount         `gorm:"foreignKey:DiscountId"`
}

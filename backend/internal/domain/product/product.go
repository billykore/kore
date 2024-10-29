package product

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

type ProductCategory struct {
	gorm.Model
	Name        string
	Description string
}

type ProductInventory struct {
	gorm.Model
	Quantity    int
	Description string
}

type Discount struct {
	gorm.Model
	Name               string
	Description        string
	DiscountPercentage float64
	Active             bool
}

type Cart struct {
	gorm.Model
	Username  string
	ProductId int
	Quantity  int
	Product   Product `gorm:"foreignKey:ProductId"`
}

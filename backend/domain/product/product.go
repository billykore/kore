package product

import (
	"github.com/billykore/kore/backend/pkg/types"
	"gorm.io/gorm"
)

// Product is product entity.
type Product struct {
	gorm.Model
	Name        string
	Description string
	SKU         string
	Price       types.Money
	CategoryId  int
	InventoryId int
	DiscountId  int
	Category    Category  `gorm:"foreignKey:CategoryId"`
	Inventory   Inventory `gorm:"foreignKey:InventoryId"`
	Discount    Discount  `gorm:"foreignKey:DiscountId"`
}

// Category is category entity.
type Category struct {
	gorm.Model
	Name        string
	Description string
}

func (Category) TableName() string {
	return "product_categories"
}

// Inventory is inventory entity.
type Inventory struct {
	gorm.Model
	Quantity    int
	Description string
}

func (Inventory) TableName() string {
	return "product_inventories"
}

// Discount is discount entity.
type Discount struct {
	gorm.Model
	Name               string
	Description        string
	DiscountPercentage float64
	Active             bool
}

// Cart is cart entity.
type Cart struct {
	gorm.Model
	Username  string
	ProductId int
	Quantity  int
	Product   Product `gorm:"foreignKey:ProductId"`
}

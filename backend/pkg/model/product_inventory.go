package model

import "gorm.io/gorm"

type ProductInventory struct {
	gorm.Model
	Quantity    int
	Description string
}
